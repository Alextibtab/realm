package main

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	pprof "net/http/pprof"
	"os"

	"entgo.io/contrib/entgql"
	"github.com/kcarretto/realm/tavern/internal/graphql"
	"github.com/kcarretto/realm/tavern/tomes"

	gqlgraphql "github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/kcarretto/realm/tavern/internal/auth"
	"github.com/kcarretto/realm/tavern/internal/cdn"
	"github.com/kcarretto/realm/tavern/internal/ent"
	"github.com/kcarretto/realm/tavern/internal/ent/migrate"
	tavernhttp "github.com/kcarretto/realm/tavern/internal/http"
	"github.com/kcarretto/realm/tavern/internal/www"
	"github.com/urfave/cli"
)

// Version of Tavern being run
const Version = "v0.0.1"

func newApp(ctx context.Context, options ...func(*Config)) (app *cli.App) {
	app = cli.NewApp()
	app.Name = "tavern"
	app.Description = "Teamserver implementation for Realm, see https://docs.realm.pub for more details"
	app.Usage = "Time for an Adventure!"
	app.Version = Version
	app.Action = cli.ActionFunc(func(*cli.Context) error {
		return run(ctx, options...)
	})
	return
}

func run(ctx context.Context, options ...func(*Config)) error {
	srv, err := NewServer(ctx, options...)
	if err != nil {
		return err
	}
	defer srv.client.Close()

	// Listen & Serve HTTP Traffic
	log.Printf("Starting HTTP server on %s", srv.HTTP.Addr)
	if err := srv.HTTP.ListenAndServe(); err != nil {
		return fmt.Errorf("stopped http server: %w", err)
	}

	return nil
}

// Server responsible for handling Tavern requests.
type Server struct {
	HTTP   *http.Server
	client *ent.Client
}

// Close should always be called to clean up a Tavern server.
func (srv *Server) Close() error {
	return srv.client.Close()
}

// NewServer initializes a Tavern HTTP server with the provided configuration.
func NewServer(ctx context.Context, options ...func(*Config)) (*Server, error) {
	// Generate server key pair
	pubKey, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatalf("[FATAL] failed to generate ed25519 keypair: %v", err)
	}
	// Initialize Config
	cfg := &Config{}
	for _, opt := range options {
		opt(cfg)
	}

	// Create Ent Client
	client, err := cfg.Connect()
	if err != nil {
		return nil, fmt.Errorf("failed to open graph: %w", err)
	}

	// Initialize Graph Schema
	if err := client.Schema.Create(
		ctx,
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		client.Close()
		return nil, fmt.Errorf("failed to initialize graph schema: %w", err)
	}

	// Load Default Tomes
	if err := tomes.UploadTomes(ctx, client, tomes.FileSystem); err != nil {
		client.Close()
		return nil, fmt.Errorf("failed to upload default tomes: %w", err)
	}

	// Initialize Test Data
	if cfg.IsTestDataEnabled() {
		createTestData(ctx, client)
	}

	// Configure Authentication
	var withAuthentication tavernhttp.Option
	if cfg.oauth.ClientID != "" {
		withAuthentication = tavernhttp.WithAuthenticationCookie(client)
	} else {
		withAuthentication = tavernhttp.WithAuthenticationBypass(client)
	}

	// Configure Request Logging
	httpLogger := log.New(os.Stderr, "[HTTP] ", log.Flags())

	// Route Map
	routes := tavernhttp.RouteMap{
		"/status":      tavernhttp.Endpoint{Handler: newStatusHandler()},
		"/oauth/login": tavernhttp.Endpoint{Handler: auth.NewOAuthLoginHandler(cfg.oauth, privKey)},
		"/oauth/authorize": tavernhttp.Endpoint{Handler: auth.NewOAuthAuthorizationHandler(
			cfg.oauth,
			pubKey,
			client,
			"https://www.googleapis.com/oauth2/v3/userinfo",
		)},
		"/graphql":    tavernhttp.Endpoint{Handler: newGraphQLHandler(client)},
		"/cdn/":       tavernhttp.Endpoint{Handler: cdn.NewDownloadHandler(client)},
		"/cdn/upload": tavernhttp.Endpoint{Handler: cdn.NewUploadHandler(client)},
		"/": tavernhttp.Endpoint{
			Handler:          www.NewHandler(httpLogger),
			LoginRedirectURI: "/oauth/login",
		},
		"/playground": tavernhttp.Endpoint{
			Handler:          playground.Handler("Tavern", "/graphql"),
			LoginRedirectURI: "/oauth/login",
		},
	}

	// Setup Profiling
	if cfg.IsPProfEnabled() {
		log.Printf("[WARN] Performance profiling is enabled, do not use in production as this may leak sensitive information")
		registerProfiler(routes)
	}

	// Create Tavern HTTP Server
	srv := tavernhttp.NewServer(
		routes,
		withAuthentication,
		tavernhttp.WithRequestLogging(httpLogger),
	)

	// Initialize HTTP Server
	if cfg.srv == nil {
		cfg.srv = &http.Server{
			Addr:    "0.0.0.0:80",
			Handler: srv,
		}
	} else {
		cfg.srv.Handler = srv
	}

	return &Server{
		HTTP:   cfg.srv,
		client: client,
	}, nil
}

func newGraphQLHandler(client *ent.Client) http.Handler {
	srv := handler.NewDefaultServer(graphql.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})

	// GraphQL Logging
	gqlLogger := log.New(os.Stderr, "[GraphQL] ", log.Flags())
	srv.AroundOperations(func(ctx context.Context, next gqlgraphql.OperationHandler) gqlgraphql.ResponseHandler {
		oc := gqlgraphql.GetOperationContext(ctx)
		reqVars, err := json.Marshal(oc.Variables)
		if err != nil {
			gqlLogger.Printf("[ERROR] failed to marshal variables to JSON: %v", err)
			return next(ctx)
		}

		authName := "unknown"
		id := auth.IdentityFromContext(ctx)
		if id != nil {
			authName = id.String()
		}

		gqlLogger.Printf("%s (%s): %s", oc.OperationName, authName, string(reqVars))
		return next(ctx)
	})

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		srv.ServeHTTP(w, req)
	})
}

func registerProfiler(router tavernhttp.RouteMap) {
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)

	// Manually add support for paths linked to by index page at /debug/pprof/
	router.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
	router.Handle("/debug/pprof/heap", pprof.Handler("heap"))
	router.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
	router.Handle("/debug/pprof/block", pprof.Handler("block"))
}

// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"bytes"
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/kcarretto/realm/tavern/ent"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Mutation() MutationResolver
	Query() QueryResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	File struct {
		CreatedAt      func(childComplexity int) int
		Hash           func(childComplexity int) int
		ID             func(childComplexity int) int
		LastModifiedAt func(childComplexity int) int
		Name           func(childComplexity int) int
		Size           func(childComplexity int) int
	}

	Job struct {
		CreatedBy func(childComplexity int) int
		ID        func(childComplexity int) int
		Name      func(childComplexity int) int
		Tasks     func(childComplexity int) int
		Tome      func(childComplexity int) int
	}

	Mutation struct {
		CreateUser func(childComplexity int, input ent.CreateUserInput) int
		UpdateUser func(childComplexity int, id int, input ent.UpdateUserInput) int
	}

	PageInfo struct {
		EndCursor       func(childComplexity int) int
		HasNextPage     func(childComplexity int) int
		HasPreviousPage func(childComplexity int) int
		StartCursor     func(childComplexity int) int
	}

	Query struct {
		Files   func(childComplexity int) int
		Jobs    func(childComplexity int) int
		Node    func(childComplexity int, id int) int
		Nodes   func(childComplexity int, ids []int) int
		Tags    func(childComplexity int) int
		Targets func(childComplexity int) int
		Tomes   func(childComplexity int) int
		Users   func(childComplexity int) int
	}

	Tag struct {
		ID      func(childComplexity int) int
		Kind    func(childComplexity int) int
		Name    func(childComplexity int) int
		Targets func(childComplexity int) int
	}

	Target struct {
		ID   func(childComplexity int) int
		Name func(childComplexity int) int
		Tags func(childComplexity int) int
	}

	Task struct {
		ID   func(childComplexity int) int
		Job  func(childComplexity int) int
		Name func(childComplexity int) int
	}

	Tome struct {
		CreatedAt      func(childComplexity int) int
		Description    func(childComplexity int) int
		Hash           func(childComplexity int) int
		ID             func(childComplexity int) int
		LastModifiedAt func(childComplexity int) int
		Name           func(childComplexity int) int
		Parameters     func(childComplexity int) int
		Size           func(childComplexity int) int
	}

	User struct {
		ID          func(childComplexity int) int
		IsActivated func(childComplexity int) int
		IsAdmin     func(childComplexity int) int
		Name        func(childComplexity int) int
		PhotoURL    func(childComplexity int) int
	}
}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e}
	_ = ec
	switch typeName + "." + field {

	case "File.createdat":
		if e.complexity.File.CreatedAt == nil {
			break
		}

		return e.complexity.File.CreatedAt(childComplexity), true

	case "File.hash":
		if e.complexity.File.Hash == nil {
			break
		}

		return e.complexity.File.Hash(childComplexity), true

	case "File.id":
		if e.complexity.File.ID == nil {
			break
		}

		return e.complexity.File.ID(childComplexity), true

	case "File.lastmodifiedat":
		if e.complexity.File.LastModifiedAt == nil {
			break
		}

		return e.complexity.File.LastModifiedAt(childComplexity), true

	case "File.name":
		if e.complexity.File.Name == nil {
			break
		}

		return e.complexity.File.Name(childComplexity), true

	case "File.size":
		if e.complexity.File.Size == nil {
			break
		}

		return e.complexity.File.Size(childComplexity), true

	case "Job.createdby":
		if e.complexity.Job.CreatedBy == nil {
			break
		}

		return e.complexity.Job.CreatedBy(childComplexity), true

	case "Job.id":
		if e.complexity.Job.ID == nil {
			break
		}

		return e.complexity.Job.ID(childComplexity), true

	case "Job.name":
		if e.complexity.Job.Name == nil {
			break
		}

		return e.complexity.Job.Name(childComplexity), true

	case "Job.tasks":
		if e.complexity.Job.Tasks == nil {
			break
		}

		return e.complexity.Job.Tasks(childComplexity), true

	case "Job.tome":
		if e.complexity.Job.Tome == nil {
			break
		}

		return e.complexity.Job.Tome(childComplexity), true

	case "Mutation.createUser":
		if e.complexity.Mutation.CreateUser == nil {
			break
		}

		args, err := ec.field_Mutation_createUser_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateUser(childComplexity, args["input"].(ent.CreateUserInput)), true

	case "Mutation.updateUser":
		if e.complexity.Mutation.UpdateUser == nil {
			break
		}

		args, err := ec.field_Mutation_updateUser_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.UpdateUser(childComplexity, args["id"].(int), args["input"].(ent.UpdateUserInput)), true

	case "PageInfo.endCursor":
		if e.complexity.PageInfo.EndCursor == nil {
			break
		}

		return e.complexity.PageInfo.EndCursor(childComplexity), true

	case "PageInfo.hasNextPage":
		if e.complexity.PageInfo.HasNextPage == nil {
			break
		}

		return e.complexity.PageInfo.HasNextPage(childComplexity), true

	case "PageInfo.hasPreviousPage":
		if e.complexity.PageInfo.HasPreviousPage == nil {
			break
		}

		return e.complexity.PageInfo.HasPreviousPage(childComplexity), true

	case "PageInfo.startCursor":
		if e.complexity.PageInfo.StartCursor == nil {
			break
		}

		return e.complexity.PageInfo.StartCursor(childComplexity), true

	case "Query.files":
		if e.complexity.Query.Files == nil {
			break
		}

		return e.complexity.Query.Files(childComplexity), true

	case "Query.jobs":
		if e.complexity.Query.Jobs == nil {
			break
		}

		return e.complexity.Query.Jobs(childComplexity), true

	case "Query.node":
		if e.complexity.Query.Node == nil {
			break
		}

		args, err := ec.field_Query_node_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Node(childComplexity, args["id"].(int)), true

	case "Query.nodes":
		if e.complexity.Query.Nodes == nil {
			break
		}

		args, err := ec.field_Query_nodes_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Nodes(childComplexity, args["ids"].([]int)), true

	case "Query.tags":
		if e.complexity.Query.Tags == nil {
			break
		}

		return e.complexity.Query.Tags(childComplexity), true

	case "Query.targets":
		if e.complexity.Query.Targets == nil {
			break
		}

		return e.complexity.Query.Targets(childComplexity), true

	case "Query.tomes":
		if e.complexity.Query.Tomes == nil {
			break
		}

		return e.complexity.Query.Tomes(childComplexity), true

	case "Query.users":
		if e.complexity.Query.Users == nil {
			break
		}

		return e.complexity.Query.Users(childComplexity), true

	case "Tag.id":
		if e.complexity.Tag.ID == nil {
			break
		}

		return e.complexity.Tag.ID(childComplexity), true

	case "Tag.kind":
		if e.complexity.Tag.Kind == nil {
			break
		}

		return e.complexity.Tag.Kind(childComplexity), true

	case "Tag.name":
		if e.complexity.Tag.Name == nil {
			break
		}

		return e.complexity.Tag.Name(childComplexity), true

	case "Tag.targets":
		if e.complexity.Tag.Targets == nil {
			break
		}

		return e.complexity.Tag.Targets(childComplexity), true

	case "Target.id":
		if e.complexity.Target.ID == nil {
			break
		}

		return e.complexity.Target.ID(childComplexity), true

	case "Target.name":
		if e.complexity.Target.Name == nil {
			break
		}

		return e.complexity.Target.Name(childComplexity), true

	case "Target.tags":
		if e.complexity.Target.Tags == nil {
			break
		}

		return e.complexity.Target.Tags(childComplexity), true

	case "Task.id":
		if e.complexity.Task.ID == nil {
			break
		}

		return e.complexity.Task.ID(childComplexity), true

	case "Task.job":
		if e.complexity.Task.Job == nil {
			break
		}

		return e.complexity.Task.Job(childComplexity), true

	case "Task.name":
		if e.complexity.Task.Name == nil {
			break
		}

		return e.complexity.Task.Name(childComplexity), true

	case "Tome.createdat":
		if e.complexity.Tome.CreatedAt == nil {
			break
		}

		return e.complexity.Tome.CreatedAt(childComplexity), true

	case "Tome.description":
		if e.complexity.Tome.Description == nil {
			break
		}

		return e.complexity.Tome.Description(childComplexity), true

	case "Tome.hash":
		if e.complexity.Tome.Hash == nil {
			break
		}

		return e.complexity.Tome.Hash(childComplexity), true

	case "Tome.id":
		if e.complexity.Tome.ID == nil {
			break
		}

		return e.complexity.Tome.ID(childComplexity), true

	case "Tome.lastmodifiedat":
		if e.complexity.Tome.LastModifiedAt == nil {
			break
		}

		return e.complexity.Tome.LastModifiedAt(childComplexity), true

	case "Tome.name":
		if e.complexity.Tome.Name == nil {
			break
		}

		return e.complexity.Tome.Name(childComplexity), true

	case "Tome.parameters":
		if e.complexity.Tome.Parameters == nil {
			break
		}

		return e.complexity.Tome.Parameters(childComplexity), true

	case "Tome.size":
		if e.complexity.Tome.Size == nil {
			break
		}

		return e.complexity.Tome.Size(childComplexity), true

	case "User.id":
		if e.complexity.User.ID == nil {
			break
		}

		return e.complexity.User.ID(childComplexity), true

	case "User.isactivated":
		if e.complexity.User.IsActivated == nil {
			break
		}

		return e.complexity.User.IsActivated(childComplexity), true

	case "User.isadmin":
		if e.complexity.User.IsAdmin == nil {
			break
		}

		return e.complexity.User.IsAdmin(childComplexity), true

	case "User.name":
		if e.complexity.User.Name == nil {
			break
		}

		return e.complexity.User.Name(childComplexity), true

	case "User.photourl":
		if e.complexity.User.PhotoURL == nil {
			break
		}

		return e.complexity.User.PhotoURL(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputCreateUserInput,
		ec.unmarshalInputFileOrder,
		ec.unmarshalInputFileWhereInput,
		ec.unmarshalInputJobOrder,
		ec.unmarshalInputJobWhereInput,
		ec.unmarshalInputTagOrder,
		ec.unmarshalInputTagWhereInput,
		ec.unmarshalInputTargetOrder,
		ec.unmarshalInputTargetWhereInput,
		ec.unmarshalInputTaskOrder,
		ec.unmarshalInputTaskWhereInput,
		ec.unmarshalInputTomeOrder,
		ec.unmarshalInputTomeWhereInput,
		ec.unmarshalInputUpdateUserInput,
		ec.unmarshalInputUserWhereInput,
	)
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Query(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}
	case ast.Mutation:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Mutation(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(parsedSchema), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "../schema/scalars.graphql", Input: `scalar Time`, BuiltIn: false},
	{Name: "../schema/schema.graphql", Input: `directive @goField(forceResolver: Boolean, name: String) on FIELD_DEFINITION | INPUT_FIELD_DEFINITION
directive @goModel(model: String, models: [String!]) on OBJECT | INPUT_OBJECT | SCALAR | ENUM | INTERFACE | UNION
"""
CreateUserInput is used for create User object.
Input was generated by ent.
"""
input CreateUserInput {
  """The name displayed for the user"""
  name: String!
  """OAuth Subject ID of the user"""
  oauthid: String!
  """URL to the user's profile photo."""
  photourl: String!
  """True if the user is active and able to authenticate"""
  isactivated: Boolean
  """True if the user is an Admin"""
  isadmin: Boolean
}
"""
Define a Relay Cursor type:
https://relay.dev/graphql/connections.htm#sec-Cursor
"""
scalar Cursor
type File implements Node {
  id: ID!
  """The name of the file, used to reference it for downloads"""
  name: String!
  """The size of the file in bytes"""
  size: Int!
  """A SHA3 digest of the content field"""
  hash: String!
  """The timestamp for when the File was created"""
  createdat: Time! @goField(name: "CreatedAt", forceResolver: false)
  """The timestamp for when the File was last modified"""
  lastmodifiedat: Time! @goField(name: "LastModifiedAt", forceResolver: false)
}
"""Ordering options for File connections"""
input FileOrder {
  """The ordering direction."""
  direction: OrderDirection! = ASC
  """The field by which to order Files."""
  field: FileOrderField!
}
"""Properties by which File connections can be ordered."""
enum FileOrderField {
  NAME
  SIZE
  CREATED_AT
  LAST_MODIFIED_AT
}
"""
FileWhereInput is used for filtering File objects.
Input was generated by ent.
"""
input FileWhereInput {
  not: FileWhereInput
  and: [FileWhereInput!]
  or: [FileWhereInput!]
  """id field predicates"""
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """name field predicates"""
  name: String
  nameNEQ: String
  nameIn: [String!]
  nameNotIn: [String!]
  nameGT: String
  nameGTE: String
  nameLT: String
  nameLTE: String
  nameContains: String
  nameHasPrefix: String
  nameHasSuffix: String
  nameEqualFold: String
  nameContainsFold: String
  """size field predicates"""
  size: Int
  sizeNEQ: Int
  sizeIn: [Int!]
  sizeNotIn: [Int!]
  sizeGT: Int
  sizeGTE: Int
  sizeLT: Int
  sizeLTE: Int
  """hash field predicates"""
  hash: String
  hashNEQ: String
  hashIn: [String!]
  hashNotIn: [String!]
  hashGT: String
  hashGTE: String
  hashLT: String
  hashLTE: String
  hashContains: String
  hashHasPrefix: String
  hashHasSuffix: String
  hashEqualFold: String
  hashContainsFold: String
  """createdAt field predicates"""
  createdat: Time
  createdatNEQ: Time
  createdatIn: [Time!]
  createdatNotIn: [Time!]
  createdatGT: Time
  createdatGTE: Time
  createdatLT: Time
  createdatLTE: Time
  """lastModifiedAt field predicates"""
  lastmodifiedat: Time
  lastmodifiedatNEQ: Time
  lastmodifiedatIn: [Time!]
  lastmodifiedatNotIn: [Time!]
  lastmodifiedatGT: Time
  lastmodifiedatGTE: Time
  lastmodifiedatLT: Time
  lastmodifiedatLTE: Time
}
type Job implements Node {
  id: ID!
  """Name of the job"""
  name: String!
  createdby: User! @goField(name: "CreatedBy", forceResolver: false)
  tome: Tome!
  tasks: [Task!]
}
"""Ordering options for Job connections"""
input JobOrder {
  """The ordering direction."""
  direction: OrderDirection! = ASC
  """The field by which to order Jobs."""
  field: JobOrderField!
}
"""Properties by which Job connections can be ordered."""
enum JobOrderField {
  NAME
}
"""
JobWhereInput is used for filtering Job objects.
Input was generated by ent.
"""
input JobWhereInput {
  not: JobWhereInput
  and: [JobWhereInput!]
  or: [JobWhereInput!]
  """id field predicates"""
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """name field predicates"""
  name: String
  nameNEQ: String
  nameIn: [String!]
  nameNotIn: [String!]
  nameGT: String
  nameGTE: String
  nameLT: String
  nameLTE: String
  nameContains: String
  nameHasPrefix: String
  nameHasSuffix: String
  nameEqualFold: String
  nameContainsFold: String
  """createdBy edge predicates"""
  hasCreatedBy: Boolean
  hasCreatedByWith: [UserWhereInput!]
  """tome edge predicates"""
  hasTome: Boolean
  hasTomeWith: [TomeWhereInput!]
  """tasks edge predicates"""
  hasTasks: Boolean
  hasTasksWith: [TaskWhereInput!]
}
"""
An object with an ID.
Follows the [Relay Global Object Identification Specification](https://relay.dev/graphql/objectidentification.htm)
"""
interface Node @goModel(model: "github.com/kcarretto/realm/tavern/ent.Noder") {
  """The id of the object."""
  id: ID!
}
"""Possible directions in which to order a list of items when provided an ` + "`" + `orderBy` + "`" + ` argument."""
enum OrderDirection {
  """Specifies an ascending order for a given ` + "`" + `orderBy` + "`" + ` argument."""
  ASC
  """Specifies a descending order for a given ` + "`" + `orderBy` + "`" + ` argument."""
  DESC
}
"""
Information about pagination in a connection.
https://relay.dev/graphql/connections.htm#sec-undefined.PageInfo
"""
type PageInfo {
  """When paginating forwards, are there more items?"""
  hasNextPage: Boolean!
  """When paginating backwards, are there more items?"""
  hasPreviousPage: Boolean!
  """When paginating backwards, the cursor to continue."""
  startCursor: Cursor
  """When paginating forwards, the cursor to continue."""
  endCursor: Cursor
}
type Query {
  """Fetches an object given its ID."""
  node(
    """ID of the object."""
    id: ID!
  ): Node
  """Lookup nodes by a list of IDs."""
  nodes(
    """The list of node IDs."""
    ids: [ID!]!
  ): [Node]!
  files: [File!]!
  jobs: [Job!]!
  tags: [Tag!]!
  targets: [Target!]!
  tomes: [Tome!]!
  users: [User!]!
}
type Tag implements Node {
  id: ID!
  """Name of the tag"""
  name: String!
  """Describes the type of tag this is"""
  kind: TagTagKind!
  targets: [Target!]!
}
"""Ordering options for Tag connections"""
input TagOrder {
  """The ordering direction."""
  direction: OrderDirection! = ASC
  """The field by which to order Tags."""
  field: TagOrderField!
}
"""Properties by which Tag connections can be ordered."""
enum TagOrderField {
  NAME
}
"""TagTagKind is enum for the field kind"""
enum TagTagKind @goModel(model: "github.com/kcarretto/realm/tavern/ent/tag.Kind") {
  group
  service
}
"""
TagWhereInput is used for filtering Tag objects.
Input was generated by ent.
"""
input TagWhereInput {
  not: TagWhereInput
  and: [TagWhereInput!]
  or: [TagWhereInput!]
  """id field predicates"""
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """name field predicates"""
  name: String
  nameNEQ: String
  nameIn: [String!]
  nameNotIn: [String!]
  nameGT: String
  nameGTE: String
  nameLT: String
  nameLTE: String
  nameContains: String
  nameHasPrefix: String
  nameHasSuffix: String
  nameEqualFold: String
  nameContainsFold: String
  """kind field predicates"""
  kind: TagTagKind
  kindNEQ: TagTagKind
  kindIn: [TagTagKind!]
  kindNotIn: [TagTagKind!]
  """targets edge predicates"""
  hasTargets: Boolean
  hasTargetsWith: [TargetWhereInput!]
}
type Target implements Node {
  id: ID!
  """Human-readable name of the target"""
  name: String!
  tags: [Tag!]
}
"""Ordering options for Target connections"""
input TargetOrder {
  """The ordering direction."""
  direction: OrderDirection! = ASC
  """The field by which to order Targets."""
  field: TargetOrderField!
}
"""Properties by which Target connections can be ordered."""
enum TargetOrderField {
  NAME
}
"""
TargetWhereInput is used for filtering Target objects.
Input was generated by ent.
"""
input TargetWhereInput {
  not: TargetWhereInput
  and: [TargetWhereInput!]
  or: [TargetWhereInput!]
  """id field predicates"""
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """name field predicates"""
  name: String
  nameNEQ: String
  nameIn: [String!]
  nameNotIn: [String!]
  nameGT: String
  nameGTE: String
  nameLT: String
  nameLTE: String
  nameContains: String
  nameHasPrefix: String
  nameHasSuffix: String
  nameEqualFold: String
  nameContainsFold: String
  """tags edge predicates"""
  hasTags: Boolean
  hasTagsWith: [TagWhereInput!]
}
type Task implements Node {
  id: ID!
  """Name of the task"""
  name: String!
  job: Job!
}
"""Ordering options for Task connections"""
input TaskOrder {
  """The ordering direction."""
  direction: OrderDirection! = ASC
  """The field by which to order Tasks."""
  field: TaskOrderField!
}
"""Properties by which Task connections can be ordered."""
enum TaskOrderField {
  NAME
}
"""
TaskWhereInput is used for filtering Task objects.
Input was generated by ent.
"""
input TaskWhereInput {
  not: TaskWhereInput
  and: [TaskWhereInput!]
  or: [TaskWhereInput!]
  """id field predicates"""
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """name field predicates"""
  name: String
  nameNEQ: String
  nameIn: [String!]
  nameNotIn: [String!]
  nameGT: String
  nameGTE: String
  nameLT: String
  nameLTE: String
  nameContains: String
  nameHasPrefix: String
  nameHasSuffix: String
  nameEqualFold: String
  nameContainsFold: String
  """job edge predicates"""
  hasJob: Boolean
  hasJobWith: [JobWhereInput!]
}
type Tome implements Node {
  id: ID!
  """Name of the tome"""
  name: String!
  """Information about the tome"""
  description: String!
  """JSON string describing what parameters are used with the tome"""
  parameters: String!
  """The size of the tome in bytes"""
  size: Int!
  """A SHA3 digest of the content field"""
  hash: String!
  """The timestamp for when the Tome was created"""
  createdat: Time! @goField(name: "CreatedAt", forceResolver: false)
  """The timestamp for when the Tome was last modified"""
  lastmodifiedat: Time! @goField(name: "LastModifiedAt", forceResolver: false)
}
"""Ordering options for Tome connections"""
input TomeOrder {
  """The ordering direction."""
  direction: OrderDirection! = ASC
  """The field by which to order Tomes."""
  field: TomeOrderField!
}
"""Properties by which Tome connections can be ordered."""
enum TomeOrderField {
  NAME
  SIZE
  CREATED_AT
  LAST_MODIFIED_AT
}
"""
TomeWhereInput is used for filtering Tome objects.
Input was generated by ent.
"""
input TomeWhereInput {
  not: TomeWhereInput
  and: [TomeWhereInput!]
  or: [TomeWhereInput!]
  """id field predicates"""
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """name field predicates"""
  name: String
  nameNEQ: String
  nameIn: [String!]
  nameNotIn: [String!]
  nameGT: String
  nameGTE: String
  nameLT: String
  nameLTE: String
  nameContains: String
  nameHasPrefix: String
  nameHasSuffix: String
  nameEqualFold: String
  nameContainsFold: String
  """description field predicates"""
  description: String
  descriptionNEQ: String
  descriptionIn: [String!]
  descriptionNotIn: [String!]
  descriptionGT: String
  descriptionGTE: String
  descriptionLT: String
  descriptionLTE: String
  descriptionContains: String
  descriptionHasPrefix: String
  descriptionHasSuffix: String
  descriptionEqualFold: String
  descriptionContainsFold: String
  """parameters field predicates"""
  parameters: String
  parametersNEQ: String
  parametersIn: [String!]
  parametersNotIn: [String!]
  parametersGT: String
  parametersGTE: String
  parametersLT: String
  parametersLTE: String
  parametersContains: String
  parametersHasPrefix: String
  parametersHasSuffix: String
  parametersEqualFold: String
  parametersContainsFold: String
  """size field predicates"""
  size: Int
  sizeNEQ: Int
  sizeIn: [Int!]
  sizeNotIn: [Int!]
  sizeGT: Int
  sizeGTE: Int
  sizeLT: Int
  sizeLTE: Int
  """hash field predicates"""
  hash: String
  hashNEQ: String
  hashIn: [String!]
  hashNotIn: [String!]
  hashGT: String
  hashGTE: String
  hashLT: String
  hashLTE: String
  hashContains: String
  hashHasPrefix: String
  hashHasSuffix: String
  hashEqualFold: String
  hashContainsFold: String
  """createdAt field predicates"""
  createdat: Time
  createdatNEQ: Time
  createdatIn: [Time!]
  createdatNotIn: [Time!]
  createdatGT: Time
  createdatGTE: Time
  createdatLT: Time
  createdatLTE: Time
  """lastModifiedAt field predicates"""
  lastmodifiedat: Time
  lastmodifiedatNEQ: Time
  lastmodifiedatIn: [Time!]
  lastmodifiedatNotIn: [Time!]
  lastmodifiedatGT: Time
  lastmodifiedatGTE: Time
  lastmodifiedatLT: Time
  lastmodifiedatLTE: Time
}
"""
UpdateUserInput is used for update User object.
Input was generated by ent.
"""
input UpdateUserInput {
  """The name displayed for the user"""
  name: String
  """URL to the user's profile photo."""
  photourl: String
  """True if the user is active and able to authenticate"""
  isactivated: Boolean
  """True if the user is an Admin"""
  isadmin: Boolean
}
type User implements Node {
  id: ID!
  """The name displayed for the user"""
  name: String!
  """URL to the user's profile photo."""
  photourl: String! @goField(name: "PhotoURL", forceResolver: false)
  """True if the user is active and able to authenticate"""
  isactivated: Boolean! @goField(name: "IsActivated", forceResolver: false)
  """True if the user is an Admin"""
  isadmin: Boolean! @goField(name: "IsAdmin", forceResolver: false)
}
"""
UserWhereInput is used for filtering User objects.
Input was generated by ent.
"""
input UserWhereInput {
  not: UserWhereInput
  and: [UserWhereInput!]
  or: [UserWhereInput!]
  """id field predicates"""
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """Name field predicates"""
  name: String
  nameNEQ: String
  nameIn: [String!]
  nameNotIn: [String!]
  nameGT: String
  nameGTE: String
  nameLT: String
  nameLTE: String
  nameContains: String
  nameHasPrefix: String
  nameHasSuffix: String
  nameEqualFold: String
  nameContainsFold: String
  """PhotoURL field predicates"""
  photourl: String
  photourlNEQ: String
  photourlIn: [String!]
  photourlNotIn: [String!]
  photourlGT: String
  photourlGTE: String
  photourlLT: String
  photourlLTE: String
  photourlContains: String
  photourlHasPrefix: String
  photourlHasSuffix: String
  photourlEqualFold: String
  photourlContainsFold: String
  """IsActivated field predicates"""
  isactivated: Boolean
  isactivatedNEQ: Boolean
  """IsAdmin field predicates"""
  isadmin: Boolean
  isadminNEQ: Boolean
}
`, BuiltIn: false},
	{Name: "../schema/user.graphql", Input: `type Mutation {
    # The input and the output are types generated by Ent.
    createUser(input: CreateUserInput!): User
    updateUser(id: ID!, input: UpdateUserInput!): User
}`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
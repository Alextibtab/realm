// Code generated by ent, DO NOT EDIT.

package repository

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the repository type in the database.
	Label = "repository"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldLastModifiedAt holds the string denoting the last_modified_at field in the database.
	FieldLastModifiedAt = "last_modified_at"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldPublicKey holds the string denoting the public_key field in the database.
	FieldPublicKey = "public_key"
	// FieldPrivateKey holds the string denoting the private_key field in the database.
	FieldPrivateKey = "private_key"
	// FieldLastImportedAt holds the string denoting the last_imported_at field in the database.
	FieldLastImportedAt = "last_imported_at"
	// EdgeTomes holds the string denoting the tomes edge name in mutations.
	EdgeTomes = "tomes"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the repository in the database.
	Table = "repositories"
	// TomesTable is the table that holds the tomes relation/edge.
	TomesTable = "tomes"
	// TomesInverseTable is the table name for the Tome entity.
	// It exists in this package in order to avoid circular dependency with the "tome" package.
	TomesInverseTable = "tomes"
	// TomesColumn is the table column denoting the tomes relation/edge.
	TomesColumn = "tome_repository"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "repositories"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "repository_owner"
)

// Columns holds all SQL columns for repository fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldLastModifiedAt,
	FieldURL,
	FieldPublicKey,
	FieldPrivateKey,
	FieldLastImportedAt,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "repositories"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"repository_owner",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "realm.pub/tavern/internal/ent/runtime"
var (
	Hooks [1]ent.Hook
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultLastModifiedAt holds the default value on creation for the "last_modified_at" field.
	DefaultLastModifiedAt func() time.Time
	// UpdateDefaultLastModifiedAt holds the default value on update for the "last_modified_at" field.
	UpdateDefaultLastModifiedAt func() time.Time
	// URLValidator is a validator for the "url" field. It is called by the builders before save.
	URLValidator func(string) error
	// PublicKeyValidator is a validator for the "public_key" field. It is called by the builders before save.
	PublicKeyValidator func(string) error
	// PrivateKeyValidator is a validator for the "private_key" field. It is called by the builders before save.
	PrivateKeyValidator func(string) error
)

// OrderOption defines the ordering options for the Repository queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByLastModifiedAt orders the results by the last_modified_at field.
func ByLastModifiedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastModifiedAt, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByPublicKey orders the results by the public_key field.
func ByPublicKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPublicKey, opts...).ToFunc()
}

// ByPrivateKey orders the results by the private_key field.
func ByPrivateKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrivateKey, opts...).ToFunc()
}

// ByLastImportedAt orders the results by the last_imported_at field.
func ByLastImportedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastImportedAt, opts...).ToFunc()
}

// ByTomesCount orders the results by tomes count.
func ByTomesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTomesStep(), opts...)
	}
}

// ByTomes orders the results by tomes terms.
func ByTomes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTomesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}
func newTomesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TomesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, TomesTable, TomesColumn),
	)
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, OwnerTable, OwnerColumn),
	)
}

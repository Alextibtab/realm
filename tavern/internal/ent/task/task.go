// Code generated by ent, DO NOT EDIT.

package task

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldLastModifiedAt holds the string denoting the last_modified_at field in the database.
	FieldLastModifiedAt = "last_modified_at"
	// FieldClaimedAt holds the string denoting the claimed_at field in the database.
	FieldClaimedAt = "claimed_at"
	// FieldExecStartedAt holds the string denoting the exec_started_at field in the database.
	FieldExecStartedAt = "exec_started_at"
	// FieldExecFinishedAt holds the string denoting the exec_finished_at field in the database.
	FieldExecFinishedAt = "exec_finished_at"
	// FieldOutput holds the string denoting the output field in the database.
	FieldOutput = "output"
	// FieldOutputSize holds the string denoting the output_size field in the database.
	FieldOutputSize = "output_size"
	// FieldError holds the string denoting the error field in the database.
	FieldError = "error"
	// EdgeQuest holds the string denoting the quest edge name in mutations.
	EdgeQuest = "quest"
	// EdgeBeacon holds the string denoting the beacon edge name in mutations.
	EdgeBeacon = "beacon"
	// EdgeReportedProcesses holds the string denoting the reported_processes edge name in mutations.
	EdgeReportedProcesses = "reported_processes"
	// Table holds the table name of the task in the database.
	Table = "tasks"
	// QuestTable is the table that holds the quest relation/edge.
	QuestTable = "tasks"
	// QuestInverseTable is the table name for the Quest entity.
	// It exists in this package in order to avoid circular dependency with the "quest" package.
	QuestInverseTable = "quests"
	// QuestColumn is the table column denoting the quest relation/edge.
	QuestColumn = "quest_tasks"
	// BeaconTable is the table that holds the beacon relation/edge.
	BeaconTable = "tasks"
	// BeaconInverseTable is the table name for the Beacon entity.
	// It exists in this package in order to avoid circular dependency with the "beacon" package.
	BeaconInverseTable = "beacons"
	// BeaconColumn is the table column denoting the beacon relation/edge.
	BeaconColumn = "task_beacon"
	// ReportedProcessesTable is the table that holds the reported_processes relation/edge.
	ReportedProcessesTable = "host_processes"
	// ReportedProcessesInverseTable is the table name for the HostProcess entity.
	// It exists in this package in order to avoid circular dependency with the "hostprocess" package.
	ReportedProcessesInverseTable = "host_processes"
	// ReportedProcessesColumn is the table column denoting the reported_processes relation/edge.
	ReportedProcessesColumn = "task_reported_processes"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldLastModifiedAt,
	FieldClaimedAt,
	FieldExecStartedAt,
	FieldExecFinishedAt,
	FieldOutput,
	FieldOutputSize,
	FieldError,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "tasks"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"quest_tasks",
	"task_beacon",
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
	// DefaultOutputSize holds the default value on creation for the "output_size" field.
	DefaultOutputSize int
	// OutputSizeValidator is a validator for the "output_size" field. It is called by the builders before save.
	OutputSizeValidator func(int) error
)

// OrderOption defines the ordering options for the Task queries.
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

// ByClaimedAt orders the results by the claimed_at field.
func ByClaimedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClaimedAt, opts...).ToFunc()
}

// ByExecStartedAt orders the results by the exec_started_at field.
func ByExecStartedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExecStartedAt, opts...).ToFunc()
}

// ByExecFinishedAt orders the results by the exec_finished_at field.
func ByExecFinishedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExecFinishedAt, opts...).ToFunc()
}

// ByOutput orders the results by the output field.
func ByOutput(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOutput, opts...).ToFunc()
}

// ByOutputSize orders the results by the output_size field.
func ByOutputSize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOutputSize, opts...).ToFunc()
}

// ByError orders the results by the error field.
func ByError(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldError, opts...).ToFunc()
}

// ByQuestField orders the results by quest field.
func ByQuestField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newQuestStep(), sql.OrderByField(field, opts...))
	}
}

// ByBeaconField orders the results by beacon field.
func ByBeaconField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBeaconStep(), sql.OrderByField(field, opts...))
	}
}

// ByReportedProcessesCount orders the results by reported_processes count.
func ByReportedProcessesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newReportedProcessesStep(), opts...)
	}
}

// ByReportedProcesses orders the results by reported_processes terms.
func ByReportedProcesses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReportedProcessesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newQuestStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(QuestInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, QuestTable, QuestColumn),
	)
}
func newBeaconStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BeaconInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, BeaconTable, BeaconColumn),
	)
}
func newReportedProcessesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReportedProcessesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ReportedProcessesTable, ReportedProcessesColumn),
	)
}

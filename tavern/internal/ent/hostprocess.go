// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"realm.pub/tavern/internal/c2/c2pb"
	"realm.pub/tavern/internal/ent/host"
	"realm.pub/tavern/internal/ent/hostprocess"
	"realm.pub/tavern/internal/ent/task"
)

// HostProcess is the model entity for the HostProcess schema.
type HostProcess struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Timestamp of when this ent was created
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Timestamp of when this ent was last updated
	LastModifiedAt time.Time `json:"last_modified_at,omitempty"`
	// ID of the process.
	Pid uint64 `json:"pid,omitempty"`
	// ID of the parent process.
	Ppid uint64 `json:"ppid,omitempty"`
	// The name of the process.
	Name string `json:"name,omitempty"`
	// The user the process is running as.
	Principal string `json:"principal,omitempty"`
	// The path to the process executable.
	Path string `json:"path,omitempty"`
	// The command used to execute the process.
	Cmd string `json:"cmd,omitempty"`
	// The environment variables set for the process.
	Env string `json:"env,omitempty"`
	// The current working directory for the process.
	Cwd string `json:"cwd,omitempty"`
	// Current process status.
	Status c2pb.Process_Status `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HostProcessQuery when eager-loading is set.
	Edges                   HostProcessEdges `json:"edges"`
	host_processes          *int
	host_process_host       *int
	task_reported_processes *int
	selectValues            sql.SelectValues
}

// HostProcessEdges holds the relations/edges for other nodes in the graph.
type HostProcessEdges struct {
	// Host the process was reported on.
	Host *Host `json:"host,omitempty"`
	// Task that reported this process.
	Task *Task `json:"task,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int
}

// HostOrErr returns the Host value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HostProcessEdges) HostOrErr() (*Host, error) {
	if e.loadedTypes[0] {
		if e.Host == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: host.Label}
		}
		return e.Host, nil
	}
	return nil, &NotLoadedError{edge: "host"}
}

// TaskOrErr returns the Task value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e HostProcessEdges) TaskOrErr() (*Task, error) {
	if e.loadedTypes[1] {
		if e.Task == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: task.Label}
		}
		return e.Task, nil
	}
	return nil, &NotLoadedError{edge: "task"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*HostProcess) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case hostprocess.FieldStatus:
			values[i] = new(c2pb.Process_Status)
		case hostprocess.FieldID, hostprocess.FieldPid, hostprocess.FieldPpid:
			values[i] = new(sql.NullInt64)
		case hostprocess.FieldName, hostprocess.FieldPrincipal, hostprocess.FieldPath, hostprocess.FieldCmd, hostprocess.FieldEnv, hostprocess.FieldCwd:
			values[i] = new(sql.NullString)
		case hostprocess.FieldCreatedAt, hostprocess.FieldLastModifiedAt:
			values[i] = new(sql.NullTime)
		case hostprocess.ForeignKeys[0]: // host_processes
			values[i] = new(sql.NullInt64)
		case hostprocess.ForeignKeys[1]: // host_process_host
			values[i] = new(sql.NullInt64)
		case hostprocess.ForeignKeys[2]: // task_reported_processes
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the HostProcess fields.
func (hp *HostProcess) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case hostprocess.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			hp.ID = int(value.Int64)
		case hostprocess.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				hp.CreatedAt = value.Time
			}
		case hostprocess.FieldLastModifiedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_modified_at", values[i])
			} else if value.Valid {
				hp.LastModifiedAt = value.Time
			}
		case hostprocess.FieldPid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pid", values[i])
			} else if value.Valid {
				hp.Pid = uint64(value.Int64)
			}
		case hostprocess.FieldPpid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ppid", values[i])
			} else if value.Valid {
				hp.Ppid = uint64(value.Int64)
			}
		case hostprocess.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				hp.Name = value.String
			}
		case hostprocess.FieldPrincipal:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field principal", values[i])
			} else if value.Valid {
				hp.Principal = value.String
			}
		case hostprocess.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				hp.Path = value.String
			}
		case hostprocess.FieldCmd:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cmd", values[i])
			} else if value.Valid {
				hp.Cmd = value.String
			}
		case hostprocess.FieldEnv:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field env", values[i])
			} else if value.Valid {
				hp.Env = value.String
			}
		case hostprocess.FieldCwd:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cwd", values[i])
			} else if value.Valid {
				hp.Cwd = value.String
			}
		case hostprocess.FieldStatus:
			if value, ok := values[i].(*c2pb.Process_Status); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value != nil {
				hp.Status = *value
			}
		case hostprocess.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field host_processes", value)
			} else if value.Valid {
				hp.host_processes = new(int)
				*hp.host_processes = int(value.Int64)
			}
		case hostprocess.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field host_process_host", value)
			} else if value.Valid {
				hp.host_process_host = new(int)
				*hp.host_process_host = int(value.Int64)
			}
		case hostprocess.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field task_reported_processes", value)
			} else if value.Valid {
				hp.task_reported_processes = new(int)
				*hp.task_reported_processes = int(value.Int64)
			}
		default:
			hp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the HostProcess.
// This includes values selected through modifiers, order, etc.
func (hp *HostProcess) Value(name string) (ent.Value, error) {
	return hp.selectValues.Get(name)
}

// QueryHost queries the "host" edge of the HostProcess entity.
func (hp *HostProcess) QueryHost() *HostQuery {
	return NewHostProcessClient(hp.config).QueryHost(hp)
}

// QueryTask queries the "task" edge of the HostProcess entity.
func (hp *HostProcess) QueryTask() *TaskQuery {
	return NewHostProcessClient(hp.config).QueryTask(hp)
}

// Update returns a builder for updating this HostProcess.
// Note that you need to call HostProcess.Unwrap() before calling this method if this HostProcess
// was returned from a transaction, and the transaction was committed or rolled back.
func (hp *HostProcess) Update() *HostProcessUpdateOne {
	return NewHostProcessClient(hp.config).UpdateOne(hp)
}

// Unwrap unwraps the HostProcess entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (hp *HostProcess) Unwrap() *HostProcess {
	_tx, ok := hp.config.driver.(*txDriver)
	if !ok {
		panic("ent: HostProcess is not a transactional entity")
	}
	hp.config.driver = _tx.drv
	return hp
}

// String implements the fmt.Stringer.
func (hp *HostProcess) String() string {
	var builder strings.Builder
	builder.WriteString("HostProcess(")
	builder.WriteString(fmt.Sprintf("id=%v, ", hp.ID))
	builder.WriteString("created_at=")
	builder.WriteString(hp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("last_modified_at=")
	builder.WriteString(hp.LastModifiedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("pid=")
	builder.WriteString(fmt.Sprintf("%v", hp.Pid))
	builder.WriteString(", ")
	builder.WriteString("ppid=")
	builder.WriteString(fmt.Sprintf("%v", hp.Ppid))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(hp.Name)
	builder.WriteString(", ")
	builder.WriteString("principal=")
	builder.WriteString(hp.Principal)
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(hp.Path)
	builder.WriteString(", ")
	builder.WriteString("cmd=")
	builder.WriteString(hp.Cmd)
	builder.WriteString(", ")
	builder.WriteString("env=")
	builder.WriteString(hp.Env)
	builder.WriteString(", ")
	builder.WriteString("cwd=")
	builder.WriteString(hp.Cwd)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", hp.Status))
	builder.WriteByte(')')
	return builder.String()
}

// HostProcesses is a parsable slice of HostProcess.
type HostProcesses []*HostProcess

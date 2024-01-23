// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"realm.pub/tavern/internal/ent/host"
	"realm.pub/tavern/internal/ent/predicate"
	"realm.pub/tavern/internal/ent/process"
	"realm.pub/tavern/internal/ent/task"
)

// ProcessUpdate is the builder for updating Process entities.
type ProcessUpdate struct {
	config
	hooks    []Hook
	mutation *ProcessMutation
}

// Where appends a list predicates to the ProcessUpdate builder.
func (pu *ProcessUpdate) Where(ps ...predicate.Process) *ProcessUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (pu *ProcessUpdate) SetLastModifiedAt(t time.Time) *ProcessUpdate {
	pu.mutation.SetLastModifiedAt(t)
	return pu
}

// SetPid sets the "pid" field.
func (pu *ProcessUpdate) SetPid(u uint64) *ProcessUpdate {
	pu.mutation.ResetPid()
	pu.mutation.SetPid(u)
	return pu
}

// AddPid adds u to the "pid" field.
func (pu *ProcessUpdate) AddPid(u int64) *ProcessUpdate {
	pu.mutation.AddPid(u)
	return pu
}

// SetName sets the "name" field.
func (pu *ProcessUpdate) SetName(s string) *ProcessUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetPrincipal sets the "principal" field.
func (pu *ProcessUpdate) SetPrincipal(s string) *ProcessUpdate {
	pu.mutation.SetPrincipal(s)
	return pu
}

// SetHostID sets the "host" edge to the Host entity by ID.
func (pu *ProcessUpdate) SetHostID(id int) *ProcessUpdate {
	pu.mutation.SetHostID(id)
	return pu
}

// SetHost sets the "host" edge to the Host entity.
func (pu *ProcessUpdate) SetHost(h *Host) *ProcessUpdate {
	return pu.SetHostID(h.ID)
}

// SetTaskID sets the "task" edge to the Task entity by ID.
func (pu *ProcessUpdate) SetTaskID(id int) *ProcessUpdate {
	pu.mutation.SetTaskID(id)
	return pu
}

// SetTask sets the "task" edge to the Task entity.
func (pu *ProcessUpdate) SetTask(t *Task) *ProcessUpdate {
	return pu.SetTaskID(t.ID)
}

// Mutation returns the ProcessMutation object of the builder.
func (pu *ProcessUpdate) Mutation() *ProcessMutation {
	return pu.mutation
}

// ClearHost clears the "host" edge to the Host entity.
func (pu *ProcessUpdate) ClearHost() *ProcessUpdate {
	pu.mutation.ClearHost()
	return pu
}

// ClearTask clears the "task" edge to the Task entity.
func (pu *ProcessUpdate) ClearTask() *ProcessUpdate {
	pu.mutation.ClearTask()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProcessUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProcessUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProcessUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProcessUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *ProcessUpdate) defaults() {
	if _, ok := pu.mutation.LastModifiedAt(); !ok {
		v := process.UpdateDefaultLastModifiedAt()
		pu.mutation.SetLastModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *ProcessUpdate) check() error {
	if v, ok := pu.mutation.Principal(); ok {
		if err := process.PrincipalValidator(v); err != nil {
			return &ValidationError{Name: "principal", err: fmt.Errorf(`ent: validator failed for field "Process.principal": %w`, err)}
		}
	}
	if _, ok := pu.mutation.HostID(); pu.mutation.HostCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Process.host"`)
	}
	if _, ok := pu.mutation.TaskID(); pu.mutation.TaskCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Process.task"`)
	}
	return nil
}

func (pu *ProcessUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(process.Table, process.Columns, sqlgraph.NewFieldSpec(process.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.LastModifiedAt(); ok {
		_spec.SetField(process.FieldLastModifiedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.Pid(); ok {
		_spec.SetField(process.FieldPid, field.TypeUint64, value)
	}
	if value, ok := pu.mutation.AddedPid(); ok {
		_spec.AddField(process.FieldPid, field.TypeUint64, value)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(process.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Principal(); ok {
		_spec.SetField(process.FieldPrincipal, field.TypeString, value)
	}
	if pu.mutation.HostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   process.HostTable,
			Columns: []string{process.HostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.HostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   process.HostTable,
			Columns: []string{process.HostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   process.TaskTable,
			Columns: []string{process.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   process.TaskTable,
			Columns: []string{process.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{process.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProcessUpdateOne is the builder for updating a single Process entity.
type ProcessUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProcessMutation
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (puo *ProcessUpdateOne) SetLastModifiedAt(t time.Time) *ProcessUpdateOne {
	puo.mutation.SetLastModifiedAt(t)
	return puo
}

// SetPid sets the "pid" field.
func (puo *ProcessUpdateOne) SetPid(u uint64) *ProcessUpdateOne {
	puo.mutation.ResetPid()
	puo.mutation.SetPid(u)
	return puo
}

// AddPid adds u to the "pid" field.
func (puo *ProcessUpdateOne) AddPid(u int64) *ProcessUpdateOne {
	puo.mutation.AddPid(u)
	return puo
}

// SetName sets the "name" field.
func (puo *ProcessUpdateOne) SetName(s string) *ProcessUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetPrincipal sets the "principal" field.
func (puo *ProcessUpdateOne) SetPrincipal(s string) *ProcessUpdateOne {
	puo.mutation.SetPrincipal(s)
	return puo
}

// SetHostID sets the "host" edge to the Host entity by ID.
func (puo *ProcessUpdateOne) SetHostID(id int) *ProcessUpdateOne {
	puo.mutation.SetHostID(id)
	return puo
}

// SetHost sets the "host" edge to the Host entity.
func (puo *ProcessUpdateOne) SetHost(h *Host) *ProcessUpdateOne {
	return puo.SetHostID(h.ID)
}

// SetTaskID sets the "task" edge to the Task entity by ID.
func (puo *ProcessUpdateOne) SetTaskID(id int) *ProcessUpdateOne {
	puo.mutation.SetTaskID(id)
	return puo
}

// SetTask sets the "task" edge to the Task entity.
func (puo *ProcessUpdateOne) SetTask(t *Task) *ProcessUpdateOne {
	return puo.SetTaskID(t.ID)
}

// Mutation returns the ProcessMutation object of the builder.
func (puo *ProcessUpdateOne) Mutation() *ProcessMutation {
	return puo.mutation
}

// ClearHost clears the "host" edge to the Host entity.
func (puo *ProcessUpdateOne) ClearHost() *ProcessUpdateOne {
	puo.mutation.ClearHost()
	return puo
}

// ClearTask clears the "task" edge to the Task entity.
func (puo *ProcessUpdateOne) ClearTask() *ProcessUpdateOne {
	puo.mutation.ClearTask()
	return puo
}

// Where appends a list predicates to the ProcessUpdate builder.
func (puo *ProcessUpdateOne) Where(ps ...predicate.Process) *ProcessUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProcessUpdateOne) Select(field string, fields ...string) *ProcessUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Process entity.
func (puo *ProcessUpdateOne) Save(ctx context.Context) (*Process, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProcessUpdateOne) SaveX(ctx context.Context) *Process {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProcessUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProcessUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *ProcessUpdateOne) defaults() {
	if _, ok := puo.mutation.LastModifiedAt(); !ok {
		v := process.UpdateDefaultLastModifiedAt()
		puo.mutation.SetLastModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *ProcessUpdateOne) check() error {
	if v, ok := puo.mutation.Principal(); ok {
		if err := process.PrincipalValidator(v); err != nil {
			return &ValidationError{Name: "principal", err: fmt.Errorf(`ent: validator failed for field "Process.principal": %w`, err)}
		}
	}
	if _, ok := puo.mutation.HostID(); puo.mutation.HostCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Process.host"`)
	}
	if _, ok := puo.mutation.TaskID(); puo.mutation.TaskCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Process.task"`)
	}
	return nil
}

func (puo *ProcessUpdateOne) sqlSave(ctx context.Context) (_node *Process, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(process.Table, process.Columns, sqlgraph.NewFieldSpec(process.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Process.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, process.FieldID)
		for _, f := range fields {
			if !process.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != process.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.LastModifiedAt(); ok {
		_spec.SetField(process.FieldLastModifiedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.Pid(); ok {
		_spec.SetField(process.FieldPid, field.TypeUint64, value)
	}
	if value, ok := puo.mutation.AddedPid(); ok {
		_spec.AddField(process.FieldPid, field.TypeUint64, value)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(process.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Principal(); ok {
		_spec.SetField(process.FieldPrincipal, field.TypeString, value)
	}
	if puo.mutation.HostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   process.HostTable,
			Columns: []string{process.HostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.HostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   process.HostTable,
			Columns: []string{process.HostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   process.TaskTable,
			Columns: []string{process.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   process.TaskTable,
			Columns: []string{process.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Process{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{process.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}

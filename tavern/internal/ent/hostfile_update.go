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
	"realm.pub/tavern/internal/ent/hostfile"
	"realm.pub/tavern/internal/ent/predicate"
	"realm.pub/tavern/internal/ent/task"
)

// HostFileUpdate is the builder for updating HostFile entities.
type HostFileUpdate struct {
	config
	hooks    []Hook
	mutation *HostFileMutation
}

// Where appends a list predicates to the HostFileUpdate builder.
func (hfu *HostFileUpdate) Where(ps ...predicate.HostFile) *HostFileUpdate {
	hfu.mutation.Where(ps...)
	return hfu
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (hfu *HostFileUpdate) SetLastModifiedAt(t time.Time) *HostFileUpdate {
	hfu.mutation.SetLastModifiedAt(t)
	return hfu
}

// SetPath sets the "path" field.
func (hfu *HostFileUpdate) SetPath(s string) *HostFileUpdate {
	hfu.mutation.SetPath(s)
	return hfu
}

// SetOwner sets the "owner" field.
func (hfu *HostFileUpdate) SetOwner(s string) *HostFileUpdate {
	hfu.mutation.SetOwner(s)
	return hfu
}

// SetNillableOwner sets the "owner" field if the given value is not nil.
func (hfu *HostFileUpdate) SetNillableOwner(s *string) *HostFileUpdate {
	if s != nil {
		hfu.SetOwner(*s)
	}
	return hfu
}

// ClearOwner clears the value of the "owner" field.
func (hfu *HostFileUpdate) ClearOwner() *HostFileUpdate {
	hfu.mutation.ClearOwner()
	return hfu
}

// SetGroup sets the "group" field.
func (hfu *HostFileUpdate) SetGroup(s string) *HostFileUpdate {
	hfu.mutation.SetGroup(s)
	return hfu
}

// SetNillableGroup sets the "group" field if the given value is not nil.
func (hfu *HostFileUpdate) SetNillableGroup(s *string) *HostFileUpdate {
	if s != nil {
		hfu.SetGroup(*s)
	}
	return hfu
}

// ClearGroup clears the value of the "group" field.
func (hfu *HostFileUpdate) ClearGroup() *HostFileUpdate {
	hfu.mutation.ClearGroup()
	return hfu
}

// SetPermissions sets the "permissions" field.
func (hfu *HostFileUpdate) SetPermissions(s string) *HostFileUpdate {
	hfu.mutation.SetPermissions(s)
	return hfu
}

// SetNillablePermissions sets the "permissions" field if the given value is not nil.
func (hfu *HostFileUpdate) SetNillablePermissions(s *string) *HostFileUpdate {
	if s != nil {
		hfu.SetPermissions(*s)
	}
	return hfu
}

// ClearPermissions clears the value of the "permissions" field.
func (hfu *HostFileUpdate) ClearPermissions() *HostFileUpdate {
	hfu.mutation.ClearPermissions()
	return hfu
}

// SetSize sets the "size" field.
func (hfu *HostFileUpdate) SetSize(i int) *HostFileUpdate {
	hfu.mutation.ResetSize()
	hfu.mutation.SetSize(i)
	return hfu
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (hfu *HostFileUpdate) SetNillableSize(i *int) *HostFileUpdate {
	if i != nil {
		hfu.SetSize(*i)
	}
	return hfu
}

// AddSize adds i to the "size" field.
func (hfu *HostFileUpdate) AddSize(i int) *HostFileUpdate {
	hfu.mutation.AddSize(i)
	return hfu
}

// SetHash sets the "hash" field.
func (hfu *HostFileUpdate) SetHash(s string) *HostFileUpdate {
	hfu.mutation.SetHash(s)
	return hfu
}

// SetNillableHash sets the "hash" field if the given value is not nil.
func (hfu *HostFileUpdate) SetNillableHash(s *string) *HostFileUpdate {
	if s != nil {
		hfu.SetHash(*s)
	}
	return hfu
}

// ClearHash clears the value of the "hash" field.
func (hfu *HostFileUpdate) ClearHash() *HostFileUpdate {
	hfu.mutation.ClearHash()
	return hfu
}

// SetContent sets the "content" field.
func (hfu *HostFileUpdate) SetContent(b []byte) *HostFileUpdate {
	hfu.mutation.SetContent(b)
	return hfu
}

// ClearContent clears the value of the "content" field.
func (hfu *HostFileUpdate) ClearContent() *HostFileUpdate {
	hfu.mutation.ClearContent()
	return hfu
}

// SetHostID sets the "host" edge to the Host entity by ID.
func (hfu *HostFileUpdate) SetHostID(id int) *HostFileUpdate {
	hfu.mutation.SetHostID(id)
	return hfu
}

// SetHost sets the "host" edge to the Host entity.
func (hfu *HostFileUpdate) SetHost(h *Host) *HostFileUpdate {
	return hfu.SetHostID(h.ID)
}

// SetTaskID sets the "task" edge to the Task entity by ID.
func (hfu *HostFileUpdate) SetTaskID(id int) *HostFileUpdate {
	hfu.mutation.SetTaskID(id)
	return hfu
}

// SetTask sets the "task" edge to the Task entity.
func (hfu *HostFileUpdate) SetTask(t *Task) *HostFileUpdate {
	return hfu.SetTaskID(t.ID)
}

// Mutation returns the HostFileMutation object of the builder.
func (hfu *HostFileUpdate) Mutation() *HostFileMutation {
	return hfu.mutation
}

// ClearHost clears the "host" edge to the Host entity.
func (hfu *HostFileUpdate) ClearHost() *HostFileUpdate {
	hfu.mutation.ClearHost()
	return hfu
}

// ClearTask clears the "task" edge to the Task entity.
func (hfu *HostFileUpdate) ClearTask() *HostFileUpdate {
	hfu.mutation.ClearTask()
	return hfu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hfu *HostFileUpdate) Save(ctx context.Context) (int, error) {
	if err := hfu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, hfu.sqlSave, hfu.mutation, hfu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hfu *HostFileUpdate) SaveX(ctx context.Context) int {
	affected, err := hfu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hfu *HostFileUpdate) Exec(ctx context.Context) error {
	_, err := hfu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hfu *HostFileUpdate) ExecX(ctx context.Context) {
	if err := hfu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hfu *HostFileUpdate) defaults() error {
	if _, ok := hfu.mutation.LastModifiedAt(); !ok {
		if hostfile.UpdateDefaultLastModifiedAt == nil {
			return fmt.Errorf("ent: uninitialized hostfile.UpdateDefaultLastModifiedAt (forgotten import ent/runtime?)")
		}
		v := hostfile.UpdateDefaultLastModifiedAt()
		hfu.mutation.SetLastModifiedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (hfu *HostFileUpdate) check() error {
	if v, ok := hfu.mutation.Path(); ok {
		if err := hostfile.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "HostFile.path": %w`, err)}
		}
	}
	if v, ok := hfu.mutation.Size(); ok {
		if err := hostfile.SizeValidator(v); err != nil {
			return &ValidationError{Name: "size", err: fmt.Errorf(`ent: validator failed for field "HostFile.size": %w`, err)}
		}
	}
	if v, ok := hfu.mutation.Hash(); ok {
		if err := hostfile.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf(`ent: validator failed for field "HostFile.hash": %w`, err)}
		}
	}
	if _, ok := hfu.mutation.HostID(); hfu.mutation.HostCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "HostFile.host"`)
	}
	if _, ok := hfu.mutation.TaskID(); hfu.mutation.TaskCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "HostFile.task"`)
	}
	return nil
}

func (hfu *HostFileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := hfu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(hostfile.Table, hostfile.Columns, sqlgraph.NewFieldSpec(hostfile.FieldID, field.TypeInt))
	if ps := hfu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hfu.mutation.LastModifiedAt(); ok {
		_spec.SetField(hostfile.FieldLastModifiedAt, field.TypeTime, value)
	}
	if value, ok := hfu.mutation.Path(); ok {
		_spec.SetField(hostfile.FieldPath, field.TypeString, value)
	}
	if value, ok := hfu.mutation.Owner(); ok {
		_spec.SetField(hostfile.FieldOwner, field.TypeString, value)
	}
	if hfu.mutation.OwnerCleared() {
		_spec.ClearField(hostfile.FieldOwner, field.TypeString)
	}
	if value, ok := hfu.mutation.Group(); ok {
		_spec.SetField(hostfile.FieldGroup, field.TypeString, value)
	}
	if hfu.mutation.GroupCleared() {
		_spec.ClearField(hostfile.FieldGroup, field.TypeString)
	}
	if value, ok := hfu.mutation.Permissions(); ok {
		_spec.SetField(hostfile.FieldPermissions, field.TypeString, value)
	}
	if hfu.mutation.PermissionsCleared() {
		_spec.ClearField(hostfile.FieldPermissions, field.TypeString)
	}
	if value, ok := hfu.mutation.Size(); ok {
		_spec.SetField(hostfile.FieldSize, field.TypeInt, value)
	}
	if value, ok := hfu.mutation.AddedSize(); ok {
		_spec.AddField(hostfile.FieldSize, field.TypeInt, value)
	}
	if value, ok := hfu.mutation.Hash(); ok {
		_spec.SetField(hostfile.FieldHash, field.TypeString, value)
	}
	if hfu.mutation.HashCleared() {
		_spec.ClearField(hostfile.FieldHash, field.TypeString)
	}
	if value, ok := hfu.mutation.Content(); ok {
		_spec.SetField(hostfile.FieldContent, field.TypeBytes, value)
	}
	if hfu.mutation.ContentCleared() {
		_spec.ClearField(hostfile.FieldContent, field.TypeBytes)
	}
	if hfu.mutation.HostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hostfile.HostTable,
			Columns: []string{hostfile.HostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hfu.mutation.HostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hostfile.HostTable,
			Columns: []string{hostfile.HostColumn},
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
	if hfu.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hostfile.TaskTable,
			Columns: []string{hostfile.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hfu.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hostfile.TaskTable,
			Columns: []string{hostfile.TaskColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, hfu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hostfile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hfu.mutation.done = true
	return n, nil
}

// HostFileUpdateOne is the builder for updating a single HostFile entity.
type HostFileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HostFileMutation
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (hfuo *HostFileUpdateOne) SetLastModifiedAt(t time.Time) *HostFileUpdateOne {
	hfuo.mutation.SetLastModifiedAt(t)
	return hfuo
}

// SetPath sets the "path" field.
func (hfuo *HostFileUpdateOne) SetPath(s string) *HostFileUpdateOne {
	hfuo.mutation.SetPath(s)
	return hfuo
}

// SetOwner sets the "owner" field.
func (hfuo *HostFileUpdateOne) SetOwner(s string) *HostFileUpdateOne {
	hfuo.mutation.SetOwner(s)
	return hfuo
}

// SetNillableOwner sets the "owner" field if the given value is not nil.
func (hfuo *HostFileUpdateOne) SetNillableOwner(s *string) *HostFileUpdateOne {
	if s != nil {
		hfuo.SetOwner(*s)
	}
	return hfuo
}

// ClearOwner clears the value of the "owner" field.
func (hfuo *HostFileUpdateOne) ClearOwner() *HostFileUpdateOne {
	hfuo.mutation.ClearOwner()
	return hfuo
}

// SetGroup sets the "group" field.
func (hfuo *HostFileUpdateOne) SetGroup(s string) *HostFileUpdateOne {
	hfuo.mutation.SetGroup(s)
	return hfuo
}

// SetNillableGroup sets the "group" field if the given value is not nil.
func (hfuo *HostFileUpdateOne) SetNillableGroup(s *string) *HostFileUpdateOne {
	if s != nil {
		hfuo.SetGroup(*s)
	}
	return hfuo
}

// ClearGroup clears the value of the "group" field.
func (hfuo *HostFileUpdateOne) ClearGroup() *HostFileUpdateOne {
	hfuo.mutation.ClearGroup()
	return hfuo
}

// SetPermissions sets the "permissions" field.
func (hfuo *HostFileUpdateOne) SetPermissions(s string) *HostFileUpdateOne {
	hfuo.mutation.SetPermissions(s)
	return hfuo
}

// SetNillablePermissions sets the "permissions" field if the given value is not nil.
func (hfuo *HostFileUpdateOne) SetNillablePermissions(s *string) *HostFileUpdateOne {
	if s != nil {
		hfuo.SetPermissions(*s)
	}
	return hfuo
}

// ClearPermissions clears the value of the "permissions" field.
func (hfuo *HostFileUpdateOne) ClearPermissions() *HostFileUpdateOne {
	hfuo.mutation.ClearPermissions()
	return hfuo
}

// SetSize sets the "size" field.
func (hfuo *HostFileUpdateOne) SetSize(i int) *HostFileUpdateOne {
	hfuo.mutation.ResetSize()
	hfuo.mutation.SetSize(i)
	return hfuo
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (hfuo *HostFileUpdateOne) SetNillableSize(i *int) *HostFileUpdateOne {
	if i != nil {
		hfuo.SetSize(*i)
	}
	return hfuo
}

// AddSize adds i to the "size" field.
func (hfuo *HostFileUpdateOne) AddSize(i int) *HostFileUpdateOne {
	hfuo.mutation.AddSize(i)
	return hfuo
}

// SetHash sets the "hash" field.
func (hfuo *HostFileUpdateOne) SetHash(s string) *HostFileUpdateOne {
	hfuo.mutation.SetHash(s)
	return hfuo
}

// SetNillableHash sets the "hash" field if the given value is not nil.
func (hfuo *HostFileUpdateOne) SetNillableHash(s *string) *HostFileUpdateOne {
	if s != nil {
		hfuo.SetHash(*s)
	}
	return hfuo
}

// ClearHash clears the value of the "hash" field.
func (hfuo *HostFileUpdateOne) ClearHash() *HostFileUpdateOne {
	hfuo.mutation.ClearHash()
	return hfuo
}

// SetContent sets the "content" field.
func (hfuo *HostFileUpdateOne) SetContent(b []byte) *HostFileUpdateOne {
	hfuo.mutation.SetContent(b)
	return hfuo
}

// ClearContent clears the value of the "content" field.
func (hfuo *HostFileUpdateOne) ClearContent() *HostFileUpdateOne {
	hfuo.mutation.ClearContent()
	return hfuo
}

// SetHostID sets the "host" edge to the Host entity by ID.
func (hfuo *HostFileUpdateOne) SetHostID(id int) *HostFileUpdateOne {
	hfuo.mutation.SetHostID(id)
	return hfuo
}

// SetHost sets the "host" edge to the Host entity.
func (hfuo *HostFileUpdateOne) SetHost(h *Host) *HostFileUpdateOne {
	return hfuo.SetHostID(h.ID)
}

// SetTaskID sets the "task" edge to the Task entity by ID.
func (hfuo *HostFileUpdateOne) SetTaskID(id int) *HostFileUpdateOne {
	hfuo.mutation.SetTaskID(id)
	return hfuo
}

// SetTask sets the "task" edge to the Task entity.
func (hfuo *HostFileUpdateOne) SetTask(t *Task) *HostFileUpdateOne {
	return hfuo.SetTaskID(t.ID)
}

// Mutation returns the HostFileMutation object of the builder.
func (hfuo *HostFileUpdateOne) Mutation() *HostFileMutation {
	return hfuo.mutation
}

// ClearHost clears the "host" edge to the Host entity.
func (hfuo *HostFileUpdateOne) ClearHost() *HostFileUpdateOne {
	hfuo.mutation.ClearHost()
	return hfuo
}

// ClearTask clears the "task" edge to the Task entity.
func (hfuo *HostFileUpdateOne) ClearTask() *HostFileUpdateOne {
	hfuo.mutation.ClearTask()
	return hfuo
}

// Where appends a list predicates to the HostFileUpdate builder.
func (hfuo *HostFileUpdateOne) Where(ps ...predicate.HostFile) *HostFileUpdateOne {
	hfuo.mutation.Where(ps...)
	return hfuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (hfuo *HostFileUpdateOne) Select(field string, fields ...string) *HostFileUpdateOne {
	hfuo.fields = append([]string{field}, fields...)
	return hfuo
}

// Save executes the query and returns the updated HostFile entity.
func (hfuo *HostFileUpdateOne) Save(ctx context.Context) (*HostFile, error) {
	if err := hfuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, hfuo.sqlSave, hfuo.mutation, hfuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hfuo *HostFileUpdateOne) SaveX(ctx context.Context) *HostFile {
	node, err := hfuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (hfuo *HostFileUpdateOne) Exec(ctx context.Context) error {
	_, err := hfuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hfuo *HostFileUpdateOne) ExecX(ctx context.Context) {
	if err := hfuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hfuo *HostFileUpdateOne) defaults() error {
	if _, ok := hfuo.mutation.LastModifiedAt(); !ok {
		if hostfile.UpdateDefaultLastModifiedAt == nil {
			return fmt.Errorf("ent: uninitialized hostfile.UpdateDefaultLastModifiedAt (forgotten import ent/runtime?)")
		}
		v := hostfile.UpdateDefaultLastModifiedAt()
		hfuo.mutation.SetLastModifiedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (hfuo *HostFileUpdateOne) check() error {
	if v, ok := hfuo.mutation.Path(); ok {
		if err := hostfile.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "HostFile.path": %w`, err)}
		}
	}
	if v, ok := hfuo.mutation.Size(); ok {
		if err := hostfile.SizeValidator(v); err != nil {
			return &ValidationError{Name: "size", err: fmt.Errorf(`ent: validator failed for field "HostFile.size": %w`, err)}
		}
	}
	if v, ok := hfuo.mutation.Hash(); ok {
		if err := hostfile.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf(`ent: validator failed for field "HostFile.hash": %w`, err)}
		}
	}
	if _, ok := hfuo.mutation.HostID(); hfuo.mutation.HostCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "HostFile.host"`)
	}
	if _, ok := hfuo.mutation.TaskID(); hfuo.mutation.TaskCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "HostFile.task"`)
	}
	return nil
}

func (hfuo *HostFileUpdateOne) sqlSave(ctx context.Context) (_node *HostFile, err error) {
	if err := hfuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(hostfile.Table, hostfile.Columns, sqlgraph.NewFieldSpec(hostfile.FieldID, field.TypeInt))
	id, ok := hfuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "HostFile.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := hfuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hostfile.FieldID)
		for _, f := range fields {
			if !hostfile.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != hostfile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := hfuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hfuo.mutation.LastModifiedAt(); ok {
		_spec.SetField(hostfile.FieldLastModifiedAt, field.TypeTime, value)
	}
	if value, ok := hfuo.mutation.Path(); ok {
		_spec.SetField(hostfile.FieldPath, field.TypeString, value)
	}
	if value, ok := hfuo.mutation.Owner(); ok {
		_spec.SetField(hostfile.FieldOwner, field.TypeString, value)
	}
	if hfuo.mutation.OwnerCleared() {
		_spec.ClearField(hostfile.FieldOwner, field.TypeString)
	}
	if value, ok := hfuo.mutation.Group(); ok {
		_spec.SetField(hostfile.FieldGroup, field.TypeString, value)
	}
	if hfuo.mutation.GroupCleared() {
		_spec.ClearField(hostfile.FieldGroup, field.TypeString)
	}
	if value, ok := hfuo.mutation.Permissions(); ok {
		_spec.SetField(hostfile.FieldPermissions, field.TypeString, value)
	}
	if hfuo.mutation.PermissionsCleared() {
		_spec.ClearField(hostfile.FieldPermissions, field.TypeString)
	}
	if value, ok := hfuo.mutation.Size(); ok {
		_spec.SetField(hostfile.FieldSize, field.TypeInt, value)
	}
	if value, ok := hfuo.mutation.AddedSize(); ok {
		_spec.AddField(hostfile.FieldSize, field.TypeInt, value)
	}
	if value, ok := hfuo.mutation.Hash(); ok {
		_spec.SetField(hostfile.FieldHash, field.TypeString, value)
	}
	if hfuo.mutation.HashCleared() {
		_spec.ClearField(hostfile.FieldHash, field.TypeString)
	}
	if value, ok := hfuo.mutation.Content(); ok {
		_spec.SetField(hostfile.FieldContent, field.TypeBytes, value)
	}
	if hfuo.mutation.ContentCleared() {
		_spec.ClearField(hostfile.FieldContent, field.TypeBytes)
	}
	if hfuo.mutation.HostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hostfile.HostTable,
			Columns: []string{hostfile.HostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hfuo.mutation.HostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   hostfile.HostTable,
			Columns: []string{hostfile.HostColumn},
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
	if hfuo.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hostfile.TaskTable,
			Columns: []string{hostfile.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := hfuo.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   hostfile.TaskTable,
			Columns: []string{hostfile.TaskColumn},
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
	_node = &HostFile{config: hfuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, hfuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hostfile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	hfuo.mutation.done = true
	return _node, nil
}

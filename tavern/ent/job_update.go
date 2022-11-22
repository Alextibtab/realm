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
	"github.com/kcarretto/realm/tavern/ent/file"
	"github.com/kcarretto/realm/tavern/ent/job"
	"github.com/kcarretto/realm/tavern/ent/predicate"
	"github.com/kcarretto/realm/tavern/ent/task"
	"github.com/kcarretto/realm/tavern/ent/tome"
)

// JobUpdate is the builder for updating Job entities.
type JobUpdate struct {
	config
	hooks    []Hook
	mutation *JobMutation
}

// Where appends a list predicates to the JobUpdate builder.
func (ju *JobUpdate) Where(ps ...predicate.Job) *JobUpdate {
	ju.mutation.Where(ps...)
	return ju
}

// SetLastModifiedAt sets the "lastModifiedAt" field.
func (ju *JobUpdate) SetLastModifiedAt(t time.Time) *JobUpdate {
	ju.mutation.SetLastModifiedAt(t)
	return ju
}

// SetName sets the "name" field.
func (ju *JobUpdate) SetName(s string) *JobUpdate {
	ju.mutation.SetName(s)
	return ju
}

// SetTomeID sets the "tome" edge to the Tome entity by ID.
func (ju *JobUpdate) SetTomeID(id int) *JobUpdate {
	ju.mutation.SetTomeID(id)
	return ju
}

// SetTome sets the "tome" edge to the Tome entity.
func (ju *JobUpdate) SetTome(t *Tome) *JobUpdate {
	return ju.SetTomeID(t.ID)
}

// SetBundleID sets the "bundle" edge to the File entity by ID.
func (ju *JobUpdate) SetBundleID(id int) *JobUpdate {
	ju.mutation.SetBundleID(id)
	return ju
}

// SetNillableBundleID sets the "bundle" edge to the File entity by ID if the given value is not nil.
func (ju *JobUpdate) SetNillableBundleID(id *int) *JobUpdate {
	if id != nil {
		ju = ju.SetBundleID(*id)
	}
	return ju
}

// SetBundle sets the "bundle" edge to the File entity.
func (ju *JobUpdate) SetBundle(f *File) *JobUpdate {
	return ju.SetBundleID(f.ID)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (ju *JobUpdate) AddTaskIDs(ids ...int) *JobUpdate {
	ju.mutation.AddTaskIDs(ids...)
	return ju
}

// AddTasks adds the "tasks" edges to the Task entity.
func (ju *JobUpdate) AddTasks(t ...*Task) *JobUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ju.AddTaskIDs(ids...)
}

// Mutation returns the JobMutation object of the builder.
func (ju *JobUpdate) Mutation() *JobMutation {
	return ju.mutation
}

// ClearTome clears the "tome" edge to the Tome entity.
func (ju *JobUpdate) ClearTome() *JobUpdate {
	ju.mutation.ClearTome()
	return ju
}

// ClearBundle clears the "bundle" edge to the File entity.
func (ju *JobUpdate) ClearBundle() *JobUpdate {
	ju.mutation.ClearBundle()
	return ju
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (ju *JobUpdate) ClearTasks() *JobUpdate {
	ju.mutation.ClearTasks()
	return ju
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (ju *JobUpdate) RemoveTaskIDs(ids ...int) *JobUpdate {
	ju.mutation.RemoveTaskIDs(ids...)
	return ju
}

// RemoveTasks removes "tasks" edges to Task entities.
func (ju *JobUpdate) RemoveTasks(t ...*Task) *JobUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ju.RemoveTaskIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ju *JobUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ju.defaults()
	if len(ju.hooks) == 0 {
		if err = ju.check(); err != nil {
			return 0, err
		}
		affected, err = ju.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*JobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ju.check(); err != nil {
				return 0, err
			}
			ju.mutation = mutation
			affected, err = ju.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ju.hooks) - 1; i >= 0; i-- {
			if ju.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ju.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ju.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ju *JobUpdate) SaveX(ctx context.Context) int {
	affected, err := ju.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ju *JobUpdate) Exec(ctx context.Context) error {
	_, err := ju.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ju *JobUpdate) ExecX(ctx context.Context) {
	if err := ju.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ju *JobUpdate) defaults() {
	if _, ok := ju.mutation.LastModifiedAt(); !ok {
		v := job.UpdateDefaultLastModifiedAt()
		ju.mutation.SetLastModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ju *JobUpdate) check() error {
	if v, ok := ju.mutation.Name(); ok {
		if err := job.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Job.name": %w`, err)}
		}
	}
	if _, ok := ju.mutation.TomeID(); ju.mutation.TomeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Job.tome"`)
	}
	return nil
}

func (ju *JobUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   job.Table,
			Columns: job.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: job.FieldID,
			},
		},
	}
	if ps := ju.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ju.mutation.LastModifiedAt(); ok {
		_spec.SetField(job.FieldLastModifiedAt, field.TypeTime, value)
	}
	if value, ok := ju.mutation.Name(); ok {
		_spec.SetField(job.FieldName, field.TypeString, value)
	}
	if ju.mutation.TomeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   job.TomeTable,
			Columns: []string{job.TomeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tome.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ju.mutation.TomeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   job.TomeTable,
			Columns: []string{job.TomeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tome.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ju.mutation.BundleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   job.BundleTable,
			Columns: []string{job.BundleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ju.mutation.BundleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   job.BundleTable,
			Columns: []string{job.BundleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ju.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   job.TasksTable,
			Columns: []string{job.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ju.mutation.RemovedTasksIDs(); len(nodes) > 0 && !ju.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   job.TasksTable,
			Columns: []string{job.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ju.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   job.TasksTable,
			Columns: []string{job.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ju.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{job.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// JobUpdateOne is the builder for updating a single Job entity.
type JobUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *JobMutation
}

// SetLastModifiedAt sets the "lastModifiedAt" field.
func (juo *JobUpdateOne) SetLastModifiedAt(t time.Time) *JobUpdateOne {
	juo.mutation.SetLastModifiedAt(t)
	return juo
}

// SetName sets the "name" field.
func (juo *JobUpdateOne) SetName(s string) *JobUpdateOne {
	juo.mutation.SetName(s)
	return juo
}

// SetTomeID sets the "tome" edge to the Tome entity by ID.
func (juo *JobUpdateOne) SetTomeID(id int) *JobUpdateOne {
	juo.mutation.SetTomeID(id)
	return juo
}

// SetTome sets the "tome" edge to the Tome entity.
func (juo *JobUpdateOne) SetTome(t *Tome) *JobUpdateOne {
	return juo.SetTomeID(t.ID)
}

// SetBundleID sets the "bundle" edge to the File entity by ID.
func (juo *JobUpdateOne) SetBundleID(id int) *JobUpdateOne {
	juo.mutation.SetBundleID(id)
	return juo
}

// SetNillableBundleID sets the "bundle" edge to the File entity by ID if the given value is not nil.
func (juo *JobUpdateOne) SetNillableBundleID(id *int) *JobUpdateOne {
	if id != nil {
		juo = juo.SetBundleID(*id)
	}
	return juo
}

// SetBundle sets the "bundle" edge to the File entity.
func (juo *JobUpdateOne) SetBundle(f *File) *JobUpdateOne {
	return juo.SetBundleID(f.ID)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (juo *JobUpdateOne) AddTaskIDs(ids ...int) *JobUpdateOne {
	juo.mutation.AddTaskIDs(ids...)
	return juo
}

// AddTasks adds the "tasks" edges to the Task entity.
func (juo *JobUpdateOne) AddTasks(t ...*Task) *JobUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return juo.AddTaskIDs(ids...)
}

// Mutation returns the JobMutation object of the builder.
func (juo *JobUpdateOne) Mutation() *JobMutation {
	return juo.mutation
}

// ClearTome clears the "tome" edge to the Tome entity.
func (juo *JobUpdateOne) ClearTome() *JobUpdateOne {
	juo.mutation.ClearTome()
	return juo
}

// ClearBundle clears the "bundle" edge to the File entity.
func (juo *JobUpdateOne) ClearBundle() *JobUpdateOne {
	juo.mutation.ClearBundle()
	return juo
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (juo *JobUpdateOne) ClearTasks() *JobUpdateOne {
	juo.mutation.ClearTasks()
	return juo
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (juo *JobUpdateOne) RemoveTaskIDs(ids ...int) *JobUpdateOne {
	juo.mutation.RemoveTaskIDs(ids...)
	return juo
}

// RemoveTasks removes "tasks" edges to Task entities.
func (juo *JobUpdateOne) RemoveTasks(t ...*Task) *JobUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return juo.RemoveTaskIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (juo *JobUpdateOne) Select(field string, fields ...string) *JobUpdateOne {
	juo.fields = append([]string{field}, fields...)
	return juo
}

// Save executes the query and returns the updated Job entity.
func (juo *JobUpdateOne) Save(ctx context.Context) (*Job, error) {
	var (
		err  error
		node *Job
	)
	juo.defaults()
	if len(juo.hooks) == 0 {
		if err = juo.check(); err != nil {
			return nil, err
		}
		node, err = juo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*JobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = juo.check(); err != nil {
				return nil, err
			}
			juo.mutation = mutation
			node, err = juo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(juo.hooks) - 1; i >= 0; i-- {
			if juo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = juo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, juo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Job)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from JobMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (juo *JobUpdateOne) SaveX(ctx context.Context) *Job {
	node, err := juo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (juo *JobUpdateOne) Exec(ctx context.Context) error {
	_, err := juo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (juo *JobUpdateOne) ExecX(ctx context.Context) {
	if err := juo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (juo *JobUpdateOne) defaults() {
	if _, ok := juo.mutation.LastModifiedAt(); !ok {
		v := job.UpdateDefaultLastModifiedAt()
		juo.mutation.SetLastModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (juo *JobUpdateOne) check() error {
	if v, ok := juo.mutation.Name(); ok {
		if err := job.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Job.name": %w`, err)}
		}
	}
	if _, ok := juo.mutation.TomeID(); juo.mutation.TomeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Job.tome"`)
	}
	return nil
}

func (juo *JobUpdateOne) sqlSave(ctx context.Context) (_node *Job, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   job.Table,
			Columns: job.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: job.FieldID,
			},
		},
	}
	id, ok := juo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Job.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := juo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, job.FieldID)
		for _, f := range fields {
			if !job.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != job.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := juo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := juo.mutation.LastModifiedAt(); ok {
		_spec.SetField(job.FieldLastModifiedAt, field.TypeTime, value)
	}
	if value, ok := juo.mutation.Name(); ok {
		_spec.SetField(job.FieldName, field.TypeString, value)
	}
	if juo.mutation.TomeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   job.TomeTable,
			Columns: []string{job.TomeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tome.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := juo.mutation.TomeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   job.TomeTable,
			Columns: []string{job.TomeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tome.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if juo.mutation.BundleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   job.BundleTable,
			Columns: []string{job.BundleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := juo.mutation.BundleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   job.BundleTable,
			Columns: []string{job.BundleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if juo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   job.TasksTable,
			Columns: []string{job.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := juo.mutation.RemovedTasksIDs(); len(nodes) > 0 && !juo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   job.TasksTable,
			Columns: []string{job.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := juo.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   job.TasksTable,
			Columns: []string{job.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Job{config: juo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, juo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{job.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kcarretto/realm/tavern/ent/job"
	"github.com/kcarretto/realm/tavern/ent/task"
	"github.com/kcarretto/realm/tavern/ent/tome"
	"github.com/kcarretto/realm/tavern/ent/user"
)

// JobCreate is the builder for creating a Job entity.
type JobCreate struct {
	config
	mutation *JobMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (jc *JobCreate) SetName(s string) *JobCreate {
	jc.mutation.SetName(s)
	return jc
}

// SetCreatedByID sets the "createdBy" edge to the User entity by ID.
func (jc *JobCreate) SetCreatedByID(id int) *JobCreate {
	jc.mutation.SetCreatedByID(id)
	return jc
}

// SetCreatedBy sets the "createdBy" edge to the User entity.
func (jc *JobCreate) SetCreatedBy(u *User) *JobCreate {
	return jc.SetCreatedByID(u.ID)
}

// SetTomeID sets the "tome" edge to the Tome entity by ID.
func (jc *JobCreate) SetTomeID(id int) *JobCreate {
	jc.mutation.SetTomeID(id)
	return jc
}

// SetTome sets the "tome" edge to the Tome entity.
func (jc *JobCreate) SetTome(t *Tome) *JobCreate {
	return jc.SetTomeID(t.ID)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (jc *JobCreate) AddTaskIDs(ids ...int) *JobCreate {
	jc.mutation.AddTaskIDs(ids...)
	return jc
}

// AddTasks adds the "tasks" edges to the Task entity.
func (jc *JobCreate) AddTasks(t ...*Task) *JobCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return jc.AddTaskIDs(ids...)
}

// Mutation returns the JobMutation object of the builder.
func (jc *JobCreate) Mutation() *JobMutation {
	return jc.mutation
}

// Save creates the Job in the database.
func (jc *JobCreate) Save(ctx context.Context) (*Job, error) {
	var (
		err  error
		node *Job
	)
	if len(jc.hooks) == 0 {
		if err = jc.check(); err != nil {
			return nil, err
		}
		node, err = jc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*JobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = jc.check(); err != nil {
				return nil, err
			}
			jc.mutation = mutation
			if node, err = jc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(jc.hooks) - 1; i >= 0; i-- {
			if jc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = jc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, jc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (jc *JobCreate) SaveX(ctx context.Context) *Job {
	v, err := jc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jc *JobCreate) Exec(ctx context.Context) error {
	_, err := jc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jc *JobCreate) ExecX(ctx context.Context) {
	if err := jc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jc *JobCreate) check() error {
	if _, ok := jc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Job.name"`)}
	}
	if v, ok := jc.mutation.Name(); ok {
		if err := job.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Job.name": %w`, err)}
		}
	}
	if _, ok := jc.mutation.CreatedByID(); !ok {
		return &ValidationError{Name: "createdBy", err: errors.New(`ent: missing required edge "Job.createdBy"`)}
	}
	if _, ok := jc.mutation.TomeID(); !ok {
		return &ValidationError{Name: "tome", err: errors.New(`ent: missing required edge "Job.tome"`)}
	}
	return nil
}

func (jc *JobCreate) sqlSave(ctx context.Context) (*Job, error) {
	_node, _spec := jc.createSpec()
	if err := sqlgraph.CreateNode(ctx, jc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (jc *JobCreate) createSpec() (*Job, *sqlgraph.CreateSpec) {
	var (
		_node = &Job{config: jc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: job.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: job.FieldID,
			},
		}
	)
	if value, ok := jc.mutation.Name(); ok {
		_spec.SetField(job.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := jc.mutation.CreatedByIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   job.CreatedByTable,
			Columns: []string{job.CreatedByColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.job_created_by = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jc.mutation.TomeIDs(); len(nodes) > 0 {
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
		_node.job_tome = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jc.mutation.TasksIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// JobCreateBulk is the builder for creating many Job entities in bulk.
type JobCreateBulk struct {
	config
	builders []*JobCreate
}

// Save creates the Job entities in the database.
func (jcb *JobCreateBulk) Save(ctx context.Context) ([]*Job, error) {
	specs := make([]*sqlgraph.CreateSpec, len(jcb.builders))
	nodes := make([]*Job, len(jcb.builders))
	mutators := make([]Mutator, len(jcb.builders))
	for i := range jcb.builders {
		func(i int, root context.Context) {
			builder := jcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*JobMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, jcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, jcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, jcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (jcb *JobCreateBulk) SaveX(ctx context.Context) []*Job {
	v, err := jcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jcb *JobCreateBulk) Exec(ctx context.Context) error {
	_, err := jcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jcb *JobCreateBulk) ExecX(ctx context.Context) {
	if err := jcb.Exec(ctx); err != nil {
		panic(err)
	}
}

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
	"github.com/kcarretto/realm/tavern/ent/predicate"
	"github.com/kcarretto/realm/tavern/ent/tag"
	"github.com/kcarretto/realm/tavern/ent/target"
)

// TargetUpdate is the builder for updating Target entities.
type TargetUpdate struct {
	config
	hooks    []Hook
	mutation *TargetMutation
}

// Where appends a list predicates to the TargetUpdate builder.
func (tu *TargetUpdate) Where(ps ...predicate.Target) *TargetUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetName sets the "name" field.
func (tu *TargetUpdate) SetName(s string) *TargetUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetLastSeenAt sets the "lastSeenAt" field.
func (tu *TargetUpdate) SetLastSeenAt(t time.Time) *TargetUpdate {
	tu.mutation.SetLastSeenAt(t)
	return tu
}

// SetNillableLastSeenAt sets the "lastSeenAt" field if the given value is not nil.
func (tu *TargetUpdate) SetNillableLastSeenAt(t *time.Time) *TargetUpdate {
	if t != nil {
		tu.SetLastSeenAt(*t)
	}
	return tu
}

// ClearLastSeenAt clears the value of the "lastSeenAt" field.
func (tu *TargetUpdate) ClearLastSeenAt() *TargetUpdate {
	tu.mutation.ClearLastSeenAt()
	return tu
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (tu *TargetUpdate) AddTagIDs(ids ...int) *TargetUpdate {
	tu.mutation.AddTagIDs(ids...)
	return tu
}

// AddTags adds the "tags" edges to the Tag entity.
func (tu *TargetUpdate) AddTags(t ...*Tag) *TargetUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTagIDs(ids...)
}

// Mutation returns the TargetMutation object of the builder.
func (tu *TargetUpdate) Mutation() *TargetMutation {
	return tu.mutation
}

// ClearTags clears all "tags" edges to the Tag entity.
func (tu *TargetUpdate) ClearTags() *TargetUpdate {
	tu.mutation.ClearTags()
	return tu
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (tu *TargetUpdate) RemoveTagIDs(ids ...int) *TargetUpdate {
	tu.mutation.RemoveTagIDs(ids...)
	return tu
}

// RemoveTags removes "tags" edges to Tag entities.
func (tu *TargetUpdate) RemoveTags(t ...*Tag) *TargetUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTagIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TargetUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TargetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TargetUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TargetUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TargetUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TargetUpdate) check() error {
	if v, ok := tu.mutation.Name(); ok {
		if err := target.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Target.name": %w`, err)}
		}
	}
	return nil
}

func (tu *TargetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   target.Table,
			Columns: target.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: target.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(target.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.LastSeenAt(); ok {
		_spec.SetField(target.FieldLastSeenAt, field.TypeTime, value)
	}
	if tu.mutation.LastSeenAtCleared() {
		_spec.ClearField(target.FieldLastSeenAt, field.TypeTime)
	}
	if tu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   target.TagsTable,
			Columns: target.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedTagsIDs(); len(nodes) > 0 && !tu.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   target.TagsTable,
			Columns: target.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   target.TagsTable,
			Columns: target.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{target.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TargetUpdateOne is the builder for updating a single Target entity.
type TargetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TargetMutation
}

// SetName sets the "name" field.
func (tuo *TargetUpdateOne) SetName(s string) *TargetUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetLastSeenAt sets the "lastSeenAt" field.
func (tuo *TargetUpdateOne) SetLastSeenAt(t time.Time) *TargetUpdateOne {
	tuo.mutation.SetLastSeenAt(t)
	return tuo
}

// SetNillableLastSeenAt sets the "lastSeenAt" field if the given value is not nil.
func (tuo *TargetUpdateOne) SetNillableLastSeenAt(t *time.Time) *TargetUpdateOne {
	if t != nil {
		tuo.SetLastSeenAt(*t)
	}
	return tuo
}

// ClearLastSeenAt clears the value of the "lastSeenAt" field.
func (tuo *TargetUpdateOne) ClearLastSeenAt() *TargetUpdateOne {
	tuo.mutation.ClearLastSeenAt()
	return tuo
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (tuo *TargetUpdateOne) AddTagIDs(ids ...int) *TargetUpdateOne {
	tuo.mutation.AddTagIDs(ids...)
	return tuo
}

// AddTags adds the "tags" edges to the Tag entity.
func (tuo *TargetUpdateOne) AddTags(t ...*Tag) *TargetUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTagIDs(ids...)
}

// Mutation returns the TargetMutation object of the builder.
func (tuo *TargetUpdateOne) Mutation() *TargetMutation {
	return tuo.mutation
}

// ClearTags clears all "tags" edges to the Tag entity.
func (tuo *TargetUpdateOne) ClearTags() *TargetUpdateOne {
	tuo.mutation.ClearTags()
	return tuo
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (tuo *TargetUpdateOne) RemoveTagIDs(ids ...int) *TargetUpdateOne {
	tuo.mutation.RemoveTagIDs(ids...)
	return tuo
}

// RemoveTags removes "tags" edges to Tag entities.
func (tuo *TargetUpdateOne) RemoveTags(t ...*Tag) *TargetUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTagIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TargetUpdateOne) Select(field string, fields ...string) *TargetUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Target entity.
func (tuo *TargetUpdateOne) Save(ctx context.Context) (*Target, error) {
	var (
		err  error
		node *Target
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TargetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Target)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TargetMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TargetUpdateOne) SaveX(ctx context.Context) *Target {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TargetUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TargetUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TargetUpdateOne) check() error {
	if v, ok := tuo.mutation.Name(); ok {
		if err := target.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Target.name": %w`, err)}
		}
	}
	return nil
}

func (tuo *TargetUpdateOne) sqlSave(ctx context.Context) (_node *Target, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   target.Table,
			Columns: target.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: target.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Target.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, target.FieldID)
		for _, f := range fields {
			if !target.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != target.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(target.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.LastSeenAt(); ok {
		_spec.SetField(target.FieldLastSeenAt, field.TypeTime, value)
	}
	if tuo.mutation.LastSeenAtCleared() {
		_spec.ClearField(target.FieldLastSeenAt, field.TypeTime)
	}
	if tuo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   target.TagsTable,
			Columns: target.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedTagsIDs(); len(nodes) > 0 && !tuo.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   target.TagsTable,
			Columns: target.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   target.TagsTable,
			Columns: target.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Target{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{target.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

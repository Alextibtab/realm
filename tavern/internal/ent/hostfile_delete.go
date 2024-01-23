// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"realm.pub/tavern/internal/ent/hostfile"
	"realm.pub/tavern/internal/ent/predicate"
)

// HostFileDelete is the builder for deleting a HostFile entity.
type HostFileDelete struct {
	config
	hooks    []Hook
	mutation *HostFileMutation
}

// Where appends a list predicates to the HostFileDelete builder.
func (hfd *HostFileDelete) Where(ps ...predicate.HostFile) *HostFileDelete {
	hfd.mutation.Where(ps...)
	return hfd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hfd *HostFileDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, hfd.sqlExec, hfd.mutation, hfd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (hfd *HostFileDelete) ExecX(ctx context.Context) int {
	n, err := hfd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hfd *HostFileDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(hostfile.Table, sqlgraph.NewFieldSpec(hostfile.FieldID, field.TypeInt))
	if ps := hfd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, hfd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	hfd.mutation.done = true
	return affected, err
}

// HostFileDeleteOne is the builder for deleting a single HostFile entity.
type HostFileDeleteOne struct {
	hfd *HostFileDelete
}

// Where appends a list predicates to the HostFileDelete builder.
func (hfdo *HostFileDeleteOne) Where(ps ...predicate.HostFile) *HostFileDeleteOne {
	hfdo.hfd.mutation.Where(ps...)
	return hfdo
}

// Exec executes the deletion query.
func (hfdo *HostFileDeleteOne) Exec(ctx context.Context) error {
	n, err := hfdo.hfd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{hostfile.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hfdo *HostFileDeleteOne) ExecX(ctx context.Context) {
	if err := hfdo.Exec(ctx); err != nil {
		panic(err)
	}
}

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
	"realm.pub/tavern/internal/ent/beacon"
	"realm.pub/tavern/internal/ent/hostfile"
	"realm.pub/tavern/internal/ent/hostprocess"
	"realm.pub/tavern/internal/ent/predicate"
	"realm.pub/tavern/internal/ent/quest"
	"realm.pub/tavern/internal/ent/task"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (tu *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (tu *TaskUpdate) SetLastModifiedAt(t time.Time) *TaskUpdate {
	tu.mutation.SetLastModifiedAt(t)
	return tu
}

// SetClaimedAt sets the "claimed_at" field.
func (tu *TaskUpdate) SetClaimedAt(t time.Time) *TaskUpdate {
	tu.mutation.SetClaimedAt(t)
	return tu
}

// SetNillableClaimedAt sets the "claimed_at" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableClaimedAt(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetClaimedAt(*t)
	}
	return tu
}

// ClearClaimedAt clears the value of the "claimed_at" field.
func (tu *TaskUpdate) ClearClaimedAt() *TaskUpdate {
	tu.mutation.ClearClaimedAt()
	return tu
}

// SetExecStartedAt sets the "exec_started_at" field.
func (tu *TaskUpdate) SetExecStartedAt(t time.Time) *TaskUpdate {
	tu.mutation.SetExecStartedAt(t)
	return tu
}

// SetNillableExecStartedAt sets the "exec_started_at" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableExecStartedAt(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetExecStartedAt(*t)
	}
	return tu
}

// ClearExecStartedAt clears the value of the "exec_started_at" field.
func (tu *TaskUpdate) ClearExecStartedAt() *TaskUpdate {
	tu.mutation.ClearExecStartedAt()
	return tu
}

// SetExecFinishedAt sets the "exec_finished_at" field.
func (tu *TaskUpdate) SetExecFinishedAt(t time.Time) *TaskUpdate {
	tu.mutation.SetExecFinishedAt(t)
	return tu
}

// SetNillableExecFinishedAt sets the "exec_finished_at" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableExecFinishedAt(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetExecFinishedAt(*t)
	}
	return tu
}

// ClearExecFinishedAt clears the value of the "exec_finished_at" field.
func (tu *TaskUpdate) ClearExecFinishedAt() *TaskUpdate {
	tu.mutation.ClearExecFinishedAt()
	return tu
}

// SetOutput sets the "output" field.
func (tu *TaskUpdate) SetOutput(s string) *TaskUpdate {
	tu.mutation.SetOutput(s)
	return tu
}

// SetNillableOutput sets the "output" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableOutput(s *string) *TaskUpdate {
	if s != nil {
		tu.SetOutput(*s)
	}
	return tu
}

// ClearOutput clears the value of the "output" field.
func (tu *TaskUpdate) ClearOutput() *TaskUpdate {
	tu.mutation.ClearOutput()
	return tu
}

// SetOutputSize sets the "output_size" field.
func (tu *TaskUpdate) SetOutputSize(i int) *TaskUpdate {
	tu.mutation.ResetOutputSize()
	tu.mutation.SetOutputSize(i)
	return tu
}

// SetNillableOutputSize sets the "output_size" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableOutputSize(i *int) *TaskUpdate {
	if i != nil {
		tu.SetOutputSize(*i)
	}
	return tu
}

// AddOutputSize adds i to the "output_size" field.
func (tu *TaskUpdate) AddOutputSize(i int) *TaskUpdate {
	tu.mutation.AddOutputSize(i)
	return tu
}

// SetError sets the "error" field.
func (tu *TaskUpdate) SetError(s string) *TaskUpdate {
	tu.mutation.SetError(s)
	return tu
}

// SetNillableError sets the "error" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableError(s *string) *TaskUpdate {
	if s != nil {
		tu.SetError(*s)
	}
	return tu
}

// ClearError clears the value of the "error" field.
func (tu *TaskUpdate) ClearError() *TaskUpdate {
	tu.mutation.ClearError()
	return tu
}

// SetQuestID sets the "quest" edge to the Quest entity by ID.
func (tu *TaskUpdate) SetQuestID(id int) *TaskUpdate {
	tu.mutation.SetQuestID(id)
	return tu
}

// SetQuest sets the "quest" edge to the Quest entity.
func (tu *TaskUpdate) SetQuest(q *Quest) *TaskUpdate {
	return tu.SetQuestID(q.ID)
}

// SetBeaconID sets the "beacon" edge to the Beacon entity by ID.
func (tu *TaskUpdate) SetBeaconID(id int) *TaskUpdate {
	tu.mutation.SetBeaconID(id)
	return tu
}

// SetBeacon sets the "beacon" edge to the Beacon entity.
func (tu *TaskUpdate) SetBeacon(b *Beacon) *TaskUpdate {
	return tu.SetBeaconID(b.ID)
}

// AddReportedFileIDs adds the "reported_files" edge to the HostFile entity by IDs.
func (tu *TaskUpdate) AddReportedFileIDs(ids ...int) *TaskUpdate {
	tu.mutation.AddReportedFileIDs(ids...)
	return tu
}

// AddReportedFiles adds the "reported_files" edges to the HostFile entity.
func (tu *TaskUpdate) AddReportedFiles(h ...*HostFile) *TaskUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return tu.AddReportedFileIDs(ids...)
}

// AddReportedProcessIDs adds the "reported_processes" edge to the HostProcess entity by IDs.
func (tu *TaskUpdate) AddReportedProcessIDs(ids ...int) *TaskUpdate {
	tu.mutation.AddReportedProcessIDs(ids...)
	return tu
}

// AddReportedProcesses adds the "reported_processes" edges to the HostProcess entity.
func (tu *TaskUpdate) AddReportedProcesses(h ...*HostProcess) *TaskUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return tu.AddReportedProcessIDs(ids...)
}

// Mutation returns the TaskMutation object of the builder.
func (tu *TaskUpdate) Mutation() *TaskMutation {
	return tu.mutation
}

// ClearQuest clears the "quest" edge to the Quest entity.
func (tu *TaskUpdate) ClearQuest() *TaskUpdate {
	tu.mutation.ClearQuest()
	return tu
}

// ClearBeacon clears the "beacon" edge to the Beacon entity.
func (tu *TaskUpdate) ClearBeacon() *TaskUpdate {
	tu.mutation.ClearBeacon()
	return tu
}

// ClearReportedFiles clears all "reported_files" edges to the HostFile entity.
func (tu *TaskUpdate) ClearReportedFiles() *TaskUpdate {
	tu.mutation.ClearReportedFiles()
	return tu
}

// RemoveReportedFileIDs removes the "reported_files" edge to HostFile entities by IDs.
func (tu *TaskUpdate) RemoveReportedFileIDs(ids ...int) *TaskUpdate {
	tu.mutation.RemoveReportedFileIDs(ids...)
	return tu
}

// RemoveReportedFiles removes "reported_files" edges to HostFile entities.
func (tu *TaskUpdate) RemoveReportedFiles(h ...*HostFile) *TaskUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return tu.RemoveReportedFileIDs(ids...)
}

// ClearReportedProcesses clears all "reported_processes" edges to the HostProcess entity.
func (tu *TaskUpdate) ClearReportedProcesses() *TaskUpdate {
	tu.mutation.ClearReportedProcesses()
	return tu
}

// RemoveReportedProcessIDs removes the "reported_processes" edge to HostProcess entities by IDs.
func (tu *TaskUpdate) RemoveReportedProcessIDs(ids ...int) *TaskUpdate {
	tu.mutation.RemoveReportedProcessIDs(ids...)
	return tu
}

// RemoveReportedProcesses removes "reported_processes" edges to HostProcess entities.
func (tu *TaskUpdate) RemoveReportedProcesses(h ...*HostProcess) *TaskUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return tu.RemoveReportedProcessIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TaskUpdate) Save(ctx context.Context) (int, error) {
	if err := tu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaskUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaskUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TaskUpdate) defaults() error {
	if _, ok := tu.mutation.LastModifiedAt(); !ok {
		if task.UpdateDefaultLastModifiedAt == nil {
			return fmt.Errorf("ent: uninitialized task.UpdateDefaultLastModifiedAt (forgotten import ent/runtime?)")
		}
		v := task.UpdateDefaultLastModifiedAt()
		tu.mutation.SetLastModifiedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tu *TaskUpdate) check() error {
	if v, ok := tu.mutation.OutputSize(); ok {
		if err := task.OutputSizeValidator(v); err != nil {
			return &ValidationError{Name: "output_size", err: fmt.Errorf(`ent: validator failed for field "Task.output_size": %w`, err)}
		}
	}
	if _, ok := tu.mutation.QuestID(); tu.mutation.QuestCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Task.quest"`)
	}
	if _, ok := tu.mutation.BeaconID(); tu.mutation.BeaconCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Task.beacon"`)
	}
	return nil
}

func (tu *TaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.LastModifiedAt(); ok {
		_spec.SetField(task.FieldLastModifiedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.ClaimedAt(); ok {
		_spec.SetField(task.FieldClaimedAt, field.TypeTime, value)
	}
	if tu.mutation.ClaimedAtCleared() {
		_spec.ClearField(task.FieldClaimedAt, field.TypeTime)
	}
	if value, ok := tu.mutation.ExecStartedAt(); ok {
		_spec.SetField(task.FieldExecStartedAt, field.TypeTime, value)
	}
	if tu.mutation.ExecStartedAtCleared() {
		_spec.ClearField(task.FieldExecStartedAt, field.TypeTime)
	}
	if value, ok := tu.mutation.ExecFinishedAt(); ok {
		_spec.SetField(task.FieldExecFinishedAt, field.TypeTime, value)
	}
	if tu.mutation.ExecFinishedAtCleared() {
		_spec.ClearField(task.FieldExecFinishedAt, field.TypeTime)
	}
	if value, ok := tu.mutation.Output(); ok {
		_spec.SetField(task.FieldOutput, field.TypeString, value)
	}
	if tu.mutation.OutputCleared() {
		_spec.ClearField(task.FieldOutput, field.TypeString)
	}
	if value, ok := tu.mutation.OutputSize(); ok {
		_spec.SetField(task.FieldOutputSize, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedOutputSize(); ok {
		_spec.AddField(task.FieldOutputSize, field.TypeInt, value)
	}
	if value, ok := tu.mutation.Error(); ok {
		_spec.SetField(task.FieldError, field.TypeString, value)
	}
	if tu.mutation.ErrorCleared() {
		_spec.ClearField(task.FieldError, field.TypeString)
	}
	if tu.mutation.QuestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.QuestTable,
			Columns: []string{task.QuestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(quest.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.QuestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.QuestTable,
			Columns: []string{task.QuestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(quest.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.BeaconCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   task.BeaconTable,
			Columns: []string{task.BeaconColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(beacon.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.BeaconIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   task.BeaconTable,
			Columns: []string{task.BeaconColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(beacon.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.ReportedFilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.ReportedFilesTable,
			Columns: []string{task.ReportedFilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostfile.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedReportedFilesIDs(); len(nodes) > 0 && !tu.mutation.ReportedFilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.ReportedFilesTable,
			Columns: []string{task.ReportedFilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostfile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ReportedFilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.ReportedFilesTable,
			Columns: []string{task.ReportedFilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostfile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.ReportedProcessesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.ReportedProcessesTable,
			Columns: []string{task.ReportedProcessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostprocess.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedReportedProcessesIDs(); len(nodes) > 0 && !tu.mutation.ReportedProcessesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.ReportedProcessesTable,
			Columns: []string{task.ReportedProcessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostprocess.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ReportedProcessesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.ReportedProcessesTable,
			Columns: []string{task.ReportedProcessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostprocess.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskMutation
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (tuo *TaskUpdateOne) SetLastModifiedAt(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetLastModifiedAt(t)
	return tuo
}

// SetClaimedAt sets the "claimed_at" field.
func (tuo *TaskUpdateOne) SetClaimedAt(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetClaimedAt(t)
	return tuo
}

// SetNillableClaimedAt sets the "claimed_at" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableClaimedAt(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetClaimedAt(*t)
	}
	return tuo
}

// ClearClaimedAt clears the value of the "claimed_at" field.
func (tuo *TaskUpdateOne) ClearClaimedAt() *TaskUpdateOne {
	tuo.mutation.ClearClaimedAt()
	return tuo
}

// SetExecStartedAt sets the "exec_started_at" field.
func (tuo *TaskUpdateOne) SetExecStartedAt(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetExecStartedAt(t)
	return tuo
}

// SetNillableExecStartedAt sets the "exec_started_at" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableExecStartedAt(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetExecStartedAt(*t)
	}
	return tuo
}

// ClearExecStartedAt clears the value of the "exec_started_at" field.
func (tuo *TaskUpdateOne) ClearExecStartedAt() *TaskUpdateOne {
	tuo.mutation.ClearExecStartedAt()
	return tuo
}

// SetExecFinishedAt sets the "exec_finished_at" field.
func (tuo *TaskUpdateOne) SetExecFinishedAt(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetExecFinishedAt(t)
	return tuo
}

// SetNillableExecFinishedAt sets the "exec_finished_at" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableExecFinishedAt(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetExecFinishedAt(*t)
	}
	return tuo
}

// ClearExecFinishedAt clears the value of the "exec_finished_at" field.
func (tuo *TaskUpdateOne) ClearExecFinishedAt() *TaskUpdateOne {
	tuo.mutation.ClearExecFinishedAt()
	return tuo
}

// SetOutput sets the "output" field.
func (tuo *TaskUpdateOne) SetOutput(s string) *TaskUpdateOne {
	tuo.mutation.SetOutput(s)
	return tuo
}

// SetNillableOutput sets the "output" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableOutput(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetOutput(*s)
	}
	return tuo
}

// ClearOutput clears the value of the "output" field.
func (tuo *TaskUpdateOne) ClearOutput() *TaskUpdateOne {
	tuo.mutation.ClearOutput()
	return tuo
}

// SetOutputSize sets the "output_size" field.
func (tuo *TaskUpdateOne) SetOutputSize(i int) *TaskUpdateOne {
	tuo.mutation.ResetOutputSize()
	tuo.mutation.SetOutputSize(i)
	return tuo
}

// SetNillableOutputSize sets the "output_size" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableOutputSize(i *int) *TaskUpdateOne {
	if i != nil {
		tuo.SetOutputSize(*i)
	}
	return tuo
}

// AddOutputSize adds i to the "output_size" field.
func (tuo *TaskUpdateOne) AddOutputSize(i int) *TaskUpdateOne {
	tuo.mutation.AddOutputSize(i)
	return tuo
}

// SetError sets the "error" field.
func (tuo *TaskUpdateOne) SetError(s string) *TaskUpdateOne {
	tuo.mutation.SetError(s)
	return tuo
}

// SetNillableError sets the "error" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableError(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetError(*s)
	}
	return tuo
}

// ClearError clears the value of the "error" field.
func (tuo *TaskUpdateOne) ClearError() *TaskUpdateOne {
	tuo.mutation.ClearError()
	return tuo
}

// SetQuestID sets the "quest" edge to the Quest entity by ID.
func (tuo *TaskUpdateOne) SetQuestID(id int) *TaskUpdateOne {
	tuo.mutation.SetQuestID(id)
	return tuo
}

// SetQuest sets the "quest" edge to the Quest entity.
func (tuo *TaskUpdateOne) SetQuest(q *Quest) *TaskUpdateOne {
	return tuo.SetQuestID(q.ID)
}

// SetBeaconID sets the "beacon" edge to the Beacon entity by ID.
func (tuo *TaskUpdateOne) SetBeaconID(id int) *TaskUpdateOne {
	tuo.mutation.SetBeaconID(id)
	return tuo
}

// SetBeacon sets the "beacon" edge to the Beacon entity.
func (tuo *TaskUpdateOne) SetBeacon(b *Beacon) *TaskUpdateOne {
	return tuo.SetBeaconID(b.ID)
}

// AddReportedFileIDs adds the "reported_files" edge to the HostFile entity by IDs.
func (tuo *TaskUpdateOne) AddReportedFileIDs(ids ...int) *TaskUpdateOne {
	tuo.mutation.AddReportedFileIDs(ids...)
	return tuo
}

// AddReportedFiles adds the "reported_files" edges to the HostFile entity.
func (tuo *TaskUpdateOne) AddReportedFiles(h ...*HostFile) *TaskUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return tuo.AddReportedFileIDs(ids...)
}

// AddReportedProcessIDs adds the "reported_processes" edge to the HostProcess entity by IDs.
func (tuo *TaskUpdateOne) AddReportedProcessIDs(ids ...int) *TaskUpdateOne {
	tuo.mutation.AddReportedProcessIDs(ids...)
	return tuo
}

// AddReportedProcesses adds the "reported_processes" edges to the HostProcess entity.
func (tuo *TaskUpdateOne) AddReportedProcesses(h ...*HostProcess) *TaskUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return tuo.AddReportedProcessIDs(ids...)
}

// Mutation returns the TaskMutation object of the builder.
func (tuo *TaskUpdateOne) Mutation() *TaskMutation {
	return tuo.mutation
}

// ClearQuest clears the "quest" edge to the Quest entity.
func (tuo *TaskUpdateOne) ClearQuest() *TaskUpdateOne {
	tuo.mutation.ClearQuest()
	return tuo
}

// ClearBeacon clears the "beacon" edge to the Beacon entity.
func (tuo *TaskUpdateOne) ClearBeacon() *TaskUpdateOne {
	tuo.mutation.ClearBeacon()
	return tuo
}

// ClearReportedFiles clears all "reported_files" edges to the HostFile entity.
func (tuo *TaskUpdateOne) ClearReportedFiles() *TaskUpdateOne {
	tuo.mutation.ClearReportedFiles()
	return tuo
}

// RemoveReportedFileIDs removes the "reported_files" edge to HostFile entities by IDs.
func (tuo *TaskUpdateOne) RemoveReportedFileIDs(ids ...int) *TaskUpdateOne {
	tuo.mutation.RemoveReportedFileIDs(ids...)
	return tuo
}

// RemoveReportedFiles removes "reported_files" edges to HostFile entities.
func (tuo *TaskUpdateOne) RemoveReportedFiles(h ...*HostFile) *TaskUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return tuo.RemoveReportedFileIDs(ids...)
}

// ClearReportedProcesses clears all "reported_processes" edges to the HostProcess entity.
func (tuo *TaskUpdateOne) ClearReportedProcesses() *TaskUpdateOne {
	tuo.mutation.ClearReportedProcesses()
	return tuo
}

// RemoveReportedProcessIDs removes the "reported_processes" edge to HostProcess entities by IDs.
func (tuo *TaskUpdateOne) RemoveReportedProcessIDs(ids ...int) *TaskUpdateOne {
	tuo.mutation.RemoveReportedProcessIDs(ids...)
	return tuo
}

// RemoveReportedProcesses removes "reported_processes" edges to HostProcess entities.
func (tuo *TaskUpdateOne) RemoveReportedProcesses(h ...*HostProcess) *TaskUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return tuo.RemoveReportedProcessIDs(ids...)
}

// Where appends a list predicates to the TaskUpdate builder.
func (tuo *TaskUpdateOne) Where(ps ...predicate.Task) *TaskUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TaskUpdateOne) Select(field string, fields ...string) *TaskUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Task entity.
func (tuo *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	if err := tuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TaskUpdateOne) defaults() error {
	if _, ok := tuo.mutation.LastModifiedAt(); !ok {
		if task.UpdateDefaultLastModifiedAt == nil {
			return fmt.Errorf("ent: uninitialized task.UpdateDefaultLastModifiedAt (forgotten import ent/runtime?)")
		}
		v := task.UpdateDefaultLastModifiedAt()
		tuo.mutation.SetLastModifiedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TaskUpdateOne) check() error {
	if v, ok := tuo.mutation.OutputSize(); ok {
		if err := task.OutputSizeValidator(v); err != nil {
			return &ValidationError{Name: "output_size", err: fmt.Errorf(`ent: validator failed for field "Task.output_size": %w`, err)}
		}
	}
	if _, ok := tuo.mutation.QuestID(); tuo.mutation.QuestCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Task.quest"`)
	}
	if _, ok := tuo.mutation.BeaconID(); tuo.mutation.BeaconCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Task.beacon"`)
	}
	return nil
}

func (tuo *TaskUpdateOne) sqlSave(ctx context.Context) (_node *Task, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Task.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, task.FieldID)
		for _, f := range fields {
			if !task.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != task.FieldID {
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
	if value, ok := tuo.mutation.LastModifiedAt(); ok {
		_spec.SetField(task.FieldLastModifiedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.ClaimedAt(); ok {
		_spec.SetField(task.FieldClaimedAt, field.TypeTime, value)
	}
	if tuo.mutation.ClaimedAtCleared() {
		_spec.ClearField(task.FieldClaimedAt, field.TypeTime)
	}
	if value, ok := tuo.mutation.ExecStartedAt(); ok {
		_spec.SetField(task.FieldExecStartedAt, field.TypeTime, value)
	}
	if tuo.mutation.ExecStartedAtCleared() {
		_spec.ClearField(task.FieldExecStartedAt, field.TypeTime)
	}
	if value, ok := tuo.mutation.ExecFinishedAt(); ok {
		_spec.SetField(task.FieldExecFinishedAt, field.TypeTime, value)
	}
	if tuo.mutation.ExecFinishedAtCleared() {
		_spec.ClearField(task.FieldExecFinishedAt, field.TypeTime)
	}
	if value, ok := tuo.mutation.Output(); ok {
		_spec.SetField(task.FieldOutput, field.TypeString, value)
	}
	if tuo.mutation.OutputCleared() {
		_spec.ClearField(task.FieldOutput, field.TypeString)
	}
	if value, ok := tuo.mutation.OutputSize(); ok {
		_spec.SetField(task.FieldOutputSize, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedOutputSize(); ok {
		_spec.AddField(task.FieldOutputSize, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.Error(); ok {
		_spec.SetField(task.FieldError, field.TypeString, value)
	}
	if tuo.mutation.ErrorCleared() {
		_spec.ClearField(task.FieldError, field.TypeString)
	}
	if tuo.mutation.QuestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.QuestTable,
			Columns: []string{task.QuestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(quest.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.QuestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.QuestTable,
			Columns: []string{task.QuestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(quest.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.BeaconCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   task.BeaconTable,
			Columns: []string{task.BeaconColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(beacon.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.BeaconIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   task.BeaconTable,
			Columns: []string{task.BeaconColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(beacon.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.ReportedFilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.ReportedFilesTable,
			Columns: []string{task.ReportedFilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostfile.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedReportedFilesIDs(); len(nodes) > 0 && !tuo.mutation.ReportedFilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.ReportedFilesTable,
			Columns: []string{task.ReportedFilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostfile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ReportedFilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.ReportedFilesTable,
			Columns: []string{task.ReportedFilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostfile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.ReportedProcessesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.ReportedProcessesTable,
			Columns: []string{task.ReportedProcessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostprocess.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedReportedProcessesIDs(); len(nodes) > 0 && !tuo.mutation.ReportedProcessesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.ReportedProcessesTable,
			Columns: []string{task.ReportedProcessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostprocess.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ReportedProcessesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.ReportedProcessesTable,
			Columns: []string{task.ReportedProcessesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(hostprocess.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Task{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}

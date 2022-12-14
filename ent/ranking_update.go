// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-ranking-api/ent/predicate"
	"go-ranking-api/ent/ranking"
	"go-ranking-api/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// RankingUpdate is the builder for updating Ranking entities.
type RankingUpdate struct {
	config
	hooks    []Hook
	mutation *RankingMutation
}

// Where appends a list predicates to the RankingUpdate builder.
func (ru *RankingUpdate) Where(ps ...predicate.Ranking) *RankingUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetScore sets the "score" field.
func (ru *RankingUpdate) SetScore(i int64) *RankingUpdate {
	ru.mutation.ResetScore()
	ru.mutation.SetScore(i)
	return ru
}

// AddScore adds i to the "score" field.
func (ru *RankingUpdate) AddScore(i int64) *RankingUpdate {
	ru.mutation.AddScore(i)
	return ru
}

// SetSongUUID sets the "song_uuid" field.
func (ru *RankingUpdate) SetSongUUID(u uuid.UUID) *RankingUpdate {
	ru.mutation.SetSongUUID(u)
	return ru
}

// SetUpdatedAt sets the "updated_at" field.
func (ru *RankingUpdate) SetUpdatedAt(t time.Time) *RankingUpdate {
	ru.mutation.SetUpdatedAt(t)
	return ru
}

// SetDeletedAt sets the "deleted_at" field.
func (ru *RankingUpdate) SetDeletedAt(t time.Time) *RankingUpdate {
	ru.mutation.SetDeletedAt(t)
	return ru
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ru *RankingUpdate) SetNillableDeletedAt(t *time.Time) *RankingUpdate {
	if t != nil {
		ru.SetDeletedAt(*t)
	}
	return ru
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ru *RankingUpdate) ClearDeletedAt() *RankingUpdate {
	ru.mutation.ClearDeletedAt()
	return ru
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ru *RankingUpdate) SetUserID(id int) *RankingUpdate {
	ru.mutation.SetUserID(id)
	return ru
}

// SetUser sets the "user" edge to the User entity.
func (ru *RankingUpdate) SetUser(u *User) *RankingUpdate {
	return ru.SetUserID(u.ID)
}

// Mutation returns the RankingMutation object of the builder.
func (ru *RankingUpdate) Mutation() *RankingMutation {
	return ru.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ru *RankingUpdate) ClearUser() *RankingUpdate {
	ru.mutation.ClearUser()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RankingUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ru.defaults()
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RankingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RankingUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RankingUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RankingUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *RankingUpdate) defaults() {
	if _, ok := ru.mutation.UpdatedAt(); !ok {
		v := ranking.UpdateDefaultUpdatedAt()
		ru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RankingUpdate) check() error {
	if v, ok := ru.mutation.Score(); ok {
		if err := ranking.ScoreValidator(v); err != nil {
			return &ValidationError{Name: "score", err: fmt.Errorf(`ent: validator failed for field "Ranking.score": %w`, err)}
		}
	}
	if _, ok := ru.mutation.UserID(); ru.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Ranking.user"`)
	}
	return nil
}

func (ru *RankingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ranking.Table,
			Columns: ranking.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ranking.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Score(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: ranking.FieldScore,
		})
	}
	if value, ok := ru.mutation.AddedScore(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: ranking.FieldScore,
		})
	}
	if value, ok := ru.mutation.SongUUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: ranking.FieldSongUUID,
		})
	}
	if value, ok := ru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: ranking.FieldUpdatedAt,
		})
	}
	if value, ok := ru.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: ranking.FieldDeletedAt,
		})
	}
	if ru.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: ranking.FieldDeletedAt,
		})
	}
	if ru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ranking.UserTable,
			Columns: []string{ranking.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ranking.UserTable,
			Columns: []string{ranking.UserColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ranking.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// RankingUpdateOne is the builder for updating a single Ranking entity.
type RankingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RankingMutation
}

// SetScore sets the "score" field.
func (ruo *RankingUpdateOne) SetScore(i int64) *RankingUpdateOne {
	ruo.mutation.ResetScore()
	ruo.mutation.SetScore(i)
	return ruo
}

// AddScore adds i to the "score" field.
func (ruo *RankingUpdateOne) AddScore(i int64) *RankingUpdateOne {
	ruo.mutation.AddScore(i)
	return ruo
}

// SetSongUUID sets the "song_uuid" field.
func (ruo *RankingUpdateOne) SetSongUUID(u uuid.UUID) *RankingUpdateOne {
	ruo.mutation.SetSongUUID(u)
	return ruo
}

// SetUpdatedAt sets the "updated_at" field.
func (ruo *RankingUpdateOne) SetUpdatedAt(t time.Time) *RankingUpdateOne {
	ruo.mutation.SetUpdatedAt(t)
	return ruo
}

// SetDeletedAt sets the "deleted_at" field.
func (ruo *RankingUpdateOne) SetDeletedAt(t time.Time) *RankingUpdateOne {
	ruo.mutation.SetDeletedAt(t)
	return ruo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ruo *RankingUpdateOne) SetNillableDeletedAt(t *time.Time) *RankingUpdateOne {
	if t != nil {
		ruo.SetDeletedAt(*t)
	}
	return ruo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ruo *RankingUpdateOne) ClearDeletedAt() *RankingUpdateOne {
	ruo.mutation.ClearDeletedAt()
	return ruo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ruo *RankingUpdateOne) SetUserID(id int) *RankingUpdateOne {
	ruo.mutation.SetUserID(id)
	return ruo
}

// SetUser sets the "user" edge to the User entity.
func (ruo *RankingUpdateOne) SetUser(u *User) *RankingUpdateOne {
	return ruo.SetUserID(u.ID)
}

// Mutation returns the RankingMutation object of the builder.
func (ruo *RankingUpdateOne) Mutation() *RankingMutation {
	return ruo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ruo *RankingUpdateOne) ClearUser() *RankingUpdateOne {
	ruo.mutation.ClearUser()
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RankingUpdateOne) Select(field string, fields ...string) *RankingUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Ranking entity.
func (ruo *RankingUpdateOne) Save(ctx context.Context) (*Ranking, error) {
	var (
		err  error
		node *Ranking
	)
	ruo.defaults()
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RankingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ruo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Ranking)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RankingMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RankingUpdateOne) SaveX(ctx context.Context) *Ranking {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RankingUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RankingUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *RankingUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdatedAt(); !ok {
		v := ranking.UpdateDefaultUpdatedAt()
		ruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RankingUpdateOne) check() error {
	if v, ok := ruo.mutation.Score(); ok {
		if err := ranking.ScoreValidator(v); err != nil {
			return &ValidationError{Name: "score", err: fmt.Errorf(`ent: validator failed for field "Ranking.score": %w`, err)}
		}
	}
	if _, ok := ruo.mutation.UserID(); ruo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Ranking.user"`)
	}
	return nil
}

func (ruo *RankingUpdateOne) sqlSave(ctx context.Context) (_node *Ranking, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   ranking.Table,
			Columns: ranking.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: ranking.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Ranking.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ranking.FieldID)
		for _, f := range fields {
			if !ranking.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != ranking.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Score(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: ranking.FieldScore,
		})
	}
	if value, ok := ruo.mutation.AddedScore(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: ranking.FieldScore,
		})
	}
	if value, ok := ruo.mutation.SongUUID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: ranking.FieldSongUUID,
		})
	}
	if value, ok := ruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: ranking.FieldUpdatedAt,
		})
	}
	if value, ok := ruo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: ranking.FieldDeletedAt,
		})
	}
	if ruo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: ranking.FieldDeletedAt,
		})
	}
	if ruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ranking.UserTable,
			Columns: []string{ranking.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ranking.UserTable,
			Columns: []string{ranking.UserColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Ranking{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ranking.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

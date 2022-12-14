// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-ranking-api/ent/asset"
	"go-ranking-api/ent/song"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AssetCreate is the builder for creating a Asset entity.
type AssetCreate struct {
	config
	mutation *AssetMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetPlatform sets the "platform" field.
func (ac *AssetCreate) SetPlatform(a asset.Platform) *AssetCreate {
	ac.mutation.SetPlatform(a)
	return ac
}

// SetNillablePlatform sets the "platform" field if the given value is not nil.
func (ac *AssetCreate) SetNillablePlatform(a *asset.Platform) *AssetCreate {
	if a != nil {
		ac.SetPlatform(*a)
	}
	return ac
}

// SetPath sets the "path" field.
func (ac *AssetCreate) SetPath(s string) *AssetCreate {
	ac.mutation.SetPath(s)
	return ac
}

// SetHash sets the "hash" field.
func (ac *AssetCreate) SetHash(s string) *AssetCreate {
	ac.mutation.SetHash(s)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AssetCreate) SetCreatedAt(t time.Time) *AssetCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AssetCreate) SetNillableCreatedAt(t *time.Time) *AssetCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AssetCreate) SetUpdatedAt(t time.Time) *AssetCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AssetCreate) SetNillableUpdatedAt(t *time.Time) *AssetCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetDeletedAt sets the "deleted_at" field.
func (ac *AssetCreate) SetDeletedAt(t time.Time) *AssetCreate {
	ac.mutation.SetDeletedAt(t)
	return ac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ac *AssetCreate) SetNillableDeletedAt(t *time.Time) *AssetCreate {
	if t != nil {
		ac.SetDeletedAt(*t)
	}
	return ac
}

// SetSongID sets the "song" edge to the Song entity by ID.
func (ac *AssetCreate) SetSongID(id int) *AssetCreate {
	ac.mutation.SetSongID(id)
	return ac
}

// SetSong sets the "song" edge to the Song entity.
func (ac *AssetCreate) SetSong(s *Song) *AssetCreate {
	return ac.SetSongID(s.ID)
}

// Mutation returns the AssetMutation object of the builder.
func (ac *AssetCreate) Mutation() *AssetMutation {
	return ac.mutation
}

// Save creates the Asset in the database.
func (ac *AssetCreate) Save(ctx context.Context) (*Asset, error) {
	var (
		err  error
		node *Asset
	)
	ac.defaults()
	if len(ac.hooks) == 0 {
		if err = ac.check(); err != nil {
			return nil, err
		}
		node, err = ac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AssetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ac.check(); err != nil {
				return nil, err
			}
			ac.mutation = mutation
			if node, err = ac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ac.hooks) - 1; i >= 0; i-- {
			if ac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ac.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ac.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Asset)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AssetMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AssetCreate) SaveX(ctx context.Context) *Asset {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AssetCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AssetCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AssetCreate) defaults() {
	if _, ok := ac.mutation.Platform(); !ok {
		v := asset.DefaultPlatform
		ac.mutation.SetPlatform(v)
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := asset.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := asset.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AssetCreate) check() error {
	if _, ok := ac.mutation.Platform(); !ok {
		return &ValidationError{Name: "platform", err: errors.New(`ent: missing required field "Asset.platform"`)}
	}
	if v, ok := ac.mutation.Platform(); ok {
		if err := asset.PlatformValidator(v); err != nil {
			return &ValidationError{Name: "platform", err: fmt.Errorf(`ent: validator failed for field "Asset.platform": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "Asset.path"`)}
	}
	if v, ok := ac.mutation.Path(); ok {
		if err := asset.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "Asset.path": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Hash(); !ok {
		return &ValidationError{Name: "hash", err: errors.New(`ent: missing required field "Asset.hash"`)}
	}
	if v, ok := ac.mutation.Hash(); ok {
		if err := asset.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf(`ent: validator failed for field "Asset.hash": %w`, err)}
		}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Asset.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Asset.updated_at"`)}
	}
	if _, ok := ac.mutation.SongID(); !ok {
		return &ValidationError{Name: "song", err: errors.New(`ent: missing required edge "Asset.song"`)}
	}
	return nil
}

func (ac *AssetCreate) sqlSave(ctx context.Context) (*Asset, error) {
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ac *AssetCreate) createSpec() (*Asset, *sqlgraph.CreateSpec) {
	var (
		_node = &Asset{config: ac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: asset.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: asset.FieldID,
			},
		}
	)
	_spec.OnConflict = ac.conflict
	if value, ok := ac.mutation.Platform(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: asset.FieldPlatform,
		})
		_node.Platform = value
	}
	if value, ok := ac.mutation.Path(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldPath,
		})
		_node.Path = value
	}
	if value, ok := ac.mutation.Hash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: asset.FieldHash,
		})
		_node.Hash = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asset.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asset.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: asset.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if nodes := ac.mutation.SongIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   asset.SongTable,
			Columns: []string{asset.SongColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: song.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.song_asset = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Asset.Create().
//		SetPlatform(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AssetUpsert) {
//			SetPlatform(v+v).
//		}).
//		Exec(ctx)
func (ac *AssetCreate) OnConflict(opts ...sql.ConflictOption) *AssetUpsertOne {
	ac.conflict = opts
	return &AssetUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Asset.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AssetCreate) OnConflictColumns(columns ...string) *AssetUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AssetUpsertOne{
		create: ac,
	}
}

type (
	// AssetUpsertOne is the builder for "upsert"-ing
	//  one Asset node.
	AssetUpsertOne struct {
		create *AssetCreate
	}

	// AssetUpsert is the "OnConflict" setter.
	AssetUpsert struct {
		*sql.UpdateSet
	}
)

// SetPlatform sets the "platform" field.
func (u *AssetUpsert) SetPlatform(v asset.Platform) *AssetUpsert {
	u.Set(asset.FieldPlatform, v)
	return u
}

// UpdatePlatform sets the "platform" field to the value that was provided on create.
func (u *AssetUpsert) UpdatePlatform() *AssetUpsert {
	u.SetExcluded(asset.FieldPlatform)
	return u
}

// SetPath sets the "path" field.
func (u *AssetUpsert) SetPath(v string) *AssetUpsert {
	u.Set(asset.FieldPath, v)
	return u
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *AssetUpsert) UpdatePath() *AssetUpsert {
	u.SetExcluded(asset.FieldPath)
	return u
}

// SetHash sets the "hash" field.
func (u *AssetUpsert) SetHash(v string) *AssetUpsert {
	u.Set(asset.FieldHash, v)
	return u
}

// UpdateHash sets the "hash" field to the value that was provided on create.
func (u *AssetUpsert) UpdateHash() *AssetUpsert {
	u.SetExcluded(asset.FieldHash)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AssetUpsert) SetCreatedAt(v time.Time) *AssetUpsert {
	u.Set(asset.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AssetUpsert) UpdateCreatedAt() *AssetUpsert {
	u.SetExcluded(asset.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AssetUpsert) SetUpdatedAt(v time.Time) *AssetUpsert {
	u.Set(asset.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AssetUpsert) UpdateUpdatedAt() *AssetUpsert {
	u.SetExcluded(asset.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AssetUpsert) SetDeletedAt(v time.Time) *AssetUpsert {
	u.Set(asset.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AssetUpsert) UpdateDeletedAt() *AssetUpsert {
	u.SetExcluded(asset.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *AssetUpsert) ClearDeletedAt() *AssetUpsert {
	u.SetNull(asset.FieldDeletedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Asset.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AssetUpsertOne) UpdateNewValues() *AssetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Asset.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AssetUpsertOne) Ignore() *AssetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AssetUpsertOne) DoNothing() *AssetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AssetCreate.OnConflict
// documentation for more info.
func (u *AssetUpsertOne) Update(set func(*AssetUpsert)) *AssetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AssetUpsert{UpdateSet: update})
	}))
	return u
}

// SetPlatform sets the "platform" field.
func (u *AssetUpsertOne) SetPlatform(v asset.Platform) *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.SetPlatform(v)
	})
}

// UpdatePlatform sets the "platform" field to the value that was provided on create.
func (u *AssetUpsertOne) UpdatePlatform() *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.UpdatePlatform()
	})
}

// SetPath sets the "path" field.
func (u *AssetUpsertOne) SetPath(v string) *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.SetPath(v)
	})
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *AssetUpsertOne) UpdatePath() *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.UpdatePath()
	})
}

// SetHash sets the "hash" field.
func (u *AssetUpsertOne) SetHash(v string) *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.SetHash(v)
	})
}

// UpdateHash sets the "hash" field to the value that was provided on create.
func (u *AssetUpsertOne) UpdateHash() *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.UpdateHash()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *AssetUpsertOne) SetCreatedAt(v time.Time) *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AssetUpsertOne) UpdateCreatedAt() *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AssetUpsertOne) SetUpdatedAt(v time.Time) *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AssetUpsertOne) UpdateUpdatedAt() *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AssetUpsertOne) SetDeletedAt(v time.Time) *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AssetUpsertOne) UpdateDeletedAt() *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *AssetUpsertOne) ClearDeletedAt() *AssetUpsertOne {
	return u.Update(func(s *AssetUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *AssetUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AssetCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AssetUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AssetUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AssetUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AssetCreateBulk is the builder for creating many Asset entities in bulk.
type AssetCreateBulk struct {
	config
	builders []*AssetCreate
	conflict []sql.ConflictOption
}

// Save creates the Asset entities in the database.
func (acb *AssetCreateBulk) Save(ctx context.Context) ([]*Asset, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Asset, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AssetMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AssetCreateBulk) SaveX(ctx context.Context) []*Asset {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AssetCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AssetCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Asset.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AssetUpsert) {
//			SetPlatform(v+v).
//		}).
//		Exec(ctx)
func (acb *AssetCreateBulk) OnConflict(opts ...sql.ConflictOption) *AssetUpsertBulk {
	acb.conflict = opts
	return &AssetUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Asset.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AssetCreateBulk) OnConflictColumns(columns ...string) *AssetUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AssetUpsertBulk{
		create: acb,
	}
}

// AssetUpsertBulk is the builder for "upsert"-ing
// a bulk of Asset nodes.
type AssetUpsertBulk struct {
	create *AssetCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Asset.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AssetUpsertBulk) UpdateNewValues() *AssetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Asset.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AssetUpsertBulk) Ignore() *AssetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AssetUpsertBulk) DoNothing() *AssetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AssetCreateBulk.OnConflict
// documentation for more info.
func (u *AssetUpsertBulk) Update(set func(*AssetUpsert)) *AssetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AssetUpsert{UpdateSet: update})
	}))
	return u
}

// SetPlatform sets the "platform" field.
func (u *AssetUpsertBulk) SetPlatform(v asset.Platform) *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.SetPlatform(v)
	})
}

// UpdatePlatform sets the "platform" field to the value that was provided on create.
func (u *AssetUpsertBulk) UpdatePlatform() *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.UpdatePlatform()
	})
}

// SetPath sets the "path" field.
func (u *AssetUpsertBulk) SetPath(v string) *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.SetPath(v)
	})
}

// UpdatePath sets the "path" field to the value that was provided on create.
func (u *AssetUpsertBulk) UpdatePath() *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.UpdatePath()
	})
}

// SetHash sets the "hash" field.
func (u *AssetUpsertBulk) SetHash(v string) *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.SetHash(v)
	})
}

// UpdateHash sets the "hash" field to the value that was provided on create.
func (u *AssetUpsertBulk) UpdateHash() *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.UpdateHash()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *AssetUpsertBulk) SetCreatedAt(v time.Time) *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AssetUpsertBulk) UpdateCreatedAt() *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AssetUpsertBulk) SetUpdatedAt(v time.Time) *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AssetUpsertBulk) UpdateUpdatedAt() *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AssetUpsertBulk) SetDeletedAt(v time.Time) *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AssetUpsertBulk) UpdateDeletedAt() *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *AssetUpsertBulk) ClearDeletedAt() *AssetUpsertBulk {
	return u.Update(func(s *AssetUpsert) {
		s.ClearDeletedAt()
	})
}

// Exec executes the query.
func (u *AssetUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AssetCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AssetCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AssetUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

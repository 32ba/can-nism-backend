// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"go-ranking-api/ent/migrate"

	"go-ranking-api/ent/asset"
	"go-ranking-api/ent/ranking"
	"go-ranking-api/ent/song"
	"go-ranking-api/ent/token"
	"go-ranking-api/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Asset is the client for interacting with the Asset builders.
	Asset *AssetClient
	// Ranking is the client for interacting with the Ranking builders.
	Ranking *RankingClient
	// Song is the client for interacting with the Song builders.
	Song *SongClient
	// Token is the client for interacting with the Token builders.
	Token *TokenClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Asset = NewAssetClient(c.config)
	c.Ranking = NewRankingClient(c.config)
	c.Song = NewSongClient(c.config)
	c.Token = NewTokenClient(c.config)
	c.User = NewUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Asset:   NewAssetClient(cfg),
		Ranking: NewRankingClient(cfg),
		Song:    NewSongClient(cfg),
		Token:   NewTokenClient(cfg),
		User:    NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Asset:   NewAssetClient(cfg),
		Ranking: NewRankingClient(cfg),
		Song:    NewSongClient(cfg),
		Token:   NewTokenClient(cfg),
		User:    NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Asset.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Asset.Use(hooks...)
	c.Ranking.Use(hooks...)
	c.Song.Use(hooks...)
	c.Token.Use(hooks...)
	c.User.Use(hooks...)
}

// AssetClient is a client for the Asset schema.
type AssetClient struct {
	config
}

// NewAssetClient returns a client for the Asset from the given config.
func NewAssetClient(c config) *AssetClient {
	return &AssetClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `asset.Hooks(f(g(h())))`.
func (c *AssetClient) Use(hooks ...Hook) {
	c.hooks.Asset = append(c.hooks.Asset, hooks...)
}

// Create returns a builder for creating a Asset entity.
func (c *AssetClient) Create() *AssetCreate {
	mutation := newAssetMutation(c.config, OpCreate)
	return &AssetCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Asset entities.
func (c *AssetClient) CreateBulk(builders ...*AssetCreate) *AssetCreateBulk {
	return &AssetCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Asset.
func (c *AssetClient) Update() *AssetUpdate {
	mutation := newAssetMutation(c.config, OpUpdate)
	return &AssetUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AssetClient) UpdateOne(a *Asset) *AssetUpdateOne {
	mutation := newAssetMutation(c.config, OpUpdateOne, withAsset(a))
	return &AssetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AssetClient) UpdateOneID(id int) *AssetUpdateOne {
	mutation := newAssetMutation(c.config, OpUpdateOne, withAssetID(id))
	return &AssetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Asset.
func (c *AssetClient) Delete() *AssetDelete {
	mutation := newAssetMutation(c.config, OpDelete)
	return &AssetDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AssetClient) DeleteOne(a *Asset) *AssetDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *AssetClient) DeleteOneID(id int) *AssetDeleteOne {
	builder := c.Delete().Where(asset.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AssetDeleteOne{builder}
}

// Query returns a query builder for Asset.
func (c *AssetClient) Query() *AssetQuery {
	return &AssetQuery{
		config: c.config,
	}
}

// Get returns a Asset entity by its id.
func (c *AssetClient) Get(ctx context.Context, id int) (*Asset, error) {
	return c.Query().Where(asset.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AssetClient) GetX(ctx context.Context, id int) *Asset {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySong queries the song edge of a Asset.
func (c *AssetClient) QuerySong(a *Asset) *SongQuery {
	query := &SongQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(asset.Table, asset.FieldID, id),
			sqlgraph.To(song.Table, song.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, asset.SongTable, asset.SongColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AssetClient) Hooks() []Hook {
	return c.hooks.Asset
}

// RankingClient is a client for the Ranking schema.
type RankingClient struct {
	config
}

// NewRankingClient returns a client for the Ranking from the given config.
func NewRankingClient(c config) *RankingClient {
	return &RankingClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `ranking.Hooks(f(g(h())))`.
func (c *RankingClient) Use(hooks ...Hook) {
	c.hooks.Ranking = append(c.hooks.Ranking, hooks...)
}

// Create returns a builder for creating a Ranking entity.
func (c *RankingClient) Create() *RankingCreate {
	mutation := newRankingMutation(c.config, OpCreate)
	return &RankingCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Ranking entities.
func (c *RankingClient) CreateBulk(builders ...*RankingCreate) *RankingCreateBulk {
	return &RankingCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Ranking.
func (c *RankingClient) Update() *RankingUpdate {
	mutation := newRankingMutation(c.config, OpUpdate)
	return &RankingUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RankingClient) UpdateOne(r *Ranking) *RankingUpdateOne {
	mutation := newRankingMutation(c.config, OpUpdateOne, withRanking(r))
	return &RankingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RankingClient) UpdateOneID(id int) *RankingUpdateOne {
	mutation := newRankingMutation(c.config, OpUpdateOne, withRankingID(id))
	return &RankingUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Ranking.
func (c *RankingClient) Delete() *RankingDelete {
	mutation := newRankingMutation(c.config, OpDelete)
	return &RankingDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RankingClient) DeleteOne(r *Ranking) *RankingDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *RankingClient) DeleteOneID(id int) *RankingDeleteOne {
	builder := c.Delete().Where(ranking.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RankingDeleteOne{builder}
}

// Query returns a query builder for Ranking.
func (c *RankingClient) Query() *RankingQuery {
	return &RankingQuery{
		config: c.config,
	}
}

// Get returns a Ranking entity by its id.
func (c *RankingClient) Get(ctx context.Context, id int) (*Ranking, error) {
	return c.Query().Where(ranking.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RankingClient) GetX(ctx context.Context, id int) *Ranking {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Ranking.
func (c *RankingClient) QueryUser(r *Ranking) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ranking.Table, ranking.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ranking.UserTable, ranking.UserColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RankingClient) Hooks() []Hook {
	return c.hooks.Ranking
}

// SongClient is a client for the Song schema.
type SongClient struct {
	config
}

// NewSongClient returns a client for the Song from the given config.
func NewSongClient(c config) *SongClient {
	return &SongClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `song.Hooks(f(g(h())))`.
func (c *SongClient) Use(hooks ...Hook) {
	c.hooks.Song = append(c.hooks.Song, hooks...)
}

// Create returns a builder for creating a Song entity.
func (c *SongClient) Create() *SongCreate {
	mutation := newSongMutation(c.config, OpCreate)
	return &SongCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Song entities.
func (c *SongClient) CreateBulk(builders ...*SongCreate) *SongCreateBulk {
	return &SongCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Song.
func (c *SongClient) Update() *SongUpdate {
	mutation := newSongMutation(c.config, OpUpdate)
	return &SongUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SongClient) UpdateOne(s *Song) *SongUpdateOne {
	mutation := newSongMutation(c.config, OpUpdateOne, withSong(s))
	return &SongUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SongClient) UpdateOneID(id int) *SongUpdateOne {
	mutation := newSongMutation(c.config, OpUpdateOne, withSongID(id))
	return &SongUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Song.
func (c *SongClient) Delete() *SongDelete {
	mutation := newSongMutation(c.config, OpDelete)
	return &SongDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SongClient) DeleteOne(s *Song) *SongDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *SongClient) DeleteOneID(id int) *SongDeleteOne {
	builder := c.Delete().Where(song.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SongDeleteOne{builder}
}

// Query returns a query builder for Song.
func (c *SongClient) Query() *SongQuery {
	return &SongQuery{
		config: c.config,
	}
}

// Get returns a Song entity by its id.
func (c *SongClient) Get(ctx context.Context, id int) (*Song, error) {
	return c.Query().Where(song.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SongClient) GetX(ctx context.Context, id int) *Song {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAsset queries the asset edge of a Song.
func (c *SongClient) QueryAsset(s *Song) *AssetQuery {
	query := &AssetQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(song.Table, song.FieldID, id),
			sqlgraph.To(asset.Table, asset.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, song.AssetTable, song.AssetColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SongClient) Hooks() []Hook {
	return c.hooks.Song
}

// TokenClient is a client for the Token schema.
type TokenClient struct {
	config
}

// NewTokenClient returns a client for the Token from the given config.
func NewTokenClient(c config) *TokenClient {
	return &TokenClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `token.Hooks(f(g(h())))`.
func (c *TokenClient) Use(hooks ...Hook) {
	c.hooks.Token = append(c.hooks.Token, hooks...)
}

// Create returns a builder for creating a Token entity.
func (c *TokenClient) Create() *TokenCreate {
	mutation := newTokenMutation(c.config, OpCreate)
	return &TokenCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Token entities.
func (c *TokenClient) CreateBulk(builders ...*TokenCreate) *TokenCreateBulk {
	return &TokenCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Token.
func (c *TokenClient) Update() *TokenUpdate {
	mutation := newTokenMutation(c.config, OpUpdate)
	return &TokenUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TokenClient) UpdateOne(t *Token) *TokenUpdateOne {
	mutation := newTokenMutation(c.config, OpUpdateOne, withToken(t))
	return &TokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TokenClient) UpdateOneID(id int) *TokenUpdateOne {
	mutation := newTokenMutation(c.config, OpUpdateOne, withTokenID(id))
	return &TokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Token.
func (c *TokenClient) Delete() *TokenDelete {
	mutation := newTokenMutation(c.config, OpDelete)
	return &TokenDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TokenClient) DeleteOne(t *Token) *TokenDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *TokenClient) DeleteOneID(id int) *TokenDeleteOne {
	builder := c.Delete().Where(token.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TokenDeleteOne{builder}
}

// Query returns a query builder for Token.
func (c *TokenClient) Query() *TokenQuery {
	return &TokenQuery{
		config: c.config,
	}
}

// Get returns a Token entity by its id.
func (c *TokenClient) Get(ctx context.Context, id int) (*Token, error) {
	return c.Query().Where(token.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TokenClient) GetX(ctx context.Context, id int) *Token {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Token.
func (c *TokenClient) QueryUser(t *Token) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(token.Table, token.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, token.UserTable, token.UserColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TokenClient) Hooks() []Hook {
	return c.hooks.Token
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRecord queries the record edge of a User.
func (c *UserClient) QueryRecord(u *User) *RankingQuery {
	query := &RankingQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(ranking.Table, ranking.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.RecordTable, user.RecordColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryToken queries the token edge of a User.
func (c *UserClient) QueryToken(u *User) *TokenQuery {
	query := &TokenQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(token.Table, token.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, user.TokenTable, user.TokenColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

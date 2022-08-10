// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-ranking-api/ent/ranking"
	"go-ranking-api/ent/user"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Ranking is the model entity for the Ranking schema.
type Ranking struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Score holds the value of the "score" field.
	Score int64 `json:"score,omitempty"`
	// SongUUID holds the value of the "song_uuid" field.
	SongUUID uuid.UUID `json:"song_uuid,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RankingQuery when eager-loading is set.
	Edges       RankingEdges `json:"edges"`
	user_record *int
}

// RankingEdges holds the relations/edges for other nodes in the graph.
type RankingEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RankingEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Ranking) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case ranking.FieldID, ranking.FieldScore:
			values[i] = new(sql.NullInt64)
		case ranking.FieldCreatedAt, ranking.FieldUpdatedAt, ranking.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case ranking.FieldSongUUID:
			values[i] = new(uuid.UUID)
		case ranking.ForeignKeys[0]: // user_record
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Ranking", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Ranking fields.
func (r *Ranking) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ranking.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case ranking.FieldScore:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field score", values[i])
			} else if value.Valid {
				r.Score = value.Int64
			}
		case ranking.FieldSongUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field song_uuid", values[i])
			} else if value != nil {
				r.SongUUID = *value
			}
		case ranking.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case ranking.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case ranking.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				r.DeletedAt = value.Time
			}
		case ranking.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_record", value)
			} else if value.Valid {
				r.user_record = new(int)
				*r.user_record = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Ranking entity.
func (r *Ranking) QueryUser() *UserQuery {
	return (&RankingClient{config: r.config}).QueryUser(r)
}

// Update returns a builder for updating this Ranking.
// Note that you need to call Ranking.Unwrap() before calling this method if this Ranking
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Ranking) Update() *RankingUpdateOne {
	return (&RankingClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Ranking entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Ranking) Unwrap() *Ranking {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Ranking is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Ranking) String() string {
	var builder strings.Builder
	builder.WriteString("Ranking(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("score=")
	builder.WriteString(fmt.Sprintf("%v", r.Score))
	builder.WriteString(", ")
	builder.WriteString("song_uuid=")
	builder.WriteString(fmt.Sprintf("%v", r.SongUUID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(r.DeletedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Rankings is a parsable slice of Ranking.
type Rankings []*Ranking

func (r Rankings) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Ranking holds the schema definition for the Ranking entity.
type Ranking struct {
	ent.Schema
}

// Fields of the Ranking.
func (Ranking) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("score").NonNegative(),
		field.UUID("song_uuid", uuid.UUID{}),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the Ranking.
func (Ranking) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("record").Unique(),
	}
}

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Asset holds the schema definition for the Asset entity.
type Asset struct {
	ent.Schema
}

func (Asset) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("platform").Edges("song").Unique(),
	}
}

// Fields of the Asset.
func (Asset) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("platform").Values("WINDOWS", "IOS").Default("WINDOWS"),
		field.String("path").NotEmpty(),
		field.String("hash").NotEmpty().Unique(),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional(),
	}
}

// Edges of the Asset.
func (Asset) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("song", Song.Type).Ref("asset").Required().Unique(),
	}
}

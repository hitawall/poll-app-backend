// file: ent/schema/polloption.go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PollOption holds the schema definition for the Option entity.
type PollOption struct {
	ent.Schema
}

// Fields of the Option.
func (PollOption) Fields() []ent.Field {
	return []ent.Field{
		field.String("text"),
		field.Int("vote_count").Default(0),
	}
}

// Edges of the Option.
func (PollOption) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("poll", Poll.Type).Ref("polloptions").Unique(),
		edge.To("votes", Vote.Type),
	}
}

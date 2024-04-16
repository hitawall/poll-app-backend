// file: ent/schema/vote.go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Vote holds the schema definition for the Vote entity.
type Vote struct {
	ent.Schema
}

// Fields of the Vote.
func (Vote) Fields() []ent.Field {
	return []ent.Field{
		field.Time("voted_on"), // Timestamp when the vote was cast
	}
}

// Edges of the Vote.
func (Vote) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("votes").Unique(), // The user who voted
		edge.To("polloption", PollOption.Type),             // The option that was voted on
	}
}

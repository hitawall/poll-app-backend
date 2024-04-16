// file: ent/schema/poll.go

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Poll holds the schema definition for the Poll entity.
type Poll struct {
	ent.Schema
}

// Fields of the Poll.
func (Poll) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
	}
}

// Edges of the Poll.
func (Poll) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("polloptions", PollOption.Type),
		edge.From("creator", User.Type).Ref("polls").Unique(),
	}
}

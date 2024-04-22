// Code generated by ent, DO NOT EDIT.

package vote

import (
	"poll-app-backend/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Vote {
	return predicate.Vote(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Vote {
	return predicate.Vote(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Vote {
	return predicate.Vote(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Vote {
	return predicate.Vote(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Vote {
	return predicate.Vote(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Vote {
	return predicate.Vote(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Vote {
	return predicate.Vote(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Vote {
	return predicate.Vote(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Vote {
	return predicate.Vote(sql.FieldLTE(FieldID, id))
}

// VotedOn applies equality check predicate on the "voted_on" field. It's identical to VotedOnEQ.
func VotedOn(v time.Time) predicate.Vote {
	return predicate.Vote(sql.FieldEQ(FieldVotedOn, v))
}

// VotedOnEQ applies the EQ predicate on the "voted_on" field.
func VotedOnEQ(v time.Time) predicate.Vote {
	return predicate.Vote(sql.FieldEQ(FieldVotedOn, v))
}

// VotedOnNEQ applies the NEQ predicate on the "voted_on" field.
func VotedOnNEQ(v time.Time) predicate.Vote {
	return predicate.Vote(sql.FieldNEQ(FieldVotedOn, v))
}

// VotedOnIn applies the In predicate on the "voted_on" field.
func VotedOnIn(vs ...time.Time) predicate.Vote {
	return predicate.Vote(sql.FieldIn(FieldVotedOn, vs...))
}

// VotedOnNotIn applies the NotIn predicate on the "voted_on" field.
func VotedOnNotIn(vs ...time.Time) predicate.Vote {
	return predicate.Vote(sql.FieldNotIn(FieldVotedOn, vs...))
}

// VotedOnGT applies the GT predicate on the "voted_on" field.
func VotedOnGT(v time.Time) predicate.Vote {
	return predicate.Vote(sql.FieldGT(FieldVotedOn, v))
}

// VotedOnGTE applies the GTE predicate on the "voted_on" field.
func VotedOnGTE(v time.Time) predicate.Vote {
	return predicate.Vote(sql.FieldGTE(FieldVotedOn, v))
}

// VotedOnLT applies the LT predicate on the "voted_on" field.
func VotedOnLT(v time.Time) predicate.Vote {
	return predicate.Vote(sql.FieldLT(FieldVotedOn, v))
}

// VotedOnLTE applies the LTE predicate on the "voted_on" field.
func VotedOnLTE(v time.Time) predicate.Vote {
	return predicate.Vote(sql.FieldLTE(FieldVotedOn, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Vote {
	return predicate.Vote(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Vote {
	return predicate.Vote(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPolloption applies the HasEdge predicate on the "polloption" edge.
func HasPolloption() predicate.Vote {
	return predicate.Vote(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, PolloptionTable, PolloptionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPolloptionWith applies the HasEdge predicate on the "polloption" edge with a given conditions (other predicates).
func HasPolloptionWith(preds ...predicate.PollOption) predicate.Vote {
	return predicate.Vote(func(s *sql.Selector) {
		step := newPolloptionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Vote) predicate.Vote {
	return predicate.Vote(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Vote) predicate.Vote {
	return predicate.Vote(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Vote) predicate.Vote {
	return predicate.Vote(sql.NotPredicates(p))
}

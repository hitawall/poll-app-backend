// Code generated by ent, DO NOT EDIT.

package vote

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the vote type in the database.
	Label = "vote"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVotedOn holds the string denoting the voted_on field in the database.
	FieldVotedOn = "voted_on"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgePolloption holds the string denoting the polloption edge name in mutations.
	EdgePolloption = "polloption"
	// Table holds the table name of the vote in the database.
	Table = "votes"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "votes"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_votes"
	// PolloptionTable is the table that holds the polloption relation/edge.
	PolloptionTable = "votes"
	// PolloptionInverseTable is the table name for the PollOption entity.
	// It exists in this package in order to avoid circular dependency with the "polloption" package.
	PolloptionInverseTable = "poll_options"
	// PolloptionColumn is the table column denoting the polloption relation/edge.
	PolloptionColumn = "vote_polloption"
)

// Columns holds all SQL columns for vote fields.
var Columns = []string{
	FieldID,
	FieldVotedOn,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "votes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_votes",
	"vote_polloption",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Vote queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByVotedOn orders the results by the voted_on field.
func ByVotedOn(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVotedOn, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByPolloptionField orders the results by polloption field.
func ByPolloptionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPolloptionStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newPolloptionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PolloptionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, PolloptionTable, PolloptionColumn),
	)
}

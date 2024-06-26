// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// EdgePolls holds the string denoting the polls edge name in mutations.
	EdgePolls = "polls"
	// EdgeVotes holds the string denoting the votes edge name in mutations.
	EdgeVotes = "votes"
	// Table holds the table name of the user in the database.
	Table = "users"
	// PollsTable is the table that holds the polls relation/edge.
	PollsTable = "polls"
	// PollsInverseTable is the table name for the Poll entity.
	// It exists in this package in order to avoid circular dependency with the "poll" package.
	PollsInverseTable = "polls"
	// PollsColumn is the table column denoting the polls relation/edge.
	PollsColumn = "user_polls"
	// VotesTable is the table that holds the votes relation/edge.
	VotesTable = "votes"
	// VotesInverseTable is the table name for the Vote entity.
	// It exists in this package in order to avoid circular dependency with the "vote" package.
	VotesInverseTable = "votes"
	// VotesColumn is the table column denoting the votes relation/edge.
	VotesColumn = "user_votes"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldName,
	FieldPassword,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByPollsCount orders the results by polls count.
func ByPollsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPollsStep(), opts...)
	}
}

// ByPolls orders the results by polls terms.
func ByPolls(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPollsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByVotesCount orders the results by votes count.
func ByVotesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newVotesStep(), opts...)
	}
}

// ByVotes orders the results by votes terms.
func ByVotes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVotesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPollsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PollsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PollsTable, PollsColumn),
	)
}
func newVotesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VotesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, VotesTable, VotesColumn),
	)
}

// Code generated by ent, DO NOT EDIT.

package polloption

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the polloption type in the database.
	Label = "poll_option"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldText holds the string denoting the text field in the database.
	FieldText = "text"
	// FieldVoteCount holds the string denoting the vote_count field in the database.
	FieldVoteCount = "vote_count"
	// EdgePoll holds the string denoting the poll edge name in mutations.
	EdgePoll = "poll"
	// EdgeVotes holds the string denoting the votes edge name in mutations.
	EdgeVotes = "votes"
	// Table holds the table name of the polloption in the database.
	Table = "poll_options"
	// PollTable is the table that holds the poll relation/edge.
	PollTable = "poll_options"
	// PollInverseTable is the table name for the Poll entity.
	// It exists in this package in order to avoid circular dependency with the "poll" package.
	PollInverseTable = "polls"
	// PollColumn is the table column denoting the poll relation/edge.
	PollColumn = "poll_polloptions"
	// VotesTable is the table that holds the votes relation/edge. The primary key declared below.
	VotesTable = "poll_option_votes"
	// VotesInverseTable is the table name for the Vote entity.
	// It exists in this package in order to avoid circular dependency with the "vote" package.
	VotesInverseTable = "votes"
)

// Columns holds all SQL columns for polloption fields.
var Columns = []string{
	FieldID,
	FieldText,
	FieldVoteCount,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "poll_options"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"poll_polloptions",
}

var (
	// VotesPrimaryKey and VotesColumn2 are the table columns denoting the
	// primary key for the votes relation (M2M).
	VotesPrimaryKey = []string{"poll_option_id", "vote_id"}
)

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

var (
	// DefaultVoteCount holds the default value on creation for the "vote_count" field.
	DefaultVoteCount int
)

// OrderOption defines the ordering options for the PollOption queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByText orders the results by the text field.
func ByText(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldText, opts...).ToFunc()
}

// ByVoteCount orders the results by the vote_count field.
func ByVoteCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVoteCount, opts...).ToFunc()
}

// ByPollField orders the results by poll field.
func ByPollField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPollStep(), sql.OrderByField(field, opts...))
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
func newPollStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PollInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, PollTable, PollColumn),
	)
}
func newVotesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VotesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, VotesTable, VotesPrimaryKey...),
	)
}

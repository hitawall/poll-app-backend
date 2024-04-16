// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"poll-app-backend/ent/poll"
	"poll-app-backend/ent/polloption"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// PollOption is the model entity for the PollOption schema.
type PollOption struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Votes holds the value of the "votes" field.
	Votes int `json:"votes,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PollOptionQuery when eager-loading is set.
	Edges            PollOptionEdges `json:"edges"`
	poll_polloptions *int
	vote_polloption  *int
	selectValues     sql.SelectValues
}

// PollOptionEdges holds the relations/edges for other nodes in the graph.
type PollOptionEdges struct {
	// Poll holds the value of the poll edge.
	Poll *Poll `json:"poll,omitempty"`
	// VotedBy holds the value of the voted_by edge.
	VotedBy []*User `json:"voted_by,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PollOrErr returns the Poll value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PollOptionEdges) PollOrErr() (*Poll, error) {
	if e.Poll != nil {
		return e.Poll, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: poll.Label}
	}
	return nil, &NotLoadedError{edge: "poll"}
}

// VotedByOrErr returns the VotedBy value or an error if the edge
// was not loaded in eager-loading.
func (e PollOptionEdges) VotedByOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.VotedBy, nil
	}
	return nil, &NotLoadedError{edge: "voted_by"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PollOption) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case polloption.FieldID, polloption.FieldVotes:
			values[i] = new(sql.NullInt64)
		case polloption.FieldText:
			values[i] = new(sql.NullString)
		case polloption.ForeignKeys[0]: // poll_polloptions
			values[i] = new(sql.NullInt64)
		case polloption.ForeignKeys[1]: // vote_polloption
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PollOption fields.
func (po *PollOption) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case polloption.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			po.ID = int(value.Int64)
		case polloption.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				po.Text = value.String
			}
		case polloption.FieldVotes:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field votes", values[i])
			} else if value.Valid {
				po.Votes = int(value.Int64)
			}
		case polloption.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field poll_polloptions", value)
			} else if value.Valid {
				po.poll_polloptions = new(int)
				*po.poll_polloptions = int(value.Int64)
			}
		case polloption.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field vote_polloption", value)
			} else if value.Valid {
				po.vote_polloption = new(int)
				*po.vote_polloption = int(value.Int64)
			}
		default:
			po.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PollOption.
// This includes values selected through modifiers, order, etc.
func (po *PollOption) Value(name string) (ent.Value, error) {
	return po.selectValues.Get(name)
}

// QueryPoll queries the "poll" edge of the PollOption entity.
func (po *PollOption) QueryPoll() *PollQuery {
	return NewPollOptionClient(po.config).QueryPoll(po)
}

// QueryVotedBy queries the "voted_by" edge of the PollOption entity.
func (po *PollOption) QueryVotedBy() *UserQuery {
	return NewPollOptionClient(po.config).QueryVotedBy(po)
}

// Update returns a builder for updating this PollOption.
// Note that you need to call PollOption.Unwrap() before calling this method if this PollOption
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *PollOption) Update() *PollOptionUpdateOne {
	return NewPollOptionClient(po.config).UpdateOne(po)
}

// Unwrap unwraps the PollOption entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *PollOption) Unwrap() *PollOption {
	_tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: PollOption is not a transactional entity")
	}
	po.config.driver = _tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *PollOption) String() string {
	var builder strings.Builder
	builder.WriteString("PollOption(")
	builder.WriteString(fmt.Sprintf("id=%v, ", po.ID))
	builder.WriteString("text=")
	builder.WriteString(po.Text)
	builder.WriteString(", ")
	builder.WriteString("votes=")
	builder.WriteString(fmt.Sprintf("%v", po.Votes))
	builder.WriteByte(')')
	return builder.String()
}

// PollOptions is a parsable slice of PollOption.
type PollOptions []*PollOption

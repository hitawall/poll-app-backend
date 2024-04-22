package handlers

import (
	"net/http"
	"poll-app-backend/ent"
	"poll-app-backend/ent/polloption"
	"poll-app-backend/ent/user"
	"poll-app-backend/ent/vote"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func VoteOption(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		optionID, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Invalid option ID", http.StatusBadRequest)
			return
		}

		u, ok := r.Context().Value("user").(*ent.User)
		if !ok {
			http.Error(w, "Authentication required", http.StatusUnauthorized)
			return
		}

		// Check if the user has already voted for this option
		exists, err := client.Vote.
			Query().
			Where(
				vote.HasPolloptionWith(polloption.IDEQ(optionID)),
				vote.HasUserWith(user.IDEQ(u.ID)),
			).
			Exist(r.Context())
		if err != nil {
			http.Error(w, "Failed to check existing votes", http.StatusInternalServerError)
			return
		}
		if exists {
			http.Error(w, "You have already voted for this option", http.StatusForbidden)
			return
		}

		// Start a transaction
		tx, err := client.Tx(r.Context())
		if err != nil {
			http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
			return
		}

		// Create the vote
		_, err = tx.Vote.
			Create().
			SetUser(u).
			SetPolloptionID(optionID).
			SetVotedOn(time.Now()). // Setting the voted_on timestamp
			Save(r.Context())
		if err != nil {
			tx.Rollback()
			http.Error(w, "Failed to record vote", http.StatusInternalServerError)
			return
		}

		// Increment votes count on the PollOption
		if _, err := tx.PollOption.
			Update().
			Where(polloption.IDEQ(optionID)).
			AddVotes(1).
			Save(r.Context()); err != nil {
			tx.Rollback()
			http.Error(w, "Failed to increment vote count", http.StatusInternalServerError)
			return
		}

		// Commit the transaction
		if err := tx.Commit(); err != nil {
			http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

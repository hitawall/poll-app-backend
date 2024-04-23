package handlers

import (
	log "github.com/sirupsen/logrus"
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

		log.Printf("Creating vote for user %d on option %d", u.ID, optionID)

		p, err := client.PollOption.
			Query().
			Where(polloption.IDEQ(optionID)).
			Only(r.Context())

		// Create the vote
		_, err = tx.Vote.
			Create().
			AddUser(u).
			AddPolloption(p).
			SetVotedOn(time.Now()). // Setting the voted_on timestamp
			Save(r.Context())
		if err != nil {
			tx.Rollback()
			http.Error(w, "Failed to record vote", http.StatusInternalServerError)
			return
		}

		// Optionally update the PollOption to include the new voter manually
		_, err = tx.PollOption.
			UpdateOneID(optionID).
			AddVotedBy(u).
			Save(r.Context())
		if err != nil {
			tx.Rollback()
			http.Error(w, "Failed to update PollOption voters", http.StatusInternalServerError)
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

func DeVoteOption(client *ent.Client) httprouter.Handle {
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

		// Start a transaction
		tx, err := client.Tx(r.Context())
		if err != nil {
			http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
			return
		}

		// Check if the user has voted for this option
		v, err := tx.Vote.
			Query().
			Where(
				vote.HasUserWith(user.IDEQ(u.ID)),
				vote.HasPolloptionWith(polloption.IDEQ(optionID)),
			).
			Only(r.Context())
		if err != nil {
			tx.Rollback() // Rollback transaction on error
			http.Error(w, "Failed to check vote", http.StatusInternalServerError)
			return
		}

		if v == nil {
			tx.Rollback() // Rollback transaction if user has not voted for this option
			http.Error(w, "You have not voted for this option", http.StatusForbidden)
			return
		}

		// Delete the vote
		err = tx.Vote.DeleteOne(v).Exec(r.Context())
		if err != nil {
			tx.Rollback() // Rollback transaction on error
			http.Error(w, "Failed to delete vote", http.StatusInternalServerError)
			return
		}

		// Decrement votes count on the option
		_, err = tx.PollOption.
			UpdateOneID(optionID).
			AddVotes(-1).
			Save(r.Context())
		if err != nil {
			tx.Rollback() // Rollback transaction on error
			http.Error(w, "Failed to decrement vote count", http.StatusInternalServerError)
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

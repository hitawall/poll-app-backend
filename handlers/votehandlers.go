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
		v, err := client.Vote.
			Create().
			SetUser(u).
			AddPolloption(p).
			SetVotedOn(time.Now()). // Setting the voted_on timestamp
			Save(r.Context())

		if err != nil {
			tx.Rollback()
			http.Error(w, "Failed to record vote", http.StatusInternalServerError)
			return
		}

		// Increment votes count on the PollOption
		if _, err := tx.PollOption.
			UpdateOne(p).
			AddVotes(v).
			AddVoteCount(1).
			Save(r.Context()); err != nil {
			log.Printf("Vote count error: %s", err)
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

		log.Printf("Removing vote for user %d on option %d", u.ID, optionID)

		p, err := client.PollOption.
			Query().
			Where(polloption.IDEQ(optionID)).
			WithVotes().
			Only(r.Context())

		if err != nil {
			http.Error(w, "Failed to fetch the poll option", http.StatusInternalServerError)
			return
		}

		userVote, err := p.QueryVotes().Where(
			vote.HasUserWith(user.IDEQ(u.ID)), // Filter to get only votes by this specific user
		).Only(r.Context())

		if userVote == nil {
			log.Println("No vote found for this user on the specified option")
			// You might want to handle this situation depending on your application's needs
			// For example, return an HTTP status to indicate no vote found
			http.Error(w, "No vote found for this user", http.StatusNotFound)
			return
		}

		// Remove Vote and decrement count on the PollOption
		if _, err := tx.PollOption.
			UpdateOne(p).
			RemoveVotes(userVote).
			AddVoteCount(-1).
			Save(r.Context()); err != nil {
			log.Printf("Vote count error: %s", err)
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

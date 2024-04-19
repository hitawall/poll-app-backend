package handlers

import (
	"encoding/json"
	"net/http"
	"poll-app-backend/ent"
	"poll-app-backend/ent/poll"
	_ "poll-app-backend/ent/poll"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CreatePollRequest struct {
	Title   string   `json:"title"`
	Options []string `json:"options"`
}

func CreatePoll(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req CreatePollRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Extract user from context (assuming your auth setup places a user entity there)
		user, ok := r.Context().Value("user").(*ent.User)
		if !ok {
			http.Error(w, "Authentication required", http.StatusUnauthorized)
			return
		}

		// Start a transaction
		tx, err := client.Tx(r.Context())
		if err != nil {
			http.Error(w, "Failed to start a transaction", http.StatusInternalServerError)
			return
		}

		// Create the poll with the creator associated
		p, err := tx.Poll.
			Create().
			SetTitle(req.Title).
			SetCreator(user).
			Save(r.Context())
		if err != nil {
			tx.Rollback()
			http.Error(w, "Failed to create poll", http.StatusInternalServerError)
			return
		}

		// Create poll options
		for _, optionText := range req.Options {
			_, err := tx.PollOption.
				Create().
				SetText(optionText).
				SetPoll(p).
				Save(r.Context())
			if err != nil {
				tx.Rollback()
				http.Error(w, "Failed to create poll option", http.StatusInternalServerError)
				return
			}
		}

		// Commit the transaction
		if err := tx.Commit(); err != nil {
			http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
			return
		}

		// Respond with the created poll data
		respondJSON(w, p, http.StatusCreated)
	}
}

func GetPolls(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		currentUser, _ := r.Context().Value("user").(*ent.User) // Get the user from context

		polls, err := client.Poll.
			Query().
			WithCreator().
			WithPolloptions().
			All(r.Context())
		if err != nil {
			http.Error(w, "Failed to retrieve polls", http.StatusInternalServerError)
			return
		}

		// Map polls to include 'createdByMe'
		results := make([]map[string]interface{}, len(polls))
		for i, p := range polls {
			results[i] = map[string]interface{}{
				"id":          p.ID,
				"title":       p.Title,
				"createdByMe": p.Edges.Creator != nil && currentUser != nil && p.Edges.Creator.ID == currentUser.ID,
			}
		}

		respondJSON(w, results, http.StatusOK)
	}
}

func GetPoll(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		pollID := ps.ByName("id") // Retrieve the poll ID from the URL parameter

		// Convert pollID from string to int
		id, err := strconv.Atoi(pollID)
		if err != nil {
			http.Error(w, "Invalid poll ID", http.StatusBadRequest)
			return
		}

		poll, err := client.Poll.
			Query().
			Where(poll.IDEQ(id)).
			WithCreator().
			WithPolloptions().
			Only(r.Context())

		if err != nil {
			if ent.IsNotFound(err) {
				http.Error(w, "Poll not found", http.StatusNotFound)
			} else {
				http.Error(w, "Failed to retrieve poll", http.StatusInternalServerError)
			}
			return
		}

		// Assuming there's a utility function to convert poll to a JSON response
		respondJSON(w, poll, http.StatusOK)
	}
}

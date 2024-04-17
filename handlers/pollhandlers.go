package handlers

import (
	"encoding/json"
	"net/http"
	"poll-app-backend/ent"
	_ "poll-app-backend/ent/poll"

	"github.com/julienschmidt/httprouter"
)

// CreatePoll creates a new poll with given title
func CreatePoll(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req struct {
			Title string `json:"title"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		p, err := client.Poll.
			Create().
			SetTitle(req.Title).
			Save(r.Context())
		if err != nil {
			http.Error(w, "Failed to create poll", http.StatusInternalServerError)
			return
		}
		respondJSON(w, p, http.StatusCreated)
	}
}

// GetPolls retrieves all polls with their options
func GetPolls(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		polls, err := client.Poll.
			Query().
			WithPolloptions().
			All(r.Context())
		if err != nil {
			http.Error(w, "Failed to retrieve polls", http.StatusInternalServerError)
			return
		}
		respondJSON(w, polls, http.StatusOK)
	}
}

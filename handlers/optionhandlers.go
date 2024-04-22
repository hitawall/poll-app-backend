package handlers

import (
	"encoding/json"
	"net/http"
	"poll-app-backend/ent"
	"poll-app-backend/ent/poll"
	"poll-app-backend/ent/polloption"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func AddOption(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		pollID, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Invalid poll ID", http.StatusBadRequest)
			return
		}

		user, ok := r.Context().Value("user").(*ent.User)
		if !ok {
			http.Error(w, "Authentication required", http.StatusUnauthorized)
			return
		}

		var req struct {
			Text string `json:"text"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Fetch the poll to verify the creator
		p, err := client.Poll.
			Query().
			Where(poll.IDEQ(pollID)).
			WithCreator().
			Only(r.Context())
		if err != nil {
			http.Error(w, "Failed to fetch poll", http.StatusInternalServerError)
			return
		}

		// Check if the current user is the creator of the poll
		if p.Edges.Creator != nil && p.Edges.Creator.ID != user.ID {
			http.Error(w, "Unauthorized to add options to this poll", http.StatusUnauthorized)
			return
		}

		// Create the option if user is authorized
		option, err := client.PollOption.
			Create().
			SetText(req.Text).
			SetPoll(p).
			Save(r.Context())
		if err != nil {
			http.Error(w, "Failed to create option", http.StatusInternalServerError)
			return
		}

		respondJSON(w, option, http.StatusCreated)
	}
}

// GetVoters lists all users who voted for a specific option
func GetVoters(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		optionID, _ := strconv.Atoi(ps.ByName("id"))
		option, err := client.PollOption.
			Query().
			Where(polloption.IDEQ(optionID)).
			WithVotedBy().
			Only(r.Context())
		if err != nil {
			http.Error(w, "Failed to retrieve voters", http.StatusInternalServerError)
			return
		}
		voters := option.QueryVotedBy().AllX(r.Context())
		respondJSON(w, voters, http.StatusOK)
	}
}

func UpdateOption(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		optionID, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Invalid option ID", http.StatusBadRequest)
			return
		}

		user, ok := r.Context().Value("user").(*ent.User)
		if !ok {
			http.Error(w, "Authentication required", http.StatusUnauthorized)
			return
		}

		var req struct {
			Text string `json:"text"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		opt, err := client.PollOption.
			Query().
			Where(polloption.IDEQ(optionID)).
			WithPoll(func(q *ent.PollQuery) {
				q.WithCreator()
			}).
			Only(r.Context())
		if err != nil {
			http.Error(w, "Failed to fetch option", http.StatusInternalServerError)
			return
		}

		if opt.Edges.Poll != nil && opt.Edges.Poll.Edges.Creator.ID != user.ID {
			http.Error(w, "Unauthorized to update this option", http.StatusUnauthorized)
			return
		}

		updatedOption, err := client.PollOption.
			UpdateOne(opt).
			SetText(req.Text).
			ClearVotedBy(). // This removes all voter connections if the text changes
			Save(r.Context())
		if err != nil {
			http.Error(w, "Failed to update option", http.StatusInternalServerError)
			return
		}

		respondJSON(w, updatedOption, http.StatusOK)
	}
}

func DeleteOption(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		optionID, err := strconv.Atoi(ps.ByName("id"))
		if err != nil {
			http.Error(w, "Invalid option ID", http.StatusBadRequest)
			return
		}

		user, ok := r.Context().Value("user").(*ent.User)
		if !ok {
			http.Error(w, "Authentication required", http.StatusUnauthorized)
			return
		}

		// Fetch the option along with its associated poll and the poll's creator
		opt, err := client.PollOption.
			Query().
			Where(polloption.IDEQ(optionID)).
			WithPoll(func(q *ent.PollQuery) {
				q.WithCreator()
			}).
			Only(r.Context())
		if err != nil {
			http.Error(w, "Failed to fetch option", http.StatusInternalServerError)
			return
		}

		// Check if the current user is the creator of the poll
		if opt.Edges.Poll == nil || opt.Edges.Poll.Edges.Creator == nil || opt.Edges.Poll.Edges.Creator.ID != user.ID {
			http.Error(w, "Unauthorized to delete this option", http.StatusUnauthorized)
			return
		}

		// Proceed to delete the option
		err = client.PollOption.
			DeleteOneID(optionID).
			Exec(r.Context())
		if err != nil {
			http.Error(w, "Failed to delete option", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Option deleted successfully"))
	}
}

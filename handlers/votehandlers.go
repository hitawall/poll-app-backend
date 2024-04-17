package handlers

import (
	"net/http"
	"poll-app-backend/ent"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// VoteOption increments the vote count for a given option
func VoteOption(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		optionID, _ := strconv.Atoi(ps.ByName("id"))
		_, err := client.PollOption.
			UpdateOneID(optionID).
			AddVotes(1). // Increment votes
			Save(r.Context())
		if err != nil {
			http.Error(w, "Failed to vote", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

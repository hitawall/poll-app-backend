package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"poll-app-backend/ent"
	"poll-app-backend/ent/user"
)

// Login handler
func LoginHandler(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		email := r.FormValue("email")
		password := r.FormValue("password")

		// Find user by email
		user, err := client.User.
			Query().
			Where(user.EmailEQ(email)).
			Only(r.Context())
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		// Compare the provided password with the stored one
		// NOTE: Store passwords securely using hashing and salting!
		if user.Password != password {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		// Generate a JWT token (implementation depends on your JWT library)

		w.WriteHeader(http.StatusOK)

		fmt.Fprintln(w, "Login successful")
	}
}

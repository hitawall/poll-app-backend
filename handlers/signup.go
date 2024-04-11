package handlers

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"poll-app-backend/ent"
)

// Signup handler
func SignupHandler(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		email := r.FormValue("email")
		name := r.FormValue("name")
		password := r.FormValue("password")

		// Create a new user
		_, err := client.User.
			Create().
			SetEmail(email).
			SetName(name).
			SetPassword(password). // NOTE: Passwords should be hashed and salted!
			Save(r.Context())
		if err != nil {
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			fmt.Println(err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "User created")
	}
}

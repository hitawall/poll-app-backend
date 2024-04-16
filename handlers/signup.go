package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"poll-app-backend/ent"
	_ "poll-app-backend/ent/user"
)

// SignupRequest defines the structure for incoming signup JSON payload
type SignupRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// SignupHandler creates a new user with a hashed password
func SignupHandler(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var req SignupRequest
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			fmt.Println("Error decoding JSON:", err)
			return
		}

		// Debugging output to check incoming data
		fmt.Printf("Received - Email: %s, Name: %s, Password: %s\n", req.Email, req.Name, req.Password)

		// Ensure there are no empty values unless expected
		if req.Email == "" || req.Name == "" || req.Password == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		// Hash the password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}

		// Create a new user with the hashed password
		_, err = client.User.
			Create().
			SetEmail(req.Email).
			SetName(req.Name).
			SetPassword(string(hashedPassword)).
			Save(r.Context())
		if err != nil {
			// Check if the error is due to a duplicate email
			if ent.IsConstraintError(err) {
				http.Error(w, "Email already in use", http.StatusConflict)
			} else {
				http.Error(w, "Error creating user", http.StatusInternalServerError)
			}
			fmt.Println("Error creating user:", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, `{"message":"User created successfully"}`)
	}
}

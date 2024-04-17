package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus" // Using structured logging
	"golang.org/x/crypto/bcrypt"
	"poll-app-backend/ent"
)

type SignupRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func SignupHandler(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var req SignupRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			log.Error("Error decoding JSON: ", err)
			return
		}

		if req.Email == "" || req.Name == "" || req.Password == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			log.Error("Error hashing password: ", err)
			return
		}

		_, err = client.User.
			Create().
			SetEmail(req.Email).
			SetName(req.Name).
			SetPassword(string(hashedPassword)).
			Save(r.Context())
		if err != nil {
			if ent.IsConstraintError(err) {
				http.Error(w, "Email already in use", http.StatusConflict)
			} else {
				http.Error(w, "Error creating user", http.StatusInternalServerError)
			}
			log.Error("Error creating user: ", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
	}
}

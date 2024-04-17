package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus" // Structured logging
	"golang.org/x/crypto/bcrypt"
	"poll-app-backend/ent"
	"poll-app-backend/ent/user"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func LoginHandler(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var creds Credentials
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			log.Error("Error decoding JSON: ", err)
			return
		}

		user, err := client.User.
			Query().
			Where(user.EmailEQ(creds.Email)).
			Only(r.Context())
		if err != nil {
			if ent.IsNotFound(err) {
				http.Error(w, "User not found", http.StatusNotFound)
			} else {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
			log.Error("Error fetching user: ", err)
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		claims := CustomClaims{
			Email: user.Email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
				Issuer:    "poll-app-backend",
			},
		}

		secretKey := "JWT_SECRET_KEY" // Ensure the environment variable is set
		if secretKey == "" {
			log.Fatal("JWT_SECRET_KEY must be set in the environment")
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(secretKey))
		if err != nil {
			http.Error(w, "Could not generate token", http.StatusInternalServerError)
			log.Error("Error generating token: ", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
	}
}

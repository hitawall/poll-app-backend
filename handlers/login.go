package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"poll-app-backend/ent"
	"poll-app-backend/ent/user"
)

// Credentials struct to hold incoming JSON data
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CustomClaims struct for JWT
type CustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

// LoginHandler authenticates a user using JSON input
func LoginHandler(client *ent.Client) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var creds Credentials
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			fmt.Println("Error decoding JSON:", err)
			return
		}

		fmt.Printf("Received - Email: %s, Password: %s\n", creds.Email, creds.Password)

		// Find user by email
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
			return
		}

		// Compare the provided password with the stored one
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
		if err != nil {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		// Create token with custom claims
		claims := CustomClaims{
			Email: user.Email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // Token expires in 24 hours
				Issuer:    "yourAppName",                         // Set the issuer of the token
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte("yourSecretKey")) // Use a secret from your environment variables
		if err != nil {
			http.Error(w, "Could not generate token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"token":"%s"}`, tokenString)
	}
}

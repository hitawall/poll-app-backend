package handlers

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"github.com/julienschmidt/httprouter"
	"poll-app-backend/ent"
	"poll-app-backend/ent/user"
)

const JWT_SECRET_KEY = "JWT_SECRET_KEY" // Use your actual secret key

func AuthMiddleware(client *ent.Client, next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(JWT_SECRET_KEY), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			userEmail := claims.Email
			if userEmail == "" {
				http.Error(w, "Email not found in token", http.StatusUnauthorized)
				return
			}

			log.Printf("Finding user with email id: %s", userEmail)
			log.Println(r.Context())
			log.Println(client)
			log.Println(client.User)
			log.Println(user.EmailEQ(userEmail))
			// Fetch user by email
			u, err := client.User.Query().Where(user.EmailEQ(userEmail)).Only(r.Context())
			if err != nil {
				log.Printf("User not found with email: %s", userEmail)
				http.Error(w, "User not found", http.StatusUnauthorized)
				return
			}

			// Store user in context
			ctx := context.WithValue(r.Context(), "user", u)
			next(w, r.WithContext(ctx), ps)
		} else {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		}
	}
}

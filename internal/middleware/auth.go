package middleware

import (
	"auth-service/internal/auth"
	"auth-service/internal/contextkey"
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")

			if tokenString == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			const bearerPrefix = "Bearer "
			if len(tokenString) < len(bearerPrefix) || tokenString[:len(bearerPrefix)] != bearerPrefix {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			tokenString = tokenString[len(bearerPrefix):]

			token, err := auth.ValidateJWT(tokenString, []byte(secret))
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			if !token.Valid {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			username, ok := claims["username"].(string)
			if !ok {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), contextkey.UsernameKey, username)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

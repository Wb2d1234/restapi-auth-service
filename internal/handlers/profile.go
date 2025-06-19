package handlers

import (
	"auth-service/internal/contextkey"
	"encoding/json"
	"net/http"
)

func ProfileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, ok := r.Context().Value(contextkey.UsernameKey).(string)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{
			"username": username,
			"message":  "Welcome to your profile!",
		})
	}
}

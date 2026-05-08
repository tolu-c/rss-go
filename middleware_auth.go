package main

import (
	"fmt"
	"net/http"

	"github.com/tolu-c/rss-go/internal/auth"
	"github.com/tolu-c/rss-go/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 401, fmt.Sprintf("Failed to get api key: %s", err))
			return
		}

		user, err := cfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Failed to get userL %v", err))
			return
		}

		handler(w, r, user)
	}
}

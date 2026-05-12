// Package response centralizes HTTP response writing for the api layer.
// Every handler that returns JSON should go through these helpers so the
// response shape (Content-Type, error envelope) stays consistent.
package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// errorBody is the canonical shape of error responses.
type errorBody struct {
	Error string `json:"error"`
}

// JSON writes payload as JSON with the given status code. If marshalling
// fails it logs and writes a bare 500 — this should not happen in practice
// because the response types in internal/api/model are always JSON-safe.
func JSON(w http.ResponseWriter, code int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("response: failed to marshal payload %v: %v", payload, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(data)
}

// Error writes a JSON error envelope. 5xx responses are logged because they
// indicate a server-side problem worth seeing in stdout.
func Error(w http.ResponseWriter, code int, message string) {
	if code >= 500 {
		log.Println("responding with 5xx error:", message)
	}
	JSON(w, code, errorBody{Error: message})
}

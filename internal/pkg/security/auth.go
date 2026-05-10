// Package security holds primitives related to authentication and authorization
// that are not tied to HTTP plumbing — pure functions that take headers /
// values and return verified information or an error.
package security

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey extracts the API key from an Authorization header of the form
//
//	Authorization: ApiKey <key>
//
// It returns an error if the header is missing or malformed.
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}

	parts := strings.Split(val, " ")
	if len(parts) != 2 {
		return "", errors.New("malformed authorization header")
	}
	if parts[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return parts[1], nil
}

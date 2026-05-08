package auth

import (
	"errors"
	"net/http"
	"strings"
)

// extracts apikey from the headers of an http request
// example:
// Authorization: ApiKey {insert apiKey here}
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("No authentication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed wrong header")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("malformed wrong header")
	}

	return vals[1], nil
}

// Auth middleware in case I need to limit requests
package api

import (
	"errors"
	"net/http"
	"strings"
)

// authMiddleware checks for the presence and validity of an authentication token.
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := extractToken(r)
		if err != nil {
			// Handle error, e.g., by sending an HTTP 401 Unauthorized response
			RespondWithError(w, http.StatusUnauthorized, err.Error())
			return
		}

		// Validate the token
		if !validateToken(token) {
			RespondWithError(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		next.ServeHTTP(w, r)
	})
}

// extractToken extracts the token from the request header.
func extractToken(r *http.Request) (string, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return "", errors.New("no token found")
	}
	// Strip the "Bearer " prefix from the token
	token = stripBearerPrefix(token)
	return token, nil
}

func stripBearerPrefix(token string) string {
	return strings.TrimPrefix(token, "Bearer ")
}

// validateToken validates the extracted token.
func validateToken(token string) bool {
	// TODO: Implement token validation
	return true
}

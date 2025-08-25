package jwt_pkg

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/jwtauth"
)

func ExtractUserIDFromToken(ctx context.Context) (string, error) {
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return "", err
	}

	userID, ok := claims["sub"].(string)
	if !ok || userID == "" {
		return "", errors.New("sub not found in token claims")
	}

	return userID, nil
}

func RequireRole(allowed ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, claims, err := jwtauth.FromContext(r.Context())

			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			role, ok := claims["role"].(string)
			if !ok || role == "" {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			for _, allowedRole := range allowed {
				if role == allowedRole {
					next.ServeHTTP(w, r)
					return
				}
			}

			http.Error(w, "Forbidden", http.StatusForbidden)
		})
	}
}

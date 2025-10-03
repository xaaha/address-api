package graph

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	"github.com/xaaha/address-api/internal/scripts"
)

type contextKey string

const apiKeyContextKey = contextKey("apikey")

// Auth is the directive implementation.
func Auth(ctx context.Context, _ any, next graphql.Resolver) (any, error) {
	// ideally this api would be fetched from the db, which sotres hashed keys, but for this small app, this should work
	validAPIKey := scripts.GetEnv().APIKey
	apiKey, ok := ctx.Value(apiKeyContextKey).(string)
	if !ok || apiKey == "" {
		return nil, fmt.Errorf("access denied: no API key provided")
	}
	if apiKey != validAPIKey {
		return nil, fmt.Errorf("access denied: invalid API key")
	}
	return next(ctx)
}

// AuthMiddleWare is the middleware funciton that checks the auth
func AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			// if no key, just proceed here. The Auth directive above will handle the error
			// We create a context with an empty value
			ctx := context.WithValue(r.Context(), apiKeyContextKey, "")
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		}
		token := strings.TrimPrefix(authHeader, "Bearer ")
		ctx := context.WithValue(r.Context(), apiKeyContextKey, token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

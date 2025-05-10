package auth

import (
	"HR-monitor/pkg/enums"
	"context"
	"errors"
	"net/http"
	"strings"
)

func JWTAuthMiddleware(jwtService *JWTService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "authorization header is required", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "invalid authorization header format", http.StatusUnauthorized)
				return
			}

			claims, err := jwtService.ValidateToken(parts[1])
			if err != nil {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), enums.UserIDKey, claims.UserID)
			ctx = context.WithValue(ctx, enums.UserRoleKey, claims.Role)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RequireRoles(roles ...enums.UserRole) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userRole := r.Context().Value(enums.UserRoleKey)
			if userRole == nil {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}

			for _, role := range roles {
				if userRole.(string) == string(role) {
					next.ServeHTTP(w, r)
					return
				}
			}

			http.Error(w, "forbidden", http.StatusForbidden)
		})
	}
}

func GetUserIDFromContext(ctx context.Context) (int, error) {
	userID, ok := ctx.Value(enums.UserIDKey).(int)
	if !ok {
		return 0, errors.New("user id not found in context")
	}
	return userID, nil
}

func GetUserRoleFromContext(ctx context.Context) (enums.UserRole, error) {
	role, ok := ctx.Value(enums.UserRoleKey).(string)
	if !ok {
		return "", errors.New("user role not found in context")
	}
	return enums.UserRole(role), nil
}

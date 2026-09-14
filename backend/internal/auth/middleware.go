package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/yahdee2702/chatting-aja/internal/httpx"
)

type AuthMiddleware struct {
	jwtHandler *JwtHandler
}

func NewAuthMiddleware(jwtHandler *JwtHandler) *AuthMiddleware {
	return &AuthMiddleware{
		jwtHandler: jwtHandler,
	}
}

func (m *AuthMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			httpx.Error(w, http.StatusBadRequest, "Invalid token")
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := m.jwtHandler.Verify(token)

		if err != nil {
			httpx.Error(w, http.StatusUnauthorized, "Token cannot be verified")
			return
		}

		ctx := context.WithValue(r.Context(), "userId", claims.Subject)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

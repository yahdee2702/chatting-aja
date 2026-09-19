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

		var token string

		if strings.Contains(authHeader, "Bearer") {
			token = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			cookieToken, err := r.Cookie("access_token")

			if err != nil {
				httpx.Error(w, http.StatusUnauthorized, "Token cannot be found")
				return
			}

			token = cookieToken.Value
		}

		claims, err := m.jwtHandler.Verify(token)

		if err != nil {
			httpx.Error(w, http.StatusUnauthorized, "Token cannot be verified")
			return
		}

		ctx := context.WithValue(r.Context(), "userId", claims.Subject)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

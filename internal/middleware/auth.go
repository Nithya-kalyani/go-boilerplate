package middleware

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Nithya-kalyani/go-boilerplate/pkg/utils"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.RenderJSON(w, http.StatusUnauthorized, map[string]string{"error": "Authorization header missing"})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			log.Printf("Invalid token: %v", err)
			utils.RenderJSON(w, http.StatusUnauthorized, map[string]string{"error": "Invalid or expired token"})
			return
		}

		next.ServeHTTP(w, r)
	})
}

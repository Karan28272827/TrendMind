package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"

	// "github.com/joho/godotenv"
	"os"
	"strings"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AuthHeader := r.Header.Get("Authorization")
		if AuthHeader == "" {
			http.Error(w, "Missing auth header", http.StatusUnauthorized)
			return
		}

		tokenArr := strings.Split(AuthHeader, "Bearer ")
		if len(tokenArr) != 2 {
			http.Error(w, "Invalid auth header format", http.StatusUnauthorized)
		}

		tokenStr := tokenArr[1]
		secret := os.Getenv("SECRET_KEY")
		if secret == "" {
			http.Error(w, "server config error", http.StatusInternalServerError)
			return
		}

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "JWT token is invalid", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

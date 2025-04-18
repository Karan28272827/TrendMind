package middleware

import (
	"net/http"

	"os"

	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Missing auth cookie", http.StatusUnauthorized)
			return
		}

		tokenStr := cookie.Value
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

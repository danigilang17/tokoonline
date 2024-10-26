package middleware

import (
	"net/http"

	"github.com/danigilang17/tokoonline/routes"
)

func JWTAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		tokenStr := cookie.Value
		claims := &routes.Claims{}
		token, err := routes.ParseToken(tokenStr, claims)

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Meneruskan claims ke handler berikutnya jika diperlukan
		r = r.WithContext(routes.SetUserIDContext(r.Context(), claims.UserID))

		next.ServeHTTP(w, r)
	}
}

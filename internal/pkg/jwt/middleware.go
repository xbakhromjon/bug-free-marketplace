package jwt

import (
	"context"
	"net/http"
)

func AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")

		if len(accessToken) == 0 {
			http.Error(w, "error:authorization header is not provided", http.StatusUnauthorized)
			return
		}

		claims, err := VerifyToken(accessToken)
		if err != nil {
			http.Error(w, "error"+err.Error(), http.StatusUnauthorized)
			return
		}

		_, ok := claims["sub"].(string)
		if !ok {
			http.Error(w, "error: Invalid subject in token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

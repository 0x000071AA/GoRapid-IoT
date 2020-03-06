package api

import (
	"log"
	"net/http"
)

type AuthenticationMiddleware struct {
	tokenUsers map[string]string
}

func (middleware *AuthenticationMiddleware) Populate() {
	middleware.tokenUsers["token"] = "user"
}

func (middleware *AuthenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")
		if user, found := middleware.tokenUsers[token]; found {
			next.ServeHTTP(w, r)
		} else {
			log.Printf("Forbidden for user %s\n", user)
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}

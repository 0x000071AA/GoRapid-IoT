package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("{ version : 1.0.0 }")
	})

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("{ health : true }")
	})

	r.HandleFunc("/users", GetUser).Methods("GET")
	r.HandleFunc("/users/{user}", CreateUser).Methods("POST")
	r.HandleFunc("/users/{user}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{user}", DeleteUser).Methods("DELETE")

	amw := AuthenticationMiddleware{tokenUsers: make(map[string]string)}
	//amw.Populate()

	r.Use(amw.Middleware)

}

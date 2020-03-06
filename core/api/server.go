package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/")

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode()
	})

	r.HandleFunc("/users").Methods("GET")
	r.HandleFunc("/users/{user}").Methods("POST")
	r.HandleFunc("/users/{user}").Methods("PUT")
	r.HandleFunc("/users/{user}").Methods("DELETE")

	amw := AuthenticationMiddleware{tokenUsers: make(map[string]string)}
	//amw.Populate()

	r.Use(amw.Middleware)
}

package main

import (
	"fmt"
	"net/http"

	"github.com/Parvis123/pokedex-apis/pokedex-apis/pokemon"
	"github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/pokemon/{name}", pokemon.GetPokemon).Methods("GET")

    r.Use(func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Set headers
            w.Header().Set("Access-Control-Allow-Origin", "*")
            w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
            w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

            // If it's a preflight request, respond with 200
            if r.Method == "OPTIONS" {
                w.WriteHeader(http.StatusOK)
                return
            }

            // Next
            next.ServeHTTP(w, r)
        })
    })

    fmt.Println("Server is now running on port 8000")
    http.ListenAndServe(":8000", r)
}
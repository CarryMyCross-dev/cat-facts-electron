package main

import (
    "fmt"
    "net/http"
    "github.com/CarryMyCross-dev/cat-facts-electron/backend/routes"
    "github.com/rs/cors"
)

func test_route(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, I am the test_route func")
}

func main() {
    mux := http.NewServeMux()
        routes.FetchCatFacts(mux)

        handler := cors.New(cors.Options{
            AllowedOrigins:   []string{"http://localhost:8081"},
            AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
            AllowedHeaders:   []string{"Content-Type"},
        }).Handler(mux)

        http.ListenAndServe(":8080", handler)
}
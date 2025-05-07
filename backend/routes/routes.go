package routes

import (
    "net/http"
    "github.com/CarryMyCross-dev/cat-facts-electron/backend/handlers"
)

func FetchCatFacts(mux *http.ServeMux) {
    mux.HandleFunc("/catfacts", handlers.CatFactsHandler)
}
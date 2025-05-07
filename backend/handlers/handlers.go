package handlers

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/CarryMyCross-dev/cat-facts-electron/backend/api"
)

func CatFactsHandler(w http.ResponseWriter, r *http.Request) {
    data, ex := catfacts.GetCatFacts()
    if ex != nil {
        log.Fatal(ex)
    }

    // convert Go object to json string
    formattedJson, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        log.Fatal(err)
    }

     // Set the response header to application/json
         // Add CORS headers
                w.Header().Set("Access-Control-Allow-Origin", "*") // Allow any origin
                w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
                w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        // Send the formatted JSON response
        w.Write(formattedJson)
}
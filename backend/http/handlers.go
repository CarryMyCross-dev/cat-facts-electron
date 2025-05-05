package handles

import (
    "fmt"
    "encoding/json"
    "log"
    "github.com/CarryMyCross-dev/cat-facts-electron/backend/api"
)

func CatFactsHandler() {
    data, ex := catfacts.GetCatFacts()
    if ex != nil {
        log.Fatal(ex)
    }

    // convert Go object to json string
    formattedJson, err := json.MarshalIndent(data, "", "  ")
    if err != nil {
        log.Fatal(err)
    }

    fmt.Print(string(formattedJson))
}
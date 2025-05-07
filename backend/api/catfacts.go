package catfacts

import (
    "fmt"
    "net/http"
    "os"
    "encoding/json"
)

type CatFactResponse struct {
    ID string   `json:"id"`
    URL string  `json:"url"`
    Width int   `json:"width"`
    Height   int           `json:"height"`
    Breeds   []interface{} `json:"breeds"`
    Favourite interface{}  `json:"favourite"`
}

func GetCatFacts() (CatFactResponse, error) {
    response, err := http.Get("https://api.thecatapi.com/v1/images/search")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    // close the response since we are completed with our request
    defer response.Body.Close()

    //decode the json
    var data []CatFactResponse
    if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
        fmt.Println("Error decoding JSON:", err)
        os.Exit(1)
    }

    // always going to have 1 item in the array. return it and the error
    return data[0], nil
}
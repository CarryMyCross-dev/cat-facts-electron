package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
    "os"
)

func main() {
    response, err := http.Get("https://api.thecatapi.com/v1/images/search")

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    defer response.Body.Close()

    responseData, err := io.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }


    fmt.Println(string(responseData))

}
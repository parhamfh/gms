package main

import (
    "fmt"
    "net/http"
    "log"
    "github.com/gorilla/mux"
)

func main() {
    http.HandleFunc("/adfad/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Received request")
        fmt.Fprintf(w, "Hey bro")
    })
    mux.NewRouter()
    log.Print(http.ListenAndServe(":8080", nil))
}

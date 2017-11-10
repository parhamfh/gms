package main

import (
    "fmt"
    "html"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)
func main() {
    router := mux.NewRouter().StrictSlash(true)
    // Routes
    router.HandleFunc("/", Index)
    // Listen
    log.Print(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Take care so long")
    fmt.Fprintf(w, "Hey, b %q", html.EscapeString(r.URL.Path))
}

package main

import (
    "encoding/json"
    "fmt"
    "html"
    "net/http"

    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    // fmt.Println("Received request", r)
    fmt.Fprintf(w, "Index request: ", html.EscapeString(r.URL.Path))
}

func TracksIndex(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Tracks request")

    tracks := Tracks{
        Track{Name: "Sonic Blast!"},
        Track{Name: "Pure Love"},
    }

    json.NewEncoder(w).Encode(tracks)
}

func TracksShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    trackId := vars["trackId"]
    fmt.Fprintln(w, "TracksShow request", trackId)
}

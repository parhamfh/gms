package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Received request", r)
	fmt.Fprintf(w, "Index request: ", html.EscapeString(r.URL.Path))
}

func TracksIndex(w http.ResponseWriter, r *http.Request) {
	rows, err := db_con.Query("SELECT * from tracks")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Empty array
	tracks := make([]Track, 0)

	var id int
	var title, artist string
	var releaseDate time.Time
	for rows.Next() {
		err := rows.Scan(&id, &title, &artist, &releaseDate)
		if err != nil {
			log.Fatal(err)
		}

		tracks = append(tracks, Track{Id: id, Title: title, Artist: artist, ReleaseDate: releaseDate})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)

	if err := json.NewEncoder(w).Encode(tracks); err != nil {
		panic(err)
	}
}

func TracksShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	trackId := vars["trackId"]
	fmt.Fprintln(w, "TracksShow request", trackId)
}

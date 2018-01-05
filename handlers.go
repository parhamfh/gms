package main

import (
	"database/sql"
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
		log.Fatal("TracksIndex: ", err)
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

	var id int
	var title, artist string
	var releaseDate time.Time

	err := db_con.QueryRow("SELECT * FROM tracks WHERE id = $1", trackId).Scan(&id, &title, &artist, &releaseDate)
	if err != nil {
		log.Println("TracksShow: ", err)
		if err == sql.ErrNoRows {
			w.WriteHeader(404)
			return
		}
		log.Fatal("TracksShow: Unexpected error.")
	}
	track := Track{id, title, artist, releaseDate}
	if err := json.NewEncoder(w).Encode(track); err != nil {
		panic(err)
	}
}

package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	log.Print(http.ListenAndServe(":8080", router))
}

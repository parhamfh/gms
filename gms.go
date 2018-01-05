package main

import (
    "log"
    "net/http"
)
func main() {
    ConnectDB()
    router := NewRouter()
    log.Print(http.ListenAndServe(":8080", router))
}

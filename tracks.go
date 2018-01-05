package main

import "time"

type Track struct {
    Id              int         `json:"id"`
	Title           string      `json:"title"`
	Artist          string      `json:"artist"`
	ReleaseDate     time.Time   `json:"releaseDate"`
}

type Tracks []Track

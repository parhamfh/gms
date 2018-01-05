package main

import "time"

type Track struct {
	Title           string      `json:"title"`
	Artist          string      `json:"artist"`
	ReleaseDate     time.Time   `json:"releaseDate"`
}

type Tracks []Track

package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "parham"
	DB_PASSWORD = ""
	DB_NAME     = "parham"
)

var db_con *sql.DB

func ConnectDB() {
	fmt.Println("## Connecting to PostgreSQL")
	db_params := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable",
		DB_USER,
		DB_PASSWORD,
		DB_NAME,
	)
	var err error
	db_con, err = sql.Open("postgres", db_params)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("## Connected to DB.")
}

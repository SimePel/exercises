package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "sime_pel"
	password = "jfkd"
	dbname   = "phone_numbers"
)

func main() {
	psInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psInfo)
	if err != nil {
		log.Fatal("open:", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("ping:", err)
	}
	fmt.Println("Successfully connected!")
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"unicode"

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

	phones := [8]string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}

	for _, v := range phones {
		_, err = db.Exec("INSERT INTO phones (phone) VALUES ($1);", v)
		if err != nil {
			log.Fatal("exec:", err)
		}
	}

	rows, err := db.Query("SELECT id, phone FROM phones;")
	if err != nil {
		log.Fatal("query:", err)
	}

	for rows.Next() {
		var id int
		var phone string
		err = rows.Scan(&id, &phone)
		if err != nil {
			log.Fatal("scan:", err)
		}

		updStatement := `UPDATE phones SET phone = $2 WHERE id = $1;`
		_, err = db.Exec(updStatement, id, Normalize(phone))
		if err != nil {
			// if err occurred, it means that we try to add duplicate,
			// because phone field is UNIQUE, that's why we remove this record from db.
			_, err = db.Exec("DELETE FROM phones WHERE id = $1;", id)
			if err != nil {
				log.Fatal("DELETE after failed UPDATE:", err)
			}
		}
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	rows.Close()
}

// Normalize converts random phone number in sequence of digits
func Normalize(phone string) (s string) {
	for _, c := range phone {
		if unicode.IsDigit(c) {
			s += string(c)
		}
	}
	return s
}

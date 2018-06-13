package main

import (
	"log"

	"github.com/SimePel/exercises/todo-cli/cmd"
	bolt "github.com/coreos/bbolt"
)

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("tasks"))
		return nil
	})
	db.Close()
	cmd.Execute()
}

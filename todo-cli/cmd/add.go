package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	bolt "github.com/coreos/bbolt"
	"github.com/spf13/cobra"
)

const addUsage = `
Usage:
  task add [here is your task]
`

func init() {
	addCmd.SetUsageTemplate(addUsage)
	addCmd.SetHelpTemplate(addCmd.Short + addUsage)
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := bolt.Open("my.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		task := strings.Join(args, " ")

		err = db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("tasks"))
			// NextSequence returns an error only if the tx is closed or not writeable.
			// That can't happen in an Update() call so I ignore the error check.
			id, _ := b.NextSequence()
			return b.Put([]byte(strconv.FormatUint(id, 10)), []byte(task))
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Added \"%s\" to your task list.", task)
	},
}

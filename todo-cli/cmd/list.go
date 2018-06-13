package cmd

import (
	"fmt"
	"log"

	bolt "github.com/coreos/bbolt"
	"github.com/spf13/cobra"
)

const listUsage = `
Usage:
  task list
`

func init() {
	listCmd.SetUsageTemplate(listUsage)
	listCmd.SetHelpTemplate(listCmd.Short + listUsage)
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := bolt.Open("my.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		db.View(func(tx *bolt.Tx) error {
			c := tx.Bucket([]byte("tasks")).Cursor()
			if k, _ := c.First(); k == nil {
				fmt.Println("You haven't any tasks.")
				return nil
			}
			fmt.Println("You have the following tasks:")
			i := 1
			for k, v := c.First(); k != nil; k, v = c.Next() {
				fmt.Printf("%v. %s\n", i, v)
				i++
			}
			return nil
		})
	},
}

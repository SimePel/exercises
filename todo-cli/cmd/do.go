package cmd

import (
	"fmt"
	"log"
	"strconv"

	bolt "github.com/coreos/bbolt"
	"github.com/spf13/cobra"
)

const doUsage = `
Usage:
  task do [index]
`

func init() {
	doCmd.SetUsageTemplate(doUsage)
	doCmd.SetHelpTemplate(doCmd.Short + doUsage)
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark a task on your TODO list as complete",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := bolt.Open("my.db", 0600, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		var task []byte
		err = db.Update(func(tx *bolt.Tx) error {
			pos, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("Invalid index")
			}
			b := tx.Bucket([]byte("tasks"))
			c := b.Cursor()
			n := getNumberOfRecords(c)
			if n < pos || pos < n {
				return fmt.Errorf("Index out of range")
			}

			k, v := c.First()
			for pos != 1 {
				k, v = c.Next()
				pos--
			}
			task = make([]byte, len(v))
			copy(task, v)
			return b.Delete([]byte(k))
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("You have completed the \"%s\" task.\n", task)
	},
}

// Don't know, why bolt hasn't this function, maybe I was badly looking for.
func getNumberOfRecords(c *bolt.Cursor) (i int) {
	for k, _ := c.First(); k != nil; k, _ = c.Next() {
		i++
	}
	return i
}

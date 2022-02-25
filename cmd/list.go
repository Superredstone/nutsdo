package cmd

import (
	"encoding/json"
	"fmt"

	nutsdoDB "github.com/Superredstone/nutsdo/db"
	"github.com/Superredstone/nutsdo/types"

	"github.com/spf13/cobra"
	"github.com/xujiajun/nutsdb"
)

func ListTodo(cmd *cobra.Command, args []string) {
	db := nutsdoDB.OpenDB()

	db.View(func(tx *nutsdb.Tx) error {
		entries, _ := tx.GetAll("db")

		for _, entry := range entries {
			var entryStruct types.Todo

			json.Unmarshal(entry.Value, &entryStruct)

			if entryStruct.Done {
				fmt.Println("[x] " + entryStruct.Name)
			} else {
				fmt.Println("[ ] " + entryStruct.Name)
			}
		}

		return nil
	})
}

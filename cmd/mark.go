package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Superredstone/nutsdo/types"

	nutsdoDB "github.com/Superredstone/nutsdo/db"

	"github.com/spf13/cobra"
	"github.com/xujiajun/nutsdb"
)

func MarkTodo(cmd *cobra.Command, args []string) {
	db := nutsdoDB.OpenDB()

	db.Update(func(tx *nutsdb.Tx) error {
		entry, err := tx.Get("db", []byte(strings.Join(args, " ")))
		if err != nil {
			fmt.Println("Can't find " + strings.Join(args, " "))
			os.Exit(0)
		}

		var entryStruct types.Todo

		json.Unmarshal(entry.Value, &entryStruct)

		if entryStruct.Done {
			entryStruct.Done = false

			marshalEntryStruct, _ := json.Marshal(entryStruct)

			err := tx.Put("db", []byte(entryStruct.Name), marshalEntryStruct, 0)
			if err != nil {
				tx.Rollback()
				os.Exit(1)
			}

			fmt.Println("[ ] " + entryStruct.Name)
		} else {
			entryStruct.Done = true

			marshalEntryStruct, _ := json.Marshal(entryStruct)

			err := tx.Put("db", []byte(entryStruct.Name), marshalEntryStruct, 0)
			if err != nil {
				tx.Rollback()
				os.Exit(1)
			}

			fmt.Println("[x] " + entryStruct.Name)
		}

		return nil
	})
}

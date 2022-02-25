package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Superredstone/nutsdo/types"

	nutsdoDB "github.com/Superredstone/nutsdo/db"

	"github.com/spf13/cobra"
)

func AddTodo(cmd *cobra.Command, args []string) {
	todoName := cmd.Flags().String("name", strings.Join(args, " "), "name of the todo")

	db := nutsdoDB.OpenDB()

	tx, err := db.Begin(true)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var todoObject = types.Todo{
		Name:     *todoName,
		Done:     false,
		Category: "none",
	}

	todoJSON, _ := json.Marshal(todoObject)

	if err = tx.Put("db", []byte(todoObject.Name), todoJSON, 0); err != nil {
		tx.Rollback()
	} else {
		if err = tx.Commit(); err != nil {
			tx.Rollback()
			os.Exit(1)
		}
	}
}

package cmd

import (
	"fmt"
	"os"
	"strings"

	nutsdoDB "github.com/Superredstone/nutsdo/db"

	"github.com/spf13/cobra"
	"github.com/xujiajun/nutsdb"
)

func DeleteTodo(cmd *cobra.Command, args []string) {
	db := nutsdoDB.OpenDB()

	db.Update(func(tx *nutsdb.Tx) error {
		_, err := tx.Get("db", []byte(strings.Join(args, " ")))
		if err != nil {
			fmt.Println("Can't find " + strings.Join(args, " "))
			os.Exit(0)
		}

		err = tx.Delete("db", []byte(strings.Join(args, " ")))
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}

		return nil
	})
}

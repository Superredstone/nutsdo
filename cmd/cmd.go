package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func Execute() {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "add",
		Short: "Add a new todo",
		Run:   AddTodo,
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:   "list",
		Short: "Lists all todos",
		Run:   ListTodo,
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:   "mark",
		Short: "Mark a todo as done or unmarks a todo",
		Run:   MarkTodo,
	})
	rootCmd.AddCommand(&cobra.Command{
		Use:   "del",
		Short: "Deletes a todo",
		Run:   DeleteTodo,
	})

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

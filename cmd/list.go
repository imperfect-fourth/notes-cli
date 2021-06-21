package cmd

import (
    "context"
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your todos",
	Run: func(cmd *cobra.Command, args []string) {

        var todos []todo
        var err error

        if showAll {
            todos, err = getAllTodos()
        } else {
            todos, err = getTodos()
        }

        if err != nil {
            fmt.Println(err)
            return
        }

        fmt.Println(todos)

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
    listCmd.Flags().BoolVar(&showAll, "all", false, "Completed todos will be shown if this flag is used")
}

var todoQuery struct {
    Todos []todo `graphql:"todos(where: {completed: {_eq: false}})"`
}

var allTodoQuery struct {
    Todos []todo
}

var showAll bool

func getTodos() ([]todo, error) {
    if err := client.Query(context.Background(), &todoQuery, nil); err != nil {
        fmt.Println(err)
        return nil, err
    }
    return todoQuery.Todos, nil
}

func getAllTodos() ([]todo, error) {
    if err := client.Query(context.Background(), &allTodoQuery, nil); err != nil {
        fmt.Println(err)
        return nil, err
    }
    return allTodoQuery.Todos, nil
}


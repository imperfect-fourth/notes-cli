package cmd

import (
    "context"
	"fmt"
    "os"

	"github.com/spf13/cobra"
    "text/tabwriter"
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

        prettyPrint(todos)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
    listCmd.Flags().BoolVar(&showAll, "all", false, "Completed todos will be shown if this flag is used")
}

var todoQuery struct {
    Todos []todo `graphql:"todos(order_by: {id: asc}, where: {completed: {_eq: false}})"`
}

var allTodoQuery struct {
    Todos []todo `graphql:"todos(order_by: {id: asc})"`
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

func prettyPrint(todos []todo) {
    w := new(tabwriter.Writer)
    w.Init(os.Stdout, 8, 8, 0, '\t', 0)

    defer w.Flush()

    fmt.Fprintf(w, "%s\t%s\t%s\t\n", "ID", "TODO", "COMPLETED")
    for i := 0; i<len(todos); i++ {
        fmt.Fprintf(w, "%d\t%s\t%t\t\n", int(todos[i].ID),
            todos[i].Body, bool(todos[i].Completed),
        )
    }
}

package cmd

import (
    "context"
	"fmt"
    "os"
    "strconv"

	"github.com/spf13/cobra"
    "github.com/hasura/go-graphql-client"
)

var toggleCmd = &cobra.Command{
	Use:   "toggle <id>",
	Short: "Toggle a todo's status",
	Run: func(cmd *cobra.Command, args []string) {
        if len(args) == 0 {
            cmd.Help()
            return
        }
        id, err := strconv.Atoi(args[0])
        if err != nil {
            fmt.Fprintln(os.Stderr, "Enter a valid ID")
            return
        }
        toggleTodo(id)
    },
}

var queryStatus struct {
	Todos struct {
		Completed graphql.Boolean
	} `graphql:"todos_by_pk(id: $id)"`
}

var mutateToggle struct {
    UpdateTodosByPk struct {
        ID graphql.Int
    } `graphql:"update_todos_by_pk(pk_columns: {id: $id}, _set: {completed: $status})"`
}

func toggleTodo(id int) {
	err := client.Query(context.Background(), &queryStatus,
		map[string]interface{}{"id":graphql.Int(id)},
	)
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
    }
	status := !bool(queryStatus.Todos.Completed)
    err = client.Mutate(context.Background(), &mutateToggle,
        map[string]interface{}{
			"id": graphql.Int(id),
			"status": graphql.Boolean(status),
		},
    )
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
    }

	if status {
		fmt.Println("Todo marked as completed!")
	} else {
		fmt.Println("Todo marked as not completed!")
	}
}

func init() {
	rootCmd.AddCommand(toggleCmd)
}

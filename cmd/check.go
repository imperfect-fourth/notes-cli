package cmd

import (
    "context"
	"fmt"
    "os"
    "strconv"

	"github.com/spf13/cobra"
    "github.com/hasura/go-graphql-client"
)

var checkCmd = &cobra.Command{
	Use:   "check <id>",
	Short: "Check off a todo from your list",
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
        checkTodo(id)
    },
}

var mutateCheck struct {
    UpdateTodosByPk struct {
        ID graphql.Int
    } `graphql:"update_todos_by_pk(pk_columns: {id: $id}, _set: {completed: true})"`
}

func checkTodo(id int) {
    err := client.Mutate(context.Background(), &mutateCheck,
        map[string]interface{}{"id": graphql.Int(id)},
    )
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
    }

    fmt.Println("Todo marked off!")
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

package cmd

import (
    "context"
	"fmt"
    "strings"

	"github.com/spf13/cobra"
    "github.com/hasura/go-graphql-client"
)

var editCmd = &cobra.Command{
	Use:   "edit --id <id> <body>",
	Short: "Edit your todo",
	Run: func(cmd *cobra.Command, args []string) {
        if len(args) == 0 {
            cmd.Help()
            return
        }
        editTodo(strings.Join(args, " "))
	},
}

var id int

var mutateEdit struct {
    UpdateTodoQuery struct {
        ID graphql.Int
    } `graphql:"update_todos_by_pk(pk_columns: {id: $id}, _set: {body: $body})"`
}

func editTodo(body string) {
    err := client.Mutate(context.Background(), &mutateEdit,
        map[string]interface{}{
            "id": graphql.Int(id),
            "body": graphql.String(body),
        },
    )

    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Successfully edited todo!")
}

func init() {
	rootCmd.AddCommand(editCmd)

    editCmd.Flags().IntVar(&id, "id", 0, "ID of the todo to edit")
    editCmd.MarkFlagRequired("id")
}

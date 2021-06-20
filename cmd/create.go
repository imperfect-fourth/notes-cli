package cmd

import (
    "context"
	"fmt"
    "strings"

	"github.com/spf13/cobra"
    "github.com/hasura/go-graphql-client"
)

var createCmd = &cobra.Command{
	Use:   "create [flags] body",
	Short: "Create a new todo",
	Run: func(cmd *cobra.Command, args []string) {
        if len(args) == 0 {
            cmd.Help()
            return
        }
        client := graphql.NewClient("http://localhost:8080/v1/graphql", nil);
        createTodo(client, strings.Join(args, " "))
	},
}

var createMutation struct {
    InsertTodosOne struct {
        ID graphql.Int
    } `graphql:"insert_todos_one(object: {body: $body})"`
}

func createTodo(client *graphql.Client, body string) {
    err := client.Mutate(context.Background(), &createMutation,
        map[string]interface{}{"body": graphql.String(body)},
    )
    if err != nil {
        fmt.Println(err)
        return;
    }
    fmt.Println("Todo successfully created!")
}

func init() {
	rootCmd.AddCommand(createCmd)
}

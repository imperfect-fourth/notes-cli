package cmd

import (
    "context"
	"fmt"

	"github.com/spf13/cobra"
    "github.com/hasura/go-graphql-client"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your todos",
	Run: func(cmd *cobra.Command, args []string) {
		getTodos()
	},
}


var todoQuery struct {
    Todos []struct {
        Title graphql.String
        Completed graphql.Boolean
    }
}


func getTodos() {
    client := graphql.NewClient("http://localhost:8080/v1/graphql", nil)
    err := client.Query(context.Background(), &todoQuery, nil)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(todoQuery)
    fmt.Println("conn")
}


func init() {
	rootCmd.AddCommand(listCmd)
}

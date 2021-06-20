package cmd

import (
    "github.com/hasura/go-graphql-client"
)

type todo struct {
    ID graphql.Int
    Body graphql.String
    Completed graphql.Boolean
}


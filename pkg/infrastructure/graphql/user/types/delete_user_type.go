package types

import (
	"github.com/graphql-go/graphql"
)

var DeleteUserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "DeleteUserOutput",
		Description: "User information",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)

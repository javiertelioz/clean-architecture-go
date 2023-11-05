package types

import (
	"github.com/graphql-go/graphql"
)

var CreateUserType = graphql.NewInputObject(
	graphql.InputObjectConfig{
		Name: "createUserInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"name": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"lastname": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"surname": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
			"email": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.InputObjectFieldConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"phone": &graphql.InputObjectFieldConfig{
				Type: graphql.String,
			},
		},
	},
)

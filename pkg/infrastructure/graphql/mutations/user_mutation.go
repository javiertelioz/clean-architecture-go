package mutations

import (
	"github.com/graphql-go/graphql"
)

var CreateUserMutation = &graphql.Field{
	Type: graphql.NewObject(graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
		},
	}),
	Description: "Crea un nuevo usuario",
	Args: graphql.FieldConfigArgument{
		"username": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return 1, nil
	},
}

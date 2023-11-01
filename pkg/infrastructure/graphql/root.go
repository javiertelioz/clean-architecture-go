package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/graphql/resolvers"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/graphql/schemas"
)

func NewSchema() graphql.Schema {
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: graphql.Fields{
		"getUserById": &graphql.Field{
			Type:        schemas.UserType,
			Description: "Get single user",
			Resolve:     resolvers.GetUser,
		},
	}}
	
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	})

	if err != nil {
		panic(err)
	}

	return schema
}

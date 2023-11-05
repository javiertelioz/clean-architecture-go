package graphql

import (
	"github.com/graphql-go/graphql"
)

func NewSchema() graphql.Schema {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    QueryType,
			Mutation: MutationType,
		},
	)
	if err != nil {
		panic(err)
	}

	return schema
}

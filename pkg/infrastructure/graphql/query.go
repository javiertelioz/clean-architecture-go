package graphql

import (
	"github.com/graphql-go/graphql"
	applicatioQuery "github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/graphql/application/query"
	userQuery "github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/graphql/user/query"
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"getUserById":               userQuery.GetUserByIDQuery(),
			"getApplicationInformation": applicatioQuery.GetApplicationQuery(),
		},
	})

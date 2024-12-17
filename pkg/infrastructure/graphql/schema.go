package graphql

import (
	"github.com/graphql-go/graphql"
	applicationQuery "github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/graphql/application/query"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/graphql/user/mutation"
	userQuery "github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/graphql/user/query"
)

func NewSchema(resolverRegistry *ResolverRegistry) graphql.Schema {
	userResolver := resolverRegistry.GetUserResolver()
	applicationResolver := resolverRegistry.GetApplicationResolver()

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graphql.NewObject(
				graphql.ObjectConfig{
					Name: "Query",
					Fields: graphql.Fields{
						"getUserById":               userQuery.GetUserByIDQuery(userResolver),
						"getApplicationInformation": applicationQuery.GetApplicationInformationQuery(applicationResolver),
					},
				},
			),
			Mutation: graphql.NewObject(
				graphql.ObjectConfig{
					Name: "Mutation",
					Fields: graphql.Fields{
						"createUser": mutation.CreateUserMutation(userResolver),
						"updateUser": mutation.UpdateUserMutation(userResolver),
						"deleteUser": mutation.DeleteUserMutation(userResolver),
					},
				},
			),
		},
	)
	if err != nil {
		panic(err)
	}

	return schema
}

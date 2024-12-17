package query

import (
	"github.com/graphql-go/graphql"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/graphql/user/resolve"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/graphql/user/types"
)

func GetUserByIDQuery(userResolver *resolve.UserResolver) *graphql.Field {
	return &graphql.Field{
		Type:        types.UserType,
		Description: "Get user by ID",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: userResolver.GetUserById,
	}
}

package mutation

import (
	"github.com/graphql-go/graphql"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/graphql/user/resolve"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/graphql/user/types"
)

func CreateUserMutation(userResolver *resolve.UserResolver) *graphql.Field {
	return &graphql.Field{
		Name:        "CreateUser",
		Type:        types.UserType,
		Description: "Create a new user",
		Args: graphql.FieldConfigArgument{
			"CreateUserInput": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(types.CreateUserType),
			},
		},
		Resolve: userResolver.CreateUser,
	}
}

func UpdateUserMutation(userResolver *resolve.UserResolver) *graphql.Field {
	return &graphql.Field{
		Name:        "UpdateUser",
		Type:        types.UserType,
		Description: "Update an existing user",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"UpdateUserInput": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(types.UpdateUserType),
			},
		},
		Resolve: userResolver.UpdateUser,
	}
}

func DeleteUserMutation(userResolver *resolve.UserResolver) *graphql.Field {
	return &graphql.Field{
		Name:        "DeleteUser",
		Type:        types.DeleteUserType,
		Description: "Delete an existing user",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: userResolver.DeleteUserById,
	}
}

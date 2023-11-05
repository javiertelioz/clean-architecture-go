package mutation

import (
	"github.com/graphql-go/graphql"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/graphql/user/resolve"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/graphql/user/types"
)

func CreateUserMutation() *graphql.Field {
	userResolve := resolve.NewUserResolver()

	return &graphql.Field{
		Name:        "CreateUser",
		Type:        types.UserType,
		Description: "Create a new user",
		Args: graphql.FieldConfigArgument{
			"CreateUserInput": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(types.CreateUserType),
			},
		},
		Resolve: userResolve.CreateUser,
	}
}

func UpdateUserMutation() *graphql.Field {
	userResolve := resolve.NewUserResolver()

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
		Resolve: userResolve.UpdateUser,
	}
}

func DeleteUserMutation() *graphql.Field {
	userResolve := resolve.NewUserResolver()

	return &graphql.Field{
		Name:        "DeleteUser",
		Type:        types.DeleteUserType,
		Description: "Delete an existing user",
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: userResolve.DeleteUserById,
	}
}

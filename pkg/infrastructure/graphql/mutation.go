package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/graphql/user/mutation"
)

var MutationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createUser": mutation.CreateUserMutation(),
			"updateUser": mutation.UpdateUserMutation(),
			"deleteUser": mutation.DeleteUserMutation(),
		},
	})

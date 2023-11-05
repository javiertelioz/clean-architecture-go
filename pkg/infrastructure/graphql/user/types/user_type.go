package types

import (
	"github.com/graphql-go/graphql"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/graphql/user/field"
)

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "UserOutput",
		Description: "User information",
		Fields:      field.UserField,
	},
)

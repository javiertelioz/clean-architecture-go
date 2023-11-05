package types

import (
	"github.com/graphql-go/graphql"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/graphql/application/field"
)

var ApplicationType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Application",
		Description: "Application information",
		Fields:      field.ApplicationField,
	},
)

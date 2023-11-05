package query

import (
	"github.com/graphql-go/graphql"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/graphql/application/resolve"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/graphql/application/types"
)

func GetApplicationQuery() *graphql.Field {
	applicationResolver := resolve.NewApplicationResolver()

	return &graphql.Field{
		Type:        types.ApplicationType,
		Description: "Get application information",
		Resolve:     applicationResolver.GetApplicationInformationResolve,
	}
}

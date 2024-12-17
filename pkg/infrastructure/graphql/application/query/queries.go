package query

import (
	"github.com/graphql-go/graphql"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/graphql/application/resolve"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/graphql/application/types"
)

func GetApplicationInformationQuery(appResolver *resolve.ApplicationResolver) *graphql.Field {
	return &graphql.Field{
		Type:        types.ApplicationType,
		Description: "Get application information",
		Resolve:     appResolver.GetApplicationInformationResolve,
	}
}

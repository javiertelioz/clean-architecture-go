package resolve

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/serializers"
)

type ApplicationResolver struct {
	appName string
}

// NewApplicationResolver godoc
func NewApplicationResolver(appName string) *ApplicationResolver {
	return &ApplicationResolver{
		appName: appName,
	}
}

func (r *ApplicationResolver) GetApplicationInformationResolve(_ graphql.ResolveParams) (interface{}, error) {
	message := fmt.Sprintf("Welcome to %s", r.appName)
	payload := serializers.NewApplicationSerializer(message)

	return payload, nil
}

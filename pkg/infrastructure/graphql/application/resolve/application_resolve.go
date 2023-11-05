package resolve

import (
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/javiertelioz/clean-architecture-go/config"
	"github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/serializers"
)

type ApplicationResolver struct {
	appName string
}

func NewApplicationResolver() *ApplicationResolver {
	appName, _ := config.GetConfig[string]("AppName")

	return &ApplicationResolver{
		appName: appName,
	}
}

func (r *ApplicationResolver) GetApplicationInformationResolve(_ graphql.ResolveParams) (interface{}, error) {
	message := fmt.Sprintf("Welcome to %s", r.appName)
	payload := serializers.NewApplicationSerializer(message)

	return payload, nil
}

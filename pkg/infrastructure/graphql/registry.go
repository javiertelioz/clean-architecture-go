package graphql

import (
	"fmt"

	"github.com/javiertelioz/clean-architecture-go/pkg/domain/contracts/services"
	applicationResolve "github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/graphql/application/resolve"
	userResolve "github.com/javiertelioz/clean-architecture-go/pkg/infrastructure/graphql/user/resolve"
)

type ResolverRegistry struct {
	resolvers map[string]interface{}
}

// NewResolverRegistry godoc
func NewResolverRegistry(
	cryptoService services.CryptoService,
	userService services.UserService,
	loggerService services.LoggerService,
	appName string,
) *ResolverRegistry {
	resolvers := map[string]interface{}{
		"user":        userResolve.NewUserResolver(cryptoService, userService, loggerService),
		"application": applicationResolve.NewApplicationResolver(appName),
	}

	return &ResolverRegistry{resolvers: resolvers}
}

// GetResolver godoc
func (r *ResolverRegistry) GetResolver(name string) (interface{}, error) {
	resolver, ok := r.resolvers[name]
	if !ok {
		return nil, fmt.Errorf("resolver '%s' not found", name)
	}

	return resolver, nil
}

func (r *ResolverRegistry) GetApplicationResolver() *applicationResolve.ApplicationResolver {
	return r.resolvers["application"].(*applicationResolve.ApplicationResolver)
}

func (r *ResolverRegistry) GetUserResolver() *userResolve.UserResolver {
	return r.resolvers["user"].(*userResolve.UserResolver)
}

package resolvers

import (
	"github.com/graphql-go/graphql"
	"github.com/javiertelioz/clean-architecture-go/pkg/domain/entity"
)

func GetUser(p graphql.ResolveParams) (interface{}, error) {
	user := entity.User{
		Name:     "John",
		LastName: "Doe",
		Email:    "john.doe@example.com",
	}
	return user, nil
}

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type GraphQLController struct {
	handler *handler.Handler
}

func NewGraphQLController(schema graphql.Schema) *GraphQLController {
	graphQLHandler := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     true,
		GraphiQL:   true,
		Playground: true,
	})

	return &GraphQLController{
		handler: graphQLHandler,
	}
}

func (c *GraphQLController) GraphQLHandler(ctx *gin.Context) {
	c.handler.ServeHTTP(ctx.Writer, ctx.Request)
}

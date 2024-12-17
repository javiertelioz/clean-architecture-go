package controllers

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var sandboxHTML = []byte(`
<!DOCTYPE html>
<html lang="en">
<body style="margin: 0; overflow-x: hidden; overflow-y: hidden">
<div id="sandbox" style="height:100vh; width:100vw;"></div>
<script src="https://embeddable-sandbox.cdn.apollographql.com/_latest/embeddable-sandbox.umd.production.min.js"></script>
<script>
new window.EmbeddedSandbox({
  target: "#sandbox",
  // Pass through your server href if you are embedding on an endpoint.
  // Otherwise, you can pass whatever endpoint you want Sandbox to start up with here.
  initialEndpoint: "http://localhost:8080/graphql",
});
// advanced options: https://www.apollographql.com/docs/studio/explorer/sandbox#embedding-sandbox
</script>
</body></html>`)

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

func (c *GraphQLController) SandboxHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(sandboxHTML)
	if err != nil {
		http.Error(w, "Failed to render sandbox HTML", http.StatusInternalServerError)
	}
}

func (c *GraphQLController) GraphQLHandler(w http.ResponseWriter, r *http.Request) {
	c.handler.ServeHTTP(w, r)
}

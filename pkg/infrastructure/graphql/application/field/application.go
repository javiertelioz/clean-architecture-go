package field

import "github.com/graphql-go/graphql"

var ApplicationField = graphql.Fields{
	"message": &graphql.Field{
		Type:        graphql.String,
		Description: "Application Name",
	},
	"version": &graphql.Field{
		Type:        graphql.String,
		Description: "Application Version",
	},
	"date": &graphql.Field{
		Type:        graphql.String,
		Description: "Application Date",
	},
}

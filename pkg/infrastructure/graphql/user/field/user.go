package field

import "github.com/graphql-go/graphql"

var UserField = graphql.Fields{
	"id": &graphql.Field{
		Type: graphql.Int,
	},
	"name": &graphql.Field{
		Type: graphql.String,
	},
	"lastname": &graphql.Field{
		Type: graphql.String,
	},
	"surname": &graphql.Field{
		Type: graphql.String,
	},
	"email": &graphql.Field{
		Type: graphql.String,
	},
	"phone": &graphql.Field{
		Type: graphql.String,
	},
}

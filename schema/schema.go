package schema

import (
	"github.com/graphql-go/graphql"
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": UserQuery["user"],
			"users": UserQuery["users"],
		},
	},
)

var MutationType = graphql.NewObject(
    graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createUser": UserMutation["createUser"],
			"updateUser": UserMutation["updateUser"],
            "deleteUser": UserMutation["deleteUser"],
		},
	},
)
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    QueryType,
	Mutation: MutationType,
})

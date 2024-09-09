package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"graph-go/schema"

	"github.com/graphql-go/graphql"
)

func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", result.Errors)
	}
	return result
}

func GraphqlHandler(w http.ResponseWriter, r *http.Request) {
	var p map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	result := ExecuteQuery(p["query"].(string), schema.Schema)
	json.NewEncoder(w).Encode(result)
}
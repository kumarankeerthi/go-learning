package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/kumarankeerthi/go-learning/graphql-go/gql"
)

// type Query struct {
// 	query string
// }

func main() {
	//var query Query
	fmt.Println("Graphql example using graphlql-go library")
	schema, _ := graphql.NewSchema(graphql.SchemaConfig{
		Query: gql.RootQuery,
	})
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		//q := json.NewDecoder(r.Body).Decode(&query)
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: r.URL.Query().Get("query"),
		})
		json.NewEncoder(w).Encode(result)
	})

	http.ListenAndServe(":1234", nil)
}

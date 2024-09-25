package main

import (
	"graph-go/database"
	"graph-go/schema"
	"log"
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/logrusorgru/aurora"
)

func main() {
	database.InitDB()
	defer database.DB.Close()

	h := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: true, // Mengaktifkan UI GraphiQL
	})

	http.Handle("/graphql", h)
	log.Println(aurora.Green("http://localhost:9122/graphql"))
	log.Fatal(http.ListenAndServe(":9122", nil))
}

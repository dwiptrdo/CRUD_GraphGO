package main

import (
	"graph-go/database"
	"graph-go/schema"
	"log"
	"net/http"

	"github.com/graphql-go/handler"
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
	log.Println("Server berjalan di port 9122")
	log.Fatal(http.ListenAndServe(":9122", nil))
}

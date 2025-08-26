package main

import (
	"fmt"
	"golang-websocket/internal/db"
	"golang-websocket/internal/db/migration"
	"golang-websocket/internal/routes"
	"log"
	"net/http"
)

func main() {
	db.ConnectDB()
	migration.Migrate()

	r := routes.RegisterRoutes()

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	fmt.Println("Server running on port 8080")

	log.Fatal(http.ListenAndServe(":8080", r))
	log.Println("Server started on :8080")
}

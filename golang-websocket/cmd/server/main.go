package main

import (
	"golang-websocket/internal/db"
	"golang-websocket/internal/routes"
	"log"
	"net/http"
)

func main(){
	db.ConnectDB()

 r:=routes.RegisterRoutes()


	http.ListenAndServe(":8080", r)

	log.Fatal(http.ListenAndServe(":8080", r))

}

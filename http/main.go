package main

import (
	"http-server/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", controllers.HomeController)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
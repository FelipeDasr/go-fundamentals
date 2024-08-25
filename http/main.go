package main

import (
	"http-server/controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/home", controllers.HomeController)
	http.HandleFunc("/home", controllers.PageRenderController)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
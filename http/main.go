package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/home", func(res http.ResponseWriter, req *http.Request) {
		type HomeResponse struct {
			Message string `json:"message"`
			Code int `json:"code"`
		}

		responseJSON, err := json.Marshal(HomeResponse{"Hello my friend", 200})
		if err != nil {
			log.Fatal(err)
			res.Write([]byte("Error"))
		}

		res.Header().Set("Content-Type", "application/json")
		res.Write(responseJSON)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
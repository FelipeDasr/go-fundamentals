package controllers

import (
	"http-server/templates"
	"net/http"
)

type UserData struct {
	Name string
	Email string
}

func PageRenderController(res http.ResponseWriter, req *http.Request) {
	user := UserData{"Felipe Dasr", "felipe@email.com"}
	templates.Templates.ExecuteTemplate(res, "home.html", user)
}
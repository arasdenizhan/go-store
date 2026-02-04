package handler

import (
	"net/http"
	"text/template"
)

func LoginGetHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template/login.html")
	if err != nil {
		ErrorHandler(w, r)
		return
	}

	err = tmpl.Execute(w, nil)
}
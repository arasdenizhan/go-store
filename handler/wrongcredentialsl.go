package handler

import (
	"net/http"
	"text/template"
)

func WrongCredentialsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)

	tmpl, _ := template.ParseFiles("template/401.html")
	tmpl.Execute(w, nil)
}
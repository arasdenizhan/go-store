package handler

import (
	"net/http"
	"text/template"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)

	tmpl, _ := template.ParseFiles("template/500.html")
	tmpl.Execute(w, nil)
}
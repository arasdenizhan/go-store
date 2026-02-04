package handler

import (
	"net/http"
	"text/template"

	"github.com/arasdenizhan/go-store/auth"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	err := auth.CheckCookie(r)
	if err != nil {
		LoginGetHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusNotFound)

	tmpl, _ := template.ParseFiles("template/404.html")
	tmpl.Execute(w, nil)
}
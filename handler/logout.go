package handler

import (
	"net/http"
	"text/template"

	"github.com/arasdenizhan/go-store/constants"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     constants.COOKIE_NAME,
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
	}
	http.SetCookie(w, &cookie)

	w.WriteHeader(http.StatusOK)

	tmpl, err := template.ParseFiles("template/logout-success.html")
	if err != nil {
		ErrorHandler(w, r)
		return
	}
	tmpl.Execute(w, nil)
}
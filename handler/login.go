package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"text/template"

	"github.com/arasdenizhan/go-store/constants"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		LoginGetHandler(w, r)
		return

	case http.MethodPost:
		loginRequest := LoginRequest{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}

		data, err := json.Marshal(loginRequest)
		if err != nil {
			ErrorHandler(w, r)
			return
		}

		resp, err := http.Post(constants.API_URL+"/auth/login", "application/json", bytes.NewReader(data))
		if err != nil {
			ErrorHandler(w, r)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode == 401 {
			WrongCredentialsHandler(w, r)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			ErrorHandler(w, r)
			return
		}

		login := LoginResponse{}
		err = json.Unmarshal(body, &login)
		if err != nil {
			ErrorHandler(w, r)
			return
		}

		cookie := http.Cookie{
			Name:     constants.COOKIE_NAME,
			Value:    login.Token,
			HttpOnly: true,
			Secure:   true,
			Path:     "/",
		}
		http.SetCookie(w, &cookie)
		tmpl, err := template.ParseFiles("template/login-success.html")
		if err != nil {
			ErrorHandler(w, r)
			return
		}

		tmpl.Execute(w, nil)
	}
}

type LoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
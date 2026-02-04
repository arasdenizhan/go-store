package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"text/template"

	"github.com/arasdenizhan/go-store/auth"
	"github.com/arasdenizhan/go-store/constants"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	err := auth.CheckCookie(r)
	if err != nil {
		LoginGetHandler(w, r)
		return
	}

	resp, err := http.Get(constants.API_URL + "/products")
	if err != nil {
		ErrorHandler(w, r)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ErrorHandler(w, r)
		return
	}

	productList := []Product{}
	err = json.Unmarshal(body, &productList)
	if err != nil {
		ErrorHandler(w, r)
		return
	}

	tmpl, err := template.ParseFiles("template/index.html")
	if err != nil {
		ErrorHandler(w, r)
		return
	}

	err = tmpl.Execute(w, productList)
	if err != nil {
		ErrorHandler(w, r)
		return
	}
}

type Product struct {
	ID          int `json:"id"`
	Title       string `json:"title"`
	Price       float64 `json:"price"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Image       string `json:"image"`
}
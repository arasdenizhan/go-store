package main

import (
	"log"
	"net/http"

	"github.com/arasdenizhan/go-store/handler"
)

func main() {
	fs := http.FileServer(http.Dir("template/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handler.NotFoundHandler)
	http.HandleFunc("/{$}", handler.ProductsHandler)
	http.HandleFunc("/cart", handler.CartHandler)
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/logout", handler.LogoutHandler)
	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
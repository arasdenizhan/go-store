package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"text/template"

	"github.com/arasdenizhan/go-store/auth"
	"github.com/arasdenizhan/go-store/constants"
)

func CartHandler(w http.ResponseWriter, r *http.Request) {
	err := auth.CheckCookie(r)
	if err != nil {
		LoginGetHandler(w, r)
		return
	}

	rawUserId, err := auth.GetUserId(r)
	if err != nil {
		ErrorHandler(w, r)
		return
	}

	userId, err := strconv.Atoi(rawUserId)
	if err != nil {
		ErrorHandler(w, r)
		return
	}

	resp, err := http.Get(constants.API_URL + "/carts")
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

	cartResponse := []CartResponse{}
	err = json.Unmarshal(body, &cartResponse)
	if err != nil {
		ErrorHandler(w, r)
		return
	}

	cartProductList := []CartProduct{}
	cartForUser := CartResponse{}
	for _, cart := range cartResponse {
		if cart.UserID == userId {
			cartForUser = cart
		}
	}

	totalPrice := float64(0)
	totalQuantity := 0
	if cartForUser.ID != 0 {
		for _, prdct := range cartForUser.Products {
			resp, err := http.Get(constants.API_URL + "/products/" + strconv.Itoa(prdct.ProductId))
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

			cartProduct := CartProduct{}
			err = json.Unmarshal(body, &cartProduct)
			if err != nil {
				ErrorHandler(w, r)
				return
			}
			cartProduct.Quantity = prdct.Quantity
			totalQuantity = totalQuantity + int(prdct.Quantity)
			cartProduct.ItemTotal = float64(prdct.Quantity) * cartProduct.Price
			totalPrice = totalPrice + cartProduct.ItemTotal
			cartProductList = append(cartProductList, cartProduct)
		}
	}

	tmpl, err := template.ParseFiles("template/cart.html")
	if err != nil {
		ErrorHandler(w, r)
		return
	}
	tmpl.Execute(w, CartDto{
		Products:   cartProductList,
		Total:      float64(totalPrice),
		TotalItems: totalQuantity,
	})
}

type CartResponse struct {
	ID       int          `json:"id"`
	UserID   int          `json:"userId"`
	Products []ProductDto `json:"products"`
}

type CartProduct struct {
	Title     string  `json:"title"`
	Price     float64 `json:"price"`
	Category  string  `json:"category"`
	Image     string  `json:"image"`
	Quantity  int
	ItemTotal float64
}

type CartDto struct {
	Products   []CartProduct
	Total      float64
	TotalItems int
}

type ProductDto struct {
	ProductId int `json:"productId"`
	Quantity  int `json:"quantity"`
}
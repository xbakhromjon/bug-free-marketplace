package controller

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"golang-project-template/internal/basket/app"
	basket "golang-project-template/internal/basket/domain"
	"net/http"
	"strconv"
)

type CartController struct {
	cartUseCase app.CartService
}

func NewCartController(cartUseCase app.CartService) *CartController {
	return &CartController{cartUseCase: cartUseCase}
}

func (cc *CartController) CreateCart(w http.ResponseWriter, r *http.Request) {
	userIDParam := chi.URLParam(r, "user_id")
	userID, err := strconv.Atoi(userIDParam)
	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}

	_, err = cc.cartUseCase.CreateCart(userID)
	if err != nil {
		http.Error(w, "Failed to create a new basket", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Basket successfully created")
}

func (cc *CartController) CreateBasketItem(w http.ResponseWriter, r *http.Request) {
	var req basket.CartItems
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := cc.cartUseCase.CreateCartItem(req)
	if err != nil {
		http.Error(w, "Failed to create cart item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Cart Item created successfully")
}

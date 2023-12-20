package controller

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"golang-project-template/internal/basket/app"
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

	_, err = cc.cartUseCase.CreateBasket(userID)
	if err != nil {
		http.Error(w, "Failed to create a new basket", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Basket successfully created")
}

package controller

import (
	"encoding/json"
	"fmt"
	"golang-project-template/internal/basket/app"
	basket "golang-project-template/internal/basket/domain"
	"log"
	"net/http"
)

type CartController struct {
	cartUseCase app.CartService
}

func NewCartController(cartUseCase app.CartService) *CartController {
	return &CartController{cartUseCase: cartUseCase}
}

func (cc *CartController) CreateCart(w http.ResponseWriter, r *http.Request) {
	var basketReq basket.Cart
	err := json.NewDecoder(r.Body).Decode(&basketReq)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
		return
	}
	userID := basketReq.UserId
	log.Println(userID)
	if userID <= 0 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	id, err := cc.cartUseCase.CreateBasket(userID)
	if err != nil {
		http.Error(w, "Failed to create a new basket", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Basket successfully created %v", id)
}

func (cc *CartController) AddItem(w http.ResponseWriter, r *http.Request) {
	var cItems basket.CartItems
	err := json.NewDecoder(r.Body).Decode(&cItems)
	if err != nil {
		http.Error(w, "Invalid JSON request", http.StatusBadRequest)
	}

	id, err := cc.cartUseCase.AddItem(&cItems)
	if err != nil {
		http.Error(w, "Failed add item to cart", http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Item successfully added to cart %v", id)
}

func (cc *CartController) GetAll(w http.ResponseWriter, r *http.Request) {
	CartReq := basket.Cart{}
	err := json.NewDecoder(r.Body).Decode(&CartReq)
	if err != nil {
		http.Error(w, "Invalid cart ID", http.StatusBadRequest)
		return
	}
	cItems, err := cc.cartUseCase.GetAll(CartReq.Id)
	if err != nil {
		if err == basket.ErrCartItemNotFound {
			http.Error(w, "Cart items not found", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to retrieve cart items", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cItems)
}

func (cc *CartController) Update(w http.ResponseWriter, r *http.Request) {
	var cItems basket.CartItems
	err := json.NewDecoder(r.Body).Decode(&cItems)
	if err != nil {
		http.Error(w, "Invalid cart ID", http.StatusBadRequest)
		return
	}
	err = cc.cartUseCase.UpdateProductQuantity(cItems.CartId, cItems.Quantity)
	if err != nil {
		if err == basket.ErrCartItemNotFound {
			http.Error(w, "Couldn't update", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to retrieve cart items", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cItems)
}

func (cc *CartController) Delete(w http.ResponseWriter, r *http.Request) {
	var cItems basket.CartItems
	err := json.NewDecoder(r.Body).Decode(&cItems)
	if err != nil {
		http.Error(w, "Invalid cart ID", http.StatusBadRequest)
		return
	}
	id, err := cc.cartUseCase.DeleteProductFromCart(cItems.CartId, cItems.ProductId)
	if err != nil {
		if err == basket.ErrCartItemNotFound {
			http.Error(w, "Couldn't delete", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to retrieve cart items", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(id)
}

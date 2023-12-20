package handler

import (
	"encoding/json"
	"golang-project-template/internal/order/app"
	"log"
	"net/http"
)

type CartHandler struct {
	Cart app.CartItemService
}
func (c*CartHandler) createBasket(w *http.ResponseWriter, r *http.Request){
	cart := cart{}
}
func (c *CartHandler) AddItem(w *http.ResponseWriter, r *http.Request) {

	cartItems := cartItems{}
	err := json.NewDecoder(r.Body).Decode( cartItems)
	if err!=nil{
		panic(err)
	}
	newItemID ,err :=c.Cart.Add(cartItems{})
	newShopId, err := h.ShopService.Create(newShop)
	if err != nil {
		log.Printf("Error creating a new shop:  %v", err)
		http.Error(w, "Failed to create shop", http.StatusInternalServerError)
		return
	}

}

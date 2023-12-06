package handler

import (
	"encoding/json"
	"golang-project-template/internal/shop/app"
	"golang-project-template/internal/shop/domain"
	"log"

	"net/http"
)

type ShopHandler struct {
	ShopService app.ShopService
}

func (h *ShopHandler) CreateShop(w http.ResponseWriter, r *http.Request) {
	newUser := domain.NewShop{}
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		log.Printf("Error decoding new user %v", err)
		http.Error(w, "Bad request", http.StatusInternalServerError)
		return
	}

	newShopId, err := h.ShopService.Create(domain.NewShop{
		Name:    newUser.Name,
		OwnerId: newUser.OwnerId,
	})
	if err != nil {
		log.Printf("Error creating a new shop:  %v", err)
		http.Error(w, "Failed to create shop", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newShopId)
}

package handler

import (
	"encoding/json"
	"fmt"
	"golang-project-template/internal/shop/app"
	"golang-project-template/internal/shop/domain"
	"log"

	"net/http"
)

type ShopHandler struct {
	ShopService app.ShopService
}

func (h *ShopHandler) CreateShop(w http.ResponseWriter, r *http.Request) {
	newShop := domain.NewShop{}
	err := json.NewDecoder(r.Body).Decode(&newShop)
	if err != nil {
		log.Printf("Error decoding new user %v", err)
		http.Error(w, "Bad request", http.StatusInternalServerError)
		return
	}

	newShopId, err := h.ShopService.Create(domain.NewShop{
		Name:    newShop.Name,
		OwnerId: newShop.OwnerId,
	})
	if err != nil {
		log.Printf("Error creating a new shop:  %v", err)
		http.Error(w, "Failed to create shop", http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(struct {
		Id int
	}{newShopId})

	fmt.Fprint(w, res)

}

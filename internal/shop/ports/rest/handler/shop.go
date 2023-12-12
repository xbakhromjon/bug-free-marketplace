package handler

import (
	"encoding/json"
	"golang-project-template/internal/shop/app"
	"golang-project-template/internal/shop/domain"
	"log"
	"strconv"

	"net/http"

	"github.com/go-chi/chi/v5"
)

type ShopHandler struct {
	ShopService app.ShopService
}

func (h *ShopHandler) CreateShop(w http.ResponseWriter, r *http.Request) {
	newShop := domain.NewShop{}
	err := json.NewDecoder(r.Body).Decode(&newShop)
	if err != nil {
		log.Printf("Error decoding new user %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	newShopId, err := h.ShopService.Create(newShop)
	if err != nil {
		log.Printf("Error creating a new shop:  %v", err)
		http.Error(w, "Failed to create shop", http.StatusInternalServerError)
		return
	}

	resdata := map[string]int{"Id": newShopId}
	jsonData, err := json.Marshal(resdata)

	if err != nil {
		log.Printf("Error marshalling a new shop id:  %v", err)
		http.Error(w, "Failed to marshall shop id", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (h *ShopHandler) GetShopById(w http.ResponseWriter, r *http.Request) {
	shopId, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Printf("Error converting shop id to number: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	shop, err := h.ShopService.GetShopById(shopId)
	if err != nil {
		log.Printf("Error getting shop by given id: %v", err)
		http.Error(w, "Failed to get the shop by given id", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(shop)

	if err != nil {
		log.Printf("Error marshalling the shop:  %v", err)
		http.Error(w, "Failed to marshall the shop", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

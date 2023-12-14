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

type ShopResponse struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	OwnerId   int    `json:"owner_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

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
	shopResponse := ShopResponse{}
	shopResponse.Id = shop.GetId()
	shopResponse.Name = shop.GetName()
	shopResponse.OwnerId = shop.GetOwnerId()
	shopResponse.CreatedAt = shop.GetCreatedAt()
	shopResponse.UpdatedAt = shop.GetUpdatedAt()

	jsonData, err := json.Marshal(shopResponse)

	if err != nil {
		log.Printf("Error marshalling the shop:  %v", err)
		http.Error(w, "Failed to marshall the shop", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func (h *ShopHandler) GetAllShops(w http.ResponseWriter, r *http.Request) {

	limit, err := ParseLimitQuery(r.URL.Query().Get("limit"))
	if err != nil {
		log.Printf("Error parsing limit: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	offset, err := ParseOffsetQuery(r.URL.Query().Get("offset"))
	if err != nil {
		log.Printf("Error parsing offset: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	search := r.URL.Query().Get("search")

	shops, err := h.ShopService.GetAllShops(limit, offset, search)
	if err != nil {
		log.Printf("Error while getting all shops: %v", err)
		http.Error(w, "Failed to get all shops", http.StatusInternalServerError)
		return
	}

	res := []ShopResponse{}
	for _, shop := range shops {
		shopResponse := ShopResponse{}

		shopResponse.Id = shop.GetId()
		shopResponse.Name = shop.GetName()
		shopResponse.OwnerId = shop.GetOwnerId()
		shopResponse.CreatedAt = shop.GetCreatedAt()
		shopResponse.UpdatedAt = shop.GetUpdatedAt()

		res = append(res, shopResponse)
	}
	jsonData, err := json.Marshal(res)
	if err != nil {
		log.Printf("Error marshalling the shops:  %v", err)
		http.Error(w, "Failed to marshall all shops", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func ParseLimitQuery(limit string) (int, error) {
	if limit == "" {
		return 10, nil
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return 0, err
	}

	return limitInt, nil

}

func ParseOffsetQuery(offset string) (int, error) {
	if offset == "" {
		return 1, nil
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return 0, err
	}

	return offsetInt, nil

}

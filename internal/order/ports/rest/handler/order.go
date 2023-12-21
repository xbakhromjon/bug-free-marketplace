package handler

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"golang-project-template/internal/order/app"
	"net/http"
	"strconv"
)

type OrderHandler struct {
	OrderService app.OrderService
}

func (o *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	basketID, err := getIdFromRequest(r)
	if err != nil {
		logError("createOrder", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = o.OrderService.CreateOrder(basketID)
	if err != nil {
		logError("createOrder", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

//func (o *OrderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
//
//}

//func (o *OrderHandler)

func getIdFromRequest(r *http.Request) (int, error) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, err
	}

	if id == 0 {
		return 0, errors.New("id can't be 0")
	}

	return int(id), nil
}

//// Vars returns the route variables for the current request, if any.
//func Vars(r *http.Request) map[string]string {
//	if rv := r.Context().Value(varsKey); rv != nil {
//		return rv.(map[string]string)
//	}
//	return nil
//}

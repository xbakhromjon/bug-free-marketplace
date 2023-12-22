package handler

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"golang-project-template/internal/order/app"
	"golang-project-template/internal/order/domain"
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

	w.WriteHeader(http.StatusCreated)
}

func (o *OrderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	orderId, err := getIdFromRequest(r)
	if err != nil {
		logError("getOrderByID", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	order, err := o.OrderService.GetOrderByID(orderId)
	if err != nil {
		if errors.Is(err, domain.ErrOrderNotFound) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		logError("getOrderById", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(order)
	if err != nil {
		logError("getOrderByID", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Context-Type", "application/json")

	_, err = w.Write(response)
	if err != nil {
		logError("getOrderByID", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (o *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := o.OrderService.GetAllOrders()
	if err != nil {
		logError("getAllOrders", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(orders)
	if err != nil {
		logError("getAllOrders", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")

	_, err = w.Write(response)
	if err != nil {
		logError("getOrderByID", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// this method changes order's status by id
func (o *OrderHandler) UpdateStatusOrder(w http.ResponseWriter, r *http.Request) {
	orderID, err := getIdFromRequest(r)
	if err != nil {
		logError("updateStatusOrder", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := o.OrderService.UpdateStatus(orderID); err != nil {
		logError("updateStatusOrder", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

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

package handler

import (
	"golang-project-template/internal/order/app"
)

type OrderHandler struct {
	OrderService app.OrderService
}

//func (o *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
//	reqBytes, err := io.ReadAll(r.Body)
//	if err != nil {
//		logError()
//	}
//}

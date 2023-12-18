package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Router(controller *CartController) http.Handler {
	router := gin.Default()

	router.POST("/carts/:user_id/create", controller.CreateBasket)
	router.GET("carts/:user_id", controller.GetBasket)
	router.POST("/carts/:user_id/add-product/:product_id/:quantity", controller.AddProductToCart)
	router.PUT("/carts/:user_id/increment/:product_id", controller.IncrementProductQuantity)
	router.PUT("/carts/:user_id/decrement/:product_id", controller.DecrementProductQuantity)
	log.Println("Server started on: 8080")
	return router
}

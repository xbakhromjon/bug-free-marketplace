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
	router.PUT("/carts/:user_id/add-product/:product_id/:quantity", controller.AddProductToCart)

	log.Println("Server starter on: 8080")
	return router
}

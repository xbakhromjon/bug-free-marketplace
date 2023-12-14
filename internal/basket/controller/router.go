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

	log.Println("Server starter on: 8080")
	return router
}

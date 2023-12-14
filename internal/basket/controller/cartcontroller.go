package controller

import (
	"github.com/gin-gonic/gin"
	"golang-project-template/internal/basket/app"
	"net/http"
)

type CartController struct {
	cartUseCase *app.CartServiceImpl
}

func (cc *CartController) CreateBasket(c *gin.Context) {
	userID, err := getUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid User id"})
		return
	}

	_, err = cc.cartUseCase.CreateCart(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a new basket"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Basket successfully created"})
}

func (cc *CartController) GetBasket(c *gin.Context) {
	userID, err := getUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user Id"})
		return
	}

	cart, err := cc.cartUseCase.GetBasket(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can't retrieve the basket"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cart": cart})
}

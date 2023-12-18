package controller

import (
	"github.com/gin-gonic/gin"
	"golang-project-template/internal/basket/app"
	"net/http"
	"strconv"
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

func (cc *CartController) AddProductToCart(c *gin.Context) {
	userId, err := getUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	productId, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product id"})
		return
	}

	quantity, err := strconv.Atoi(c.Param("quantity"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity"})
		return
	}

	updateBasket, err := cc.cartUseCase.AddProductToCart(userId, productId, quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product to the cart"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"cart": updateBasket})
}

func (cc *CartController) IncrementProductQuantity(c *gin.Context) {
	userId, err := getUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	productId, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product id"})
	}

	err = cc.cartUseCase.IncrementProductQuantity(userId, productId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to increment quantity"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Quantity is incremented"})
}

func (cc *CartController) DecrementProductQuantity(c *gin.Context) {
	userId, err := getUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
	}

	productId, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product id"})
	}

	err = cc.cartUseCase.DecrementProductQuantity(userId, productId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decrement"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Quantity is decremented"})
}

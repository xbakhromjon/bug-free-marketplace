package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func getUserIdFromContext(c *gin.Context) (int, error) {
	userIDParam := c.Param("user_id")
	userId, err := strconv.Atoi(userIDParam)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

package handler
package handler

import (
	"fmt"
	"net/http"

	"github.com/asadbek21coder/fintracker2/internal/domain"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase domain.UserUseCase
}

// Imlement UserHandler mehtods here

package app

package app

import (
	"github.com/asadbek21coder/fintracker2/internal/domain"
)

type UserUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) domain.UserUseCase {
	return UserUseCase{userRepo: repo}
}

// Implement userUseCase methods here

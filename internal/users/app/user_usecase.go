package app

import (
	"errors"
	"golang-project-template/internal/users/domain"

	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepository domain.UserRepository
	f              domain.UserFactory
}

func NewUserUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) RegisterUser(user *domain.User) (*domain.User, error) {

	user, err := u.userRepository.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) LoginUser(phoneNumber, pass string) (bool, error) {
	user, err := u.userRepository.GetByPhoneNumber(phoneNumber)
	if err != nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	if err != nil {
		return false, errors.New("Invalid phone number or password")
	}

	return true, nil
}

func (u *userUsecase) GetUserDataByID(userID int) (*domain.User, error) {
	user, err := u.userRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}
	retrievedUser := u.f.MapUserData(user.ID, user.Name, user.PhoneNumber, user.Role, user.CreatedAt)

	return retrievedUser, nil
}

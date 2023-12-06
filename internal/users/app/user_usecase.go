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

type UserUsecase interface {
	RegisterMerchantUser(user *domain.NewUser) (int, error)
	RegisterCustomer(user *domain.NewUser) (int, error)
	LoginUser(phoneNumber, pass string) (bool, error)
	GetUserDataPhoneNumber(phoneNumber string) (*domain.User, error)
}

func NewUserUsecase(userRepository domain.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) RegisterMerchantUser(newUser *domain.NewUser) (int, error) {
	userFromFactory := u.f.CreateMerchantUser(newUser)
	id, err := u.userRepository.Save(userFromFactory)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *userUsecase) RegisterCustomer(newUser *domain.NewUser) (int, error) {
	userFromFactory := u.f.CreateCustomerUser(newUser)
	id, err := u.userRepository.Save(userFromFactory)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *userUsecase) LoginUser(phoneNumber, pass string) (bool, error) {
	user, err := u.userRepository.FindOneByPhoneNumber(phoneNumber)
	if err != nil {
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(pass))
	if err != nil {
		return false, errors.New("Invalid phone number or password")
	}

	return true, nil
}

func (u *userUsecase) GetUserDataPhoneNumber(phoneNumber string) (*domain.User, error) {
	user, err := u.userRepository.FindOneByPhoneNumber(phoneNumber)
	if err != nil {
		return nil, err
	}

	return user, nil
}

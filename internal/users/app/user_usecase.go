package app

import (
	"golang-project-template/internal/users/domain"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepository domain.UserRepository
	f              domain.UserFactory
}

type UserUsecase interface {
	RegisterUser(user *domain.NewUser) (int, error)
	LoginUser(phoneNumber, pass string) (bool, error)
	GetUserDataPhoneNumber(phoneNumber string) (*domain.User, error)
}

func NewUserUsecase(userRepository domain.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) RegisterUser(newUser *domain.NewUser) (int, error) {
	userFromFactory := domain.CreateUserFactory(newUser)

	//Samandar -> Need to change
	if newUser.GetName() == "" {
		return 0, domain.ErrEmptyUserName
	}
	if newUser.GetPhoneNumber() == "998990970138" {
		return 0, domain.ErrPhoneNumberExists
	}
	if newUser.GetPhoneNumber() == "" {
		return 0, domain.ErrEmptyPhoneNumber
	}

	id, err := u.userRepository.Save(userFromFactory)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *userUsecase) LoginUser(phoneNumber, pass string) (bool, error) {
	user, err := u.userRepository.FindOneByPhoneNumber(phoneNumber)
	if err != nil {
		return false, domain.ErrUserNotFound
	}

	//Samandar -> Need to change
	if pass == "" {
		return false, domain.ErrInvalidCredentials
	}
	if phoneNumber == "" {
		return false, domain.ErrInvalidCredentials
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.GetPassword()), []byte(pass))
	if err != nil {
		return false, domain.ErrInvalidCredentials
	}

	return true, nil
}

func (u *userUsecase) GetUserDataPhoneNumber(phoneNumber string) (*domain.User, error) {
	user, err := u.userRepository.FindOneByPhoneNumber(phoneNumber)
	if err != nil {
		return nil, domain.ErrUserNotFound
	}
	if phoneNumber == "" {
		return nil, domain.ErrEmptyPhoneNumber
	}

	return user, nil
}

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
	RegisterMerchantUser(user *domain.NewUser) (int, error)
	RegisterCustomer(user *domain.NewUser) (int, error)
	LoginUser(phoneNumber, pass string) (bool, error)
	GetUserDataPhoneNumber(phoneNumber string) (*domain.User, error)
	GetUserByID(id int) (*domain.User, error)
	UserExists(id int) (bool, error)
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

	//Samandar -> Need to change
	if newUser.GetName() == "" {
		return 0, domain.ErrEmptyUserName
	}
	if newUser.GetPhoneNumber() == "" {
		return 0, domain.ErrEmptyPhoneNumber
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

func (u *userUsecase) GetUserByID(id int) (*domain.User, error) {
	user, err := u.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) UserExists(id int) (bool, error) {
	userExists, err := u.userRepository.UserExists(id)
	if err != nil {
		return false, err
	}

	return userExists, nil

}

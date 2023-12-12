package app

import (
	"golang-project-template/internal/users/domain"
	"strings"
)

type userUsecase struct {
	userRepository domain.UserRepository
	f              domain.UserFactory
}

type UserUsecase interface {
	RegisterMerchantUser(user *domain.NewUser) (int, error)
	RegisterCustomer(user *domain.NewUser) (int, error)
	RegisterAdmin(user *domain.NewUser) (int, error)
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
	err := validateUserInfoForRegister(userFromFactory.GetName(), userFromFactory.GetPhoneNumber(), userFromFactory.GetPassword())
	if err != nil {
		return 0, err
	}

	id, err := u.userRepository.Save(userFromFactory)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *userUsecase) RegisterCustomer(newUser *domain.NewUser) (int, error) {
	userFromFactory := u.f.CreateCustomerUser(newUser)

	err := validateUserInfoForRegister(
		userFromFactory.GetName(),
		userFromFactory.GetPhoneNumber(),
		userFromFactory.GetPassword())
	if err != nil {
		return 0, err
	}

	id, err := u.userRepository.Save(userFromFactory)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *userUsecase) RegisterAdmin(newUser *domain.NewUser) (int, error) {
	userFromFactory := u.f.CreateAdminUser(newUser)
	err := validateUserInfoForRegister(
		userFromFactory.GetName(),
		userFromFactory.GetPhoneNumber(),
		userFromFactory.GetPassword())
	if err != nil {
		return 0, err
	}

	id, err := u.userRepository.Save(userFromFactory)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (u *userUsecase) LoginUser(phoneNumber, pass string) (bool, error) {

	err := validateUserInfoForLogin(phoneNumber, pass)
	if err != nil {
		return false, err
	}

	ok, err := u.userRepository.UserExistByPhone(phoneNumber)
	if !ok || err != nil {
		return false, domain.ErrUserNotFound
	}

	return true, nil
}

func (u *userUsecase) GetUserDataPhoneNumber(phoneNumber string) (*domain.User, error) {

	if strings.TrimSpace(phoneNumber) == "" {
		return nil, domain.ErrEmptyPhoneNumber
	}

	user, err := u.userRepository.FindOneByPhoneNumber(phoneNumber)
	if err != nil {
		return nil, domain.ErrUserNotFound
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

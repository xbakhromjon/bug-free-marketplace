package app

import (
	"errors"
	"golang-project-template/internal/users/domain"
	"regexp"
	"strings"
)

func validatePassword(password string) error {
	if len(password) < 8 {
		return domain.ErrInvalidPassword
	}

	var (
		uppercase = regexp.MustCompile(`[A-Z]`)
		lowercase = regexp.MustCompile(`[a-z]`)
		digit     = regexp.MustCompile(`[0-9]`)
	)

	if !uppercase.MatchString(password) || !lowercase.MatchString(password) {
		return domain.ErrInvalidPassword
	}
	if !digit.MatchString(password) {
		return domain.ErrInvalidPassword
	}

	return nil
}

func validateUserInfoForRegister(name, phoneNumber, password string) error {

	if strings.TrimSpace(name) == "" {
		return domain.ErrEmptyUserName
	}
	if strings.TrimSpace(phoneNumber) == "" {
		return domain.ErrEmptyPhoneNumber
	}
	if validatePassword(password) == domain.ErrInvalidPassword {
		return domain.ErrInvalidPassword
	}

	return nil
}

func validateUserInfoForLogin(phoneNumber, password string) error {

	if strings.TrimSpace(phoneNumber) == "" {
		return domain.ErrInvalidCredentials
	}
	if validatePassword(password) == domain.ErrInvalidPassword {
		return domain.ErrInvalidCredentials
	}

	return nil
}

func validatePhoneNumberCount(phoneNumber string) error {
	if len(phoneNumber) > 12 {
		return errors.New("phone number count is more than 12")
	}
	return nil
}

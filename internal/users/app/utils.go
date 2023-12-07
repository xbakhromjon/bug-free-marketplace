package app

import (
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

func validateUserInfo(user *domain.User) error {

	if strings.TrimSpace(user.GetName()) == "" {
		return domain.ErrEmptyUserName
	}
	if strings.TrimSpace(user.GetPhoneNumber()) == "" {
		return domain.ErrEmptyPhoneNumber
	}
	if validatePassword(user.GetPassword()) == domain.ErrInvalidPassword {
		return domain.ErrInvalidPassword
	}

	return nil
}

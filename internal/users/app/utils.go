package app

import (
	"errors"
	"golang-project-template/internal/users/domain"
	"regexp"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
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

func CreateToken(user string, duration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(duration).Unix(),
		"iat": time.Now().Unix(),
		"sub": user,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte("marketplace"))
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyToken(token string) (jwt.MapClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte("marketplace"), nil
	}

	jwtToken, err := jwt.Parse(token, keyFunc)
	if err != nil {
		if ver, ok := err.(*jwt.ValidationError); ok {
			if ver.Errors & jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("expired token")
			}
		}
		return nil, errors.New("invalid token")
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

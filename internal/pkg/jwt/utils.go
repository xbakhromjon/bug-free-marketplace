package jwt

import (
	"errors"
	"golang-project-template/internal/pkg/config"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(sub string) (string, error) {
	claims := jwt.MapClaims{
		"exp": time.Now().Add(15 * time.Minute).Unix(),
		"iat": time.Now().Unix(),
		"sub": sub,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecretKey := config.NewConfig().JwtSecretKey
	token, err := jwtToken.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyToken(token string) (jwt.MapClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		jwtSecretKey := config.NewConfig().JwtSecretKey
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(jwtSecretKey), nil
	}

	jwtToken, err := jwt.Parse(token, keyFunc)
	if err != nil {
		if ver, ok := err.(*jwt.ValidationError); ok {
			if ver.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("expired token")
			}
		}
		return nil, errors.New(err.Error())
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New(err.Error())
	}

	return claims, nil
}

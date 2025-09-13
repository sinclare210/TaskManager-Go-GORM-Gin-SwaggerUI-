package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(Email string, Id uint) (string, error) {
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		return "", errors.New("SECRET_KEY not set in environment")
	}
	ParsedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Email": Email,
		"Id":    Id,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	})

	return ParsedToken.SignedString([]byte(secret))
}

func VerifyToken(token string) (string, uint, error) {
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		return "", 0, errors.New("SECRET_KEY not set in environment")
	}
	Parsedtoken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invaid signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return "", 0, err
	}

	if !Parsedtoken.Valid {
		return "", 0, errors.New("invalid token")
	}

	claims, ok := Parsedtoken.Claims.(jwt.MapClaims)
	if !ok {
		return "", 0, err
	}

	emailVal, ok := claims["Email"]
	if !ok {
		return "", 0, fmt.Errorf("email claim missing")
	}

	email, ok := emailVal.(string)
	if !ok {
		return "", 0, fmt.Errorf("email claim is not a string")
	}

	Id := uint(claims["Id"].(float64))
	if Id == 0 {
		return "", 0, err
	}

	return email, Id, nil

}

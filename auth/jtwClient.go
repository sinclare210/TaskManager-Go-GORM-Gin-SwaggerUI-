package auth

import (
	"errors"
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
		"Emali": Email,
		"Id":    Id,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	})

	return ParsedToken.SignedString([]byte(secret))
}

func VerifyToken(token string) (any, error) {
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		return nil, errors.New("SECRET_KEY not set in environment")
	}
	Parsedtoken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invaid signing method")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if !Parsedtoken.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := Parsedtoken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	return claims, nil

}

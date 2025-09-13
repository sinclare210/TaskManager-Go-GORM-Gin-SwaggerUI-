package auth

import "golang.org/x/crypto/bcrypt"

func GenerateHash(password string) (string, error) {
	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(HashedPassword), err
}

func VerifyHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

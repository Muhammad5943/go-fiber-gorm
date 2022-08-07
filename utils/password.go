package utils

import "golang.org/x/crypto/bcrypt"

func HashingPassword(password string) (string, error) {
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(hashedByte), nil
}

func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil // if err == nil ? true : false
}

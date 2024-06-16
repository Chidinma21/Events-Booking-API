package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), 14)
	return string(bytes), err
}

func ComparePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
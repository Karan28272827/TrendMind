package utils

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	var hashedPass, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal("Error hashing password", err)
	}

	return string(hashedPass)
}

func CompareHashAndPassword(dbPassword string, userPassword string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(userPassword)); err != nil {
		fmt.Println("Invalid password", err)
		return false
	}
	return true
}

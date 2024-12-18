package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

func HashPassword(password string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(res), err
}

func CheckHashedPassword(password, hashed string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return false
	}
	return true
}

func GenerateRandomName() string {
	prefix := "光迹"
	lettersAndNumbers := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	length := 6

	rand.Seed(time.Now().UnixNano())

	randomPart := make([]byte, length)
	for i := range randomPart {
		randomPart[i] = lettersAndNumbers[rand.Intn(len(lettersAndNumbers))]
	}

	return fmt.Sprintf("%s%s", prefix, string(randomPart))
}

package utils

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/mail"
	"strings"
)

func ValitEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ValitUuId(id string) (string, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {

	}
	return uuid.String(), nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func SplitTextToArray(s string) []string {
	listString := strings.Split(s, ",")
	return listString
}

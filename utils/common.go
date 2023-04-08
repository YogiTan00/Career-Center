package utils

import (
	"CareerCenter/utils/exceptions"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net/mail"
	"strings"
	"time"
)

func ValitEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ValitUuId(s string) (string, error) {
	id, err := uuid.Parse(s)
	if err != nil {

	}
	return id.String(), nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func RandomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
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

func JoinTextFromArray(s []string) string {
	text := strings.Join(s, ",")
	return text
}

func ToDate(s string) (time.Time, error) {
	conv, err := time.Parse("2006-01-02", s)
	if err != nil {
		return time.Time{}, exceptions.ErrCustomString("err format date")
	}
	return conv, nil
}

func ToOnlyDateResponse(date time.Time) string {
	toString := date.Format("2006-01-02")
	return toString
}

package utils

import (
	"CareerCenter/utils/exceptions"
	"encoding/json"
	"fmt"
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

func RandomStringNumber(n int) string {
	rand.Seed(time.Now().UnixNano())
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
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

func PrettyPrint(data interface{}) {
	jsonBytes, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(string(jsonBytes))
}

func Color(color string, message string) string {
	const (
		black   = "\033[30m"
		red     = "\033[31m"
		green   = "\033[32m"
		yellow  = "\033[33m"
		blue    = "\033[34m"
		magenta = "\033[35m"
		cyan    = "\033[36m"
		white   = "\033[37m"
	)

	var (
		result string
		reset  = "\033[0m"
	)
	switch strings.ToLower(color) {
	case "black":
		result = black
	case "red":
		result = red
	case "green":
		result = green
	case "yellow":
		result = yellow
	case "blue":
		result = blue
	case "magenta":
		result = magenta
	case "cyan":
		result = cyan
	case "white":
		result = white
	default:
		result = black
	}
	result = result + message + reset
	return result
}

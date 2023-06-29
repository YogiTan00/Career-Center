package utils

import (
	"CareerCenter/domain/valueobject"
	"CareerCenter/utils/exceptions"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	Email string
	Role  string
}

func (u *User) Admin() bool {
	if u.Role != string(valueobject.ROLE_ADMIN) {
		return false
	}
	return true
}

func GenerateToken(email string, role string) (*http.Cookie, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString([]byte("K6ct3EnySXFglJghxQrwLmu6RR403UzT"))
	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	}

	return cookie, nil
}

func ValidateTokenFromHeader(r *http.Request) (*User, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, exceptions.ErrCustomString("authorization header is missing")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("K6ct3EnySXFglJghxQrwLmu6RR403UzT"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, exceptions.ErrCustomString("invalid token")
	}

	user := &User{
		Email: claims["email"].(string),
		Role:  claims["role"].(string),
	}

	return user, nil
}

func RemoveJwtInCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Value:   "",
		Expires: time.Unix(0, 0),
	})
}

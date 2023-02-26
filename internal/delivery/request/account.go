package request

import "github.com/golang-jwt/jwt/v4"

type RequestRegister struct {
	Email    string `json:"username"`
	Nama     string `json:"nama"`
	Password string `json:"password"`
}

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

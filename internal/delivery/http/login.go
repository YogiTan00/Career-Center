package http

import (
	"CareerCenter/domain/entity"
	"CareerCenter/internal/delivery/request"
	response2 "CareerCenter/internal/delivery/response"
	"context"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

func (h *RegisterHandler) Login(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = context.TODO()
		req     request.RequestLogin
		decoder = json.NewDecoder(r.Body)
		jwtKey  = []byte("my_secret_key")
		creds   request.RequestLogin
	)
	errDecode := decoder.Decode(&req)
	if errDecode != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error decode data"))
		return
	}
	buildLogin := &entity.AccountDTO{
		Email:    req.Email,
		Password: req.Password,
	}

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &request.Claims{
		Email: creds.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	err := h.UCRegister.Login(ctx, buildLogin.Email, buildLogin.Password)
	if err != nil {
		response, errMap := response2.MapResponse(1, "wrong email or password")
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(response)
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
		response, errMap := response2.MapResponseLogin(0, "success login", tokenString)
		if errMap != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error mapping data"))
		}
		w.Write(response)
	}
}

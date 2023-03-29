package request

import (
	"CareerCenter/domain/entity/account"
)

type RequestRegister struct {
	Email    string `json:"email"`
	Nama     string `json:"name"`
	Password string `json:"password"`
}

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewRegisterRequest(req *RequestRegister) *account.AccountDTO {
	return &account.AccountDTO{
		Email:    req.Email,
		Nama:     req.Nama,
		Password: req.Password,
	}
}

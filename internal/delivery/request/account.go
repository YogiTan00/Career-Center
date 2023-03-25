package request

import (
	"CareerCenter/domain/entity"
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

func NewRegisterRequest(req *RequestRegister) *entity.AccountDTO {
	return &entity.AccountDTO{
		Email:    req.Email,
		Nama:     req.Nama,
		Password: req.Password,
	}
}

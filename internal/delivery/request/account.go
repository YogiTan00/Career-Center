package request

import (
	"CareerCenter/domain/entity/account"
)

type RequestRegister struct {
	Email    string `json:"email"`
	Nama     string `json:"name"`
	Password string `json:"password"`
}

func NewRegisterRequest(req *RequestRegister) *account.AccountDTO {
	return &account.AccountDTO{
		Email:    req.Email,
		Name:     req.Nama,
		Password: req.Password,
	}
}

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestPassword struct {
	OldPassword     string `json:"oldPassword"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

func NewPassword(req *RequestPassword) *account.UpdatePasswordDTO {
	return &account.UpdatePasswordDTO{
		OldPassword:     req.OldPassword,
		NewPassword:     req.NewPassword,
		ConfirmPassword: req.ConfirmPassword,
	}
}

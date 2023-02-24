package entity

import (
	"CareerCenter/utils"
	"errors"
)

type Account struct {
	email    string
	nama     string
	password string
}

type AccountDTO struct {
	Email    string
	Nama     string
	Password string
}

func NewAccount(dto *AccountDTO) (*Account, error) {
	err := dto.Validation()
	if err != nil {
		return nil, err
	}
	return &Account{
		email:    dto.Email,
		nama:     dto.Nama,
		password: dto.Password,
	}, nil
}

func (g *Account) GetEmail() string {
	return g.email
}
func (g *Account) GetNama() string {
	return g.nama
}
func (g *Account) GetPassword() string {
	return g.password
}

func (dto *AccountDTO) Validation() error {
	email := utils.ValitEmail(dto.Email)
	if email != true {
		return errors.New("error format email")
	}
	return nil
}

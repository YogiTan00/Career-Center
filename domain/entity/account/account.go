package account

import (
	"CareerCenter/utils"
	"errors"
	"time"
)

type Account struct {
	email     string
	nama      string
	password  string
	createdAt time.Time
	updateAt  time.Time
}

type AccountDTO struct {
	Email     string
	Nama      string
	Password  string
	CreatedAt time.Time
	UpdateAt  time.Time
}

func NewAccount(dto *AccountDTO) (*Account, error) {
	err := dto.Validation()
	if err != nil {
		return nil, err
	}
	return &Account{
		email:     dto.Email,
		nama:      dto.Nama,
		password:  dto.Password,
		createdAt: dto.CreatedAt,
		updateAt:  dto.UpdateAt,
	}, nil
}

func (dto *AccountDTO) Validation() error {
	email := utils.ValitEmail(dto.Email)
	if email != true {
		return errors.New("error format email")
	}

	timeNow := time.Now()
	dto.CreatedAt = timeNow
	dto.UpdateAt = timeNow

	return nil
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
func (g *Account) GetCreatedAt() time.Time {
	return g.createdAt
}

func (g *Account) GetUpdatedAt() time.Time {
	return g.updateAt
}

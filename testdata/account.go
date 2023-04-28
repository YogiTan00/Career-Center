package testdata

import (
	"CareerCenter/domain/entity/account"
	"CareerCenter/utils"
	"time"
)

func TestDataAccount() *account.Account {
	test := utils.RandomString(10)
	pw, _ := utils.HashPassword("abcd123")
	register, _ := account.NewAccount(&account.AccountDTO{
		Email:    test + "@gmail.com",
		Name:     "yogi",
		Password: pw,
	})
	return register
}

func TestDataAccountDTO() *account.AccountDTO {
	test := utils.RandomString(10)
	pw, _ := utils.HashPassword("abcd123")
	return &account.AccountDTO{
		Email:    test + "@gmail.com",
		Name:     "yogi",
		Password: pw,
	}
}

func TestDataPassword() *account.UpdatePasswordDTO {
	return &account.UpdatePasswordDTO{
		OldPassword:     "abcd123",
		NewPassword:     "abcd1234",
		ConfirmPassword: "abcd1234",
		UpdatedAt:       time.Now(),
	}
}

package testdata

import (
	"CareerCenter/domain/entity/account"
	"CareerCenter/utils"
)

func TestDataAccount() *account.Account {
	test := utils.RandomString(10)
	register, _ := account.NewAccount(&account.AccountDTO{
		Email:    test + "@gmail.com",
		Nama:     "yogi",
		Password: "abcd123",
	})
	return register
}

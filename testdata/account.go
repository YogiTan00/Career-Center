package testdata

import (
	"CareerCenter/domain/entity"
	"CareerCenter/utils"
)

func TestDataAccount() *entity.Account {
	test := utils.RandomString(10)
	register, _ := entity.NewAccount(&entity.AccountDTO{
		Email:    test + "@gmail.com",
		Nama:     "yogi",
		Password: "abcd123",
	})
	return register
}

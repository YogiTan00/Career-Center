package models

import (
	"time"
)

type AccountModel struct {
	Id         string    `dbq:"id"`
	Email      string    `dbq:"email"`
	Nama       string    `dbq:"name"`
	Password   string    `dbq:"password"`
	Role       string    `dbq:"role"`
	CodeOtp    string    `dbq:"code_otp"`
	ExpiredOtp time.Time `dbq:"expired_otp"`
	CreatedAt  time.Time `dbq:"created_at"`
	UpdateAt   time.Time `dbq:"updated_at"`
	DeletedAt  time.Time `dbq:"deleted_at"`
	RegisterBy string    `dbq:"register_by"`
}

func GetTableNameAccount() string {
	return "account"
}

func TableAccount() []string {
	return []string{
		"id",
		"email",
		"name",
		"password",
		"role",
		"code_otp",
		"expired_otp",
		"created_at",
		"updated_at",
		"deleted_at",
		"register_by",
	}
}

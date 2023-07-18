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
	CreatedAt  time.Time `dbq:"created_at"`
	UpdateAt   time.Time `dbq:"updated_at"`
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
		"created_at",
		"updated_at",
		"register_by",
	}
}

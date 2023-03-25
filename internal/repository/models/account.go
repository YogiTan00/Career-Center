package models

import (
	"time"
)

type AccountModel struct {
	Email     string    `dbq:"email"`
	Nama      string    `dbq:"name"`
	Password  string    `dbq:"password"`
	CreatedAt time.Time `dbq:"created_at"`
	UpdateAt  time.Time `dbq:"updated_at"`
}

func GetTableNameAccount() string {
	return "account"
}

func TableAccount() []string {
	return []string{
		"email",
		"name",
		"password",
		"created_at",
		"updated_at",
	}
}

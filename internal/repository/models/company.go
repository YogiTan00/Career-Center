package models

import (
	"time"
)

type CompanyModel struct {
	Id          string    `dbq:"id"`
	Email       string    `dbq:"email"`
	Name        string    `dbq:"name"`
	TypeCompany string    `dbq:"type_company"`
	Address     string    `dbq:"address"`
	Logo        string    `dbq:"logo"`
	Profile     string    `dbq:"profile"`
	Website     string    `dbq:"website"`
	Location    string    `dbq:"location"`
	CreatedAt   time.Time `dbq:"created_at"`
	UpdatedAt   time.Time `dbq:"updated_at"`
	DeletedAt   time.Time `dbq:"deleted_at"`
}

func GetTableNameCompany() string {
	return "company"
}

func TableCompany() []string {
	return []string{
		"id",
		"email",
		"name",
		"type_company",
		"address",
		"logo",
		"profile",
		"website",
		"location",
		"created_at",
		"updated_at",
	}
}

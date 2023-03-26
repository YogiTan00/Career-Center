package models

import (
	"CareerCenter/domain/valueobject"
	"time"
)

type CompanyModel struct {
	Id          string                      `dbq:"id"`
	Name        string                      `dbq:"name"`
	TypeCompany valueobject.TypeCompanyEnum `dbq:"type_company"`
	Address     string                      `dbq:"address"`
	Logo        string                      `dbq:"logo"`
	Profile     string                      `dbq:"profile"`
	Website     string                      `dbq:"website"`
	Location    string                      `dbq:"location"`
	CreatedAt   time.Time                   `dbq:"created_at"`
	UpdatedAt   time.Time                   `dbq:"updated_at"`
}

func GetTableNameCompany() string {
	return "company"
}

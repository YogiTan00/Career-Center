package entity

import (
	"CareerCenter/domain/valueobject"
	"time"
)

type Company struct {
	id          string
	name        string
	typeCompany *valueobject.TypeCompany
	address     string
	logo        string
	about       *About
	jobs        []*Jobs
	createdAt   time.Time
	updatedAt   time.Time
}

type About struct {
	profile  string
	website  string
	location string
}

type CompanyDTO struct {
	Id          string
	Name        string
	TypeCompany *valueobject.TypeCompany
	Address     string
	Logo        string
	About       *AboutDTO
	Jobs        []*JobsDTO
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type AboutDTO struct {
	Profile  string
	Website  string
	Location string
}

package models

import (
	"time"
)

type JobsModel struct {
	Id             string    `dbq:"id"`
	CompanyId      string    `dbq:"company_id"`
	Position       string    `dbq:"position"`
	Company        string    `dbq:"company"`
	Logo           string    `dbq:"logo"`
	Address        string    `dbq:"address"`
	Status         bool      `dbq:"status"`
	Qualification  string    `dbq:"qualification"`
	JobDescription string    `dbq:"job_description"`
	CreatedAt      time.Time `dbq:"created_at"`
	UpdatedAt      time.Time `dbq:"updated_at"`
	DeletedAt      time.Time `dbq:"deleted_at"`
}

func GetTableNameJobs() string {
	return "jobs"
}

func TableJob() []string {
	return []string{
		"id",
		"company_id",
		"position",
		"company",
		"logo",
		"address",
		"status",
		"qualification",
		"job_description",
		"created_at",
		"updated_at",
		"deleted_at",
	}
}

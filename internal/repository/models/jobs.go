package models

import (
	"time"
)

type JobsModel struct {
	Id          string    `dbq:"id"`
	Position    string    `dbq:"position"`
	Company     string    `dbq:"company"`
	Logo        string    `dbq:"logo"`
	Address     string    `dbq:"address"`
	CreatedDate time.Time `dbq:"created_at"`
}

func GetTableNameJobs() string {
	return "jobs"
}

func TableJobs() []string {
	return []string{
		"id",
		"position",
		"company",
		"logo",
		"address",
		"created_date",
	}
}

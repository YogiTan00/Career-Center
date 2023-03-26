package models

import (
	"time"
)

type JobsModel struct {
	Id             string    `dbq:"id"`
	Position       string    `dbq:"position"`
	Company        string    `dbq:"company"`
	Logo           string    `dbq:"logo"`
	Address        string    `dbq:"address"`
	Status         bool      `dbq:"status"`
	SendDate       time.Time `dbq:"send_date"`
	Qualification  string    `dbq:"qualification"`
	JobDescription string    `dbq:"job_description"`
	Category       string    `dbq:"category"`
	Description    string    `dbq:"description"`
	CreatedAt      time.Time `dbq:"created_at"`
	UpdatedAt      time.Time `dbq:"updated_at"`
}

func GetTableNameJobs() string {
	return "jobs"
}

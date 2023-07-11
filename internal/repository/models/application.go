package models

import "time"

type ApplicationModel struct {
	Id          string    `dbq:"id"`
	CompanyId   string    `dbq:"company_id"`
	JobId       string    `dbq:"job_id"`
	Email       string    `dbq:"email"`
	Name        string    `dbq:"name"`
	Skill       string    `dbq:"skill"`
	PhoneNumber string    `dbq:"phone_number"`
	CvResume    string    `dbq:"cv_resume"`
	Portofolio  string    `dbq:"portofolio"`
	CreatedAt   time.Time `dbq:"created_at"`
	UpdateAt    time.Time `dbq:"updated_at"`
}

func GetTableNameApplication() string {
	return "application"
}

func TableApplication() []string {
	return []string{
		"id",
		"company_id",
		"job_id",
		"email",
		"name",
		"skill",
		"phone_number",
		"cv_resume",
		"portofolio",
		"created_at",
		"updated_at",
	}
}

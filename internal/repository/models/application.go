package models

import "time"

type ApplicationModel struct {
	Id          string    `dbq:"id"`
	CompanyId   string    `dbq:"company_d"`
	Email       string    `dbq:"email"`
	Nama        string    `dbq:"name"`
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

package models

import (
	"time"
)

type ProfileModel struct {
	Email       string    `dbq:"email"`
	Name        string    `dbq:"name"`
	Photo       string    `dbq:"photo"`
	Skill       string    `dbq:"skill"`
	PhoneNumber string    `dbq:"phone_number"`
	Ability     string    `dbq:"ability"`
	Language    string    `dbq:"language"`
	CvResume    string    `dbq:"cv_resume"`
	Portofolio  string    `dbq:"portofolio"`
	CreatedAt   time.Time `dbq:"created_at"`
	UpdatedAt   time.Time `dbq:"updated_at"`
}

func GetTableNameProfile() string {
	return "profile"
}

func TableProfile() []string {
	return []string{
		"email",
		"name",
		"photo",
		"skill",
		"phone_number",
		"ability",
		"language",
		"cv_resume",
		"portofolio",
		"created_at",
		"updated_at",
	}
}

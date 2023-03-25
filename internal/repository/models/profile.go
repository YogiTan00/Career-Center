package models

import (
	"time"
)

type ProfileModel struct {
	Email      string    `dbq:"email"`
	Nama       string    `dbq:"name"`
	Photo      string    `dbq:"photo"`
	Skill      string    `dbq:"skill"`
	NoTlp      string    `dbq:"no_tlp"`
	Ability    string    `dbq:"ability"`
	Language   string    `dbq:"language"`
	Cv         string    `dbq:"cv"`
	Portofolio string    `dbq:"portofolio"`
	CreatedAt  time.Time `dbq:"created_at"`
	UpdateAt   time.Time `dbq:"updated_at"`
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
		"no_tlp",
		"language",
		"cv",
		"portofolio",
		"created_at",
		"updated_at",
	}
}

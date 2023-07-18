package profile

import "time"

type EducationModel struct {
	Id             string    `dbq:"id"`
	Email          string    `dbq:"email"`
	Level          string    `dbq:"level"`
	Name           string    `dbq:"name"`
	Major          string    `dbq:"major"`
	StillEducation bool      `dbq:"still_education"`
	StartEdu       time.Time `dbq:"start_education"`
	EndEdu         time.Time `dbq:"end_education"`
	Description    string    `dbq:"description"`
	CreatedAt      time.Time `dbq:"created_at"`
	UpdatedAt      time.Time `dbq:"updated_at"`
}

func GetTableNameEducation() string {
	return "education"
}

func TableEducation() []string {
	return []string{
		"id",
		"email",
		"level",
		"name",
		"major",
		"still_education",
		"start_education",
		"end_education",
		"description",
		"created_at",
		"updated_at",
	}
}

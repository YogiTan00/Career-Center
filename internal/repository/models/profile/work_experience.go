package profile

import "time"

type WorkExperienceModel struct {
	Id              string    `dbq:"id"`
	Email           string    `dbq:"email"`
	Name            string    `dbq:"name"`
	SkillExperience string    `dbq:"skill_experience"`
	StillWorking    bool      `dbq:"still_working"`
	StartWork       time.Time `dbq:"start_work"`
	EndWork         time.Time `dbq:"end_work"`
	Description     string    `dbq:"description"`
	CreatedAt       time.Time `dbq:"created_at"`
	UpdatedAt       time.Time `dbq:"updated_at"`
}

func GetTableNameWorkExperience() string {
	return "work_experience"
}

func TableWorkExperience() []string {
	return []string{
		"id",
		"email",
		"name",
		"skill_experience",
		"still_working",
		"start_work",
		"end_work",
		"description",
		"created_at",
		"updated_at",
	}
}

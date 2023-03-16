package entity

type ProfileUser struct {
	name           string
	photo          string
	skill          string
	email          string
	noTlp          int
	workExperience *WorkExperience
}

type WorkExperience struct {
	skillExperience string
	name            string
	date            string
	description     string
}

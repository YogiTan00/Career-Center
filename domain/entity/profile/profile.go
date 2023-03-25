package profile

import (
	"CareerCenter/utils"
	"errors"
	"time"
)

type ProfileUser struct {
	name           string
	photo          string
	skill          string
	email          string
	noTlp          string
	workExperience *WorkExperience
	education      *Education
	ability        []string
	language       []string
	cv             string
	portofolio     string
	createdAt      time.Time
	updateAt       time.Time
}

type ProfileUserDTO struct {
	Name           string
	Photo          string
	Skill          string
	Email          string
	NoTlp          string
	WorkExperience *WorkExperienceDTO
	Education      *EducationDTO
	Ability        []string
	Language       []string
	Cv             string
	Portofolio     string
	CreatedAt      time.Time
	UpdateAt       time.Time
}

func NewProfile(dto *ProfileUserDTO) (*ProfileUser, error) {
	workExperiencet := NewWorkExperience(dto.WorkExperience)
	education := NewEducation(dto.Education)
	return &ProfileUser{
		name:           dto.Name,
		photo:          dto.Photo,
		skill:          dto.Skill,
		email:          dto.Email, //Validation
		noTlp:          dto.NoTlp,
		workExperience: workExperiencet,
		education:      education,
		ability:        dto.Ability,
		language:       dto.Language,
		cv:             dto.Cv,         //Validation
		portofolio:     dto.Portofolio, //Validation
		createdAt:      dto.CreatedAt,
		updateAt:       dto.UpdateAt,
	}, nil
}

func (dto *ProfileUserDTO) Validation() error {
	email := utils.ValitEmail(dto.Email)
	if email != true {
		return errors.New("error format email")
	}
	return nil
}

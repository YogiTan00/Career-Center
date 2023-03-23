package profile

import (
	"CareerCenter/utils"
	"errors"
	uuid2 "github.com/google/uuid"
)

type ProfileUser struct {
	id             string
	name           string
	photo          *PhotoProfile
	skill          string
	email          string
	noTlp          int
	workExperience *WorkExperience
	education      *Education
	ability        []string
	language       []string
	cv             string
	portofolio     string
}

type ProfileUserDTO struct {
	Id             string
	Name           string
	Photo          *PhotoProfileDTO
	Skill          string
	Email          string
	NoTlp          int
	WorkExperience *WorkExperienceDTO
	Education      *EducationDTO
	Ability        []string
	Language       []string
	Cv             string
	Portofolio     string
}

func NewProfile(dto *ProfileUserDTO) (*ProfileUser, error) {
	uuid, err := uuid2.NewUUID()
	if err != nil {
		return nil, err
	}
	photo := NewPhoto(dto.Photo)
	workExperiencet := NewWorkExperience(dto.WorkExperience)
	education := NewEducation(dto.Education)
	return &ProfileUser{
		id:             uuid.String(),
		name:           dto.Name,
		photo:          photo,
		skill:          dto.Skill,
		email:          dto.Email, //Validation
		noTlp:          dto.NoTlp,
		workExperience: workExperiencet,
		education:      education,
		ability:        dto.Ability,
		language:       dto.Language,
		cv:             dto.Cv,         //Validation
		portofolio:     dto.Portofolio, //Validation
	}, nil
}

func (dto *ProfileUserDTO) Validation() error {
	email := utils.ValitEmail(dto.Email)
	if email != true {
		return errors.New("error format email")
	}
	return nil
}

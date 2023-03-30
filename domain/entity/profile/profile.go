package profile

import (
	"CareerCenter/utils"
	"CareerCenter/utils/exceptions"
	"time"
)

type ProfileUser struct {
	name           string
	photo          string
	skill          string
	email          string
	phoneNumber    string
	workExperience WorkExperience
	education      Education
	ability        []string
	language       []string
	cvResume       string
	portofolio     string
	createdAt      time.Time
	updatedAt      time.Time
}

type ProfileUserDTO struct {
	Name           string
	Photo          string
	Skill          string
	Email          string
	PhoneNumber    string
	WorkExperience WorkExperienceDTO
	Education      EducationDTO
	Ability        []string
	Language       []string
	CvResume       string
	Portofolio     string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewProfile(dto *ProfileUserDTO) (*ProfileUser, error) {
	err := dto.Validation()
	if err != nil {
		return nil, err
	}
	workExperiencet := NewWorkExperience(dto.WorkExperience)
	education := NewEducation(dto.Education)
	return &ProfileUser{
		name:           dto.Name,
		photo:          dto.Photo,
		skill:          dto.Skill,
		email:          dto.Email,
		phoneNumber:    dto.PhoneNumber,
		workExperience: workExperiencet,
		education:      education,
		ability:        dto.Ability,
		language:       dto.Language,
		cvResume:       dto.CvResume,
		portofolio:     dto.Portofolio,
		createdAt:      dto.CreatedAt,
		updatedAt:      dto.UpdatedAt,
	}, nil
}

func (dto *ProfileUserDTO) Validation() error {
	if len(dto.Email) > 0 {
		email := utils.ValitEmail(dto.Email)
		if email != true {
			return exceptions.ErrCustomString("error format email")
		}
	}

	if len(dto.PhoneNumber) > 0 {
		if !utils.IsNumber(dto.PhoneNumber) {
			return exceptions.ErrCustomString("invalid input only number")
		}
	}

	timeNow := time.Now()
	if dto.CreatedAt.IsZero() && dto.UpdatedAt.IsZero() {
		dto.CreatedAt = timeNow
		dto.UpdatedAt = timeNow
	}

	if !dto.CreatedAt.IsZero() && dto.UpdatedAt.IsZero() {
		dto.UpdatedAt = timeNow
	}
	return nil
}

func (data *ProfileUser) GetName() string {
	return data.name
}
func (data *ProfileUser) GetPhoto() string {
	return data.photo
}
func (data *ProfileUser) GetSkill() string {
	return data.skill
}
func (data *ProfileUser) GetEmail() string {
	return data.email
}
func (data *ProfileUser) GetPhoneNumber() string {
	return data.phoneNumber
}
func (data *ProfileUser) GetWorkExperience() WorkExperience {
	return data.workExperience
}
func (data *ProfileUser) GetEducation() Education {
	return data.education
}
func (data *ProfileUser) GetAbility() []string {
	return data.ability
}
func (data *ProfileUser) GetLanguage() []string {
	return data.language
}
func (data *ProfileUser) GetCvResume() string {
	return data.cvResume
}
func (data *ProfileUser) GetPortofolio() string {
	return data.portofolio
}
func (data *ProfileUser) GetCreatedAt() time.Time {
	return data.createdAt
}
func (data *ProfileUser) GetUpdatedAt() time.Time {
	return data.updatedAt
}

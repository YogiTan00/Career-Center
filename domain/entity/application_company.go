package entity

import (
	"CareerCenter/domain/entity/profile"
	"CareerCenter/utils"
	"time"

	uuid2 "github.com/google/uuid"
)

type Application struct {
	id          string
	companyId   string
	email       string
	name        string
	skill       string
	phoneNumber string
	cvResume    string
	portofolio  string
	createdAt   time.Time
	updatedAt   time.Time
}

type ApplicationDTO struct {
	Id          string
	CompanyId   string
	Email       string
	Name        string
	Skill       string
	PhoneNumber string
	CvResume    string
	Portofolio  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewApplication(dto *ApplicationDTO) (*Application, error) {
	uuid, err := uuid2.NewUUID()
	if err != nil {
		return nil, err
	}
	timeNow := time.Now()
	return &Application{
		id:          uuid.String(),
		companyId:   dto.CompanyId,
		email:       dto.Email,
		name:        dto.Name,
		skill:       dto.Skill,
		phoneNumber: dto.PhoneNumber,
		cvResume:    dto.CvResume,
		portofolio:  dto.Portofolio,
		createdAt:   timeNow,
		updatedAt:   timeNow,
	}, nil
}

func NewApplicationFromProfile(dto *profile.ProfileUserDTO, companyId string) (*Application, error) {
	uuid, err := uuid2.NewUUID()
	if err != nil {
		return nil, err
	}
	company, err := utils.ValitUuId(companyId)
	if err != nil {
		return nil, err
	}
	timeNow := time.Now()
	return &Application{
		id:          uuid.String(),
		companyId:   company,
		email:       dto.Email,
		name:        dto.Name,
		skill:       dto.Skill,
		phoneNumber: dto.PhoneNumber,
		cvResume:    dto.CvResume,
		portofolio:  dto.Portofolio,
		createdAt:   timeNow,
		updatedAt:   timeNow,
	}, nil
}

func (data *Application) GetId() string {
	return data.id
}
func (data *Application) GetCompanyId() string {
	return data.companyId
}
func (data *Application) GetEmail() string {
	return data.email
}
func (data *Application) GetName() string {
	return data.name
}
func (data *Application) GetSkill() string {
	return data.skill
}
func (data *Application) GetPhoneNumber() string {
	return data.phoneNumber
}
func (data *Application) GetCvResume() string {
	return data.cvResume
}
func (data *Application) GetPortofolio() string {
	return data.portofolio
}
func (data *Application) GetCreatedAt() time.Time {
	return data.createdAt
}
func (data *Application) GetUpdatedAt() time.Time {
	return data.updatedAt
}

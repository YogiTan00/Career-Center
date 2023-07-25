package entity

import (
	"CareerCenter/domain/entity/profile"
	"CareerCenter/domain/valueobject"
	"CareerCenter/utils/exceptions"
	"github.com/google/uuid"
	"time"
)

type Company struct {
	id          string
	email       string
	name        string
	typeCompany *valueobject.TypeCompany
	address     string
	logo        string
	about       *About
	jobs        []*Jobs
	createdAt   time.Time
	updatedAt   time.Time
}

type About struct {
	profile  string
	website  string
	location string
}

type CompanyDTO struct {
	Id          string
	Email       string
	Name        string
	TypeCompany *valueobject.TypeCompany
	Address     string
	Logo        string
	About       *AboutDTO
	Jobs        []*JobsDTO
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type AboutDTO struct {
	Profile  string
	Website  string
	Location string
}

func NewAboutCompany(dto *AboutDTO) *About {
	return &About{
		profile:  dto.Profile,
		website:  dto.Website,
		location: dto.Location,
	}
}

func NewCompany(dto *CompanyDTO) (*Company, error) {
	err := dto.Validation()
	if err != nil {
		return nil, err
	}
	about := NewAboutCompany(dto.About)
	timeNow := time.Now()
	return &Company{
		id:          uuid.New().String(),
		email:       dto.Email,
		name:        dto.Name,
		typeCompany: dto.TypeCompany,
		address:     dto.Address,
		logo:        dto.Logo,
		about:       about,
		createdAt:   timeNow,
		updatedAt:   timeNow,
	}, nil
}

func (dto *CompanyDTO) Validation() error {
	if dto.Email == "" {
		return exceptions.ErrIsRequire("email")
	}

	return nil
}

func (data *Company) GetId() string {
	return data.id
}
func (data *Company) GetEmail() string {
	return data.email
}
func (data *Company) GetName() string {
	return data.name
}
func (data *Company) GetTypeCompany() *valueobject.TypeCompany {
	return data.typeCompany
}
func (data *Company) GetAddress() string {
	return data.address
}
func (data *Company) GetLogo() string {
	return data.logo
}
func (data *Company) About() *About {
	return data.about
}
func (data *Company) GetCreatedAt() time.Time {
	return data.createdAt
}
func (data *Company) GetUpdateAt() time.Time {
	return data.updatedAt
}

func (data *About) GetProfile() string {
	return data.profile
}
func (data *About) GetWebsite() string {
	return data.website
}
func (data *About) GetLocation() string {
	return data.location
}

func (data *CompanyDTO) SetLogo(dto *profile.ProfileUserDTO) {
	data.Logo = dto.Photo
}

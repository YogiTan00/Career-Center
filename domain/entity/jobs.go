package entity

import (
	uuid2 "github.com/google/uuid"
	"time"
)

type Jobs struct {
	id             string
	position       string
	company        string
	logo           string
	address        string
	status         bool
	sendDate       time.Time
	qualification  string
	jobDescription string
	category       string
	description    string
	createdAt      time.Time
	updatedAt      time.Time
}

type JobsDTO struct {
	Id             string
	Position       string
	Company        string
	Logo           string
	Address        string
	Status         bool
	SendDate       time.Time
	Qualification  string
	JobDescription string
	Category       string
	Description    string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func NewJobs(dto *JobsDTO) (*Jobs, error) {
	uuid, err := uuid2.NewUUID()
	if err != nil {
		return nil, err
	}
	return &Jobs{
		id:        uuid.String(),
		position:  dto.Position,
		company:   dto.Company,
		logo:      dto.Logo,
		address:   dto.Address,
		createdAt: dto.CreatedAt,
	}, nil
}

func (g *Jobs) GetId() string {
	return g.id
}

func (g *Jobs) GetPosition() string {
	return g.position
}

func (g *Jobs) GetCompany() string {
	return g.company
}

func (g *Jobs) GetLogo() string {
	return g.logo
}

func (g *Jobs) GetAddress() string {
	return g.address
}

func (g *Jobs) GetCreatedAt() time.Time {
	return g.createdAt
}

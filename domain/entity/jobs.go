package entity

import (
	"CareerCenter/utils/exceptions"
	"github.com/google/uuid"
	"time"
)

type Jobs struct {
	id             string
	companyId      string
	position       string
	company        string
	logo           string
	address        string
	status         bool
	qualification  string
	jobDescription string
	createdAt      time.Time
	updatedAt      time.Time
	deletedAt      time.Time
	applicant      int
}

type JobsDTO struct {
	Id             string
	CompanyId      string
	Position       string
	Company        string
	Logo           string
	Address        string
	Status         bool
	Qualification  string
	JobDescription string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	deletedAt      time.Time
	Applicant      int
}

func NewJobs(dto *JobsDTO) (*Jobs, error) {
	timeNow := time.Now()
	err := dto.Validation()
	if err != nil {
		return nil, err
	}
	return &Jobs{
		id:             uuid.New().String(),
		companyId:      dto.CompanyId,
		position:       dto.Position,
		company:        dto.Company,
		logo:           dto.Logo,
		address:        dto.Address,
		status:         dto.Status,
		qualification:  dto.Qualification,
		jobDescription: dto.JobDescription,
		createdAt:      timeNow,
		updatedAt:      timeNow,
	}, nil
}

func (dto *JobsDTO) Validation() error {
	if dto.CompanyId == "" {
		return exceptions.ErrIsRequire("company")
	}
	if dto.Position == "" {
		return exceptions.ErrIsRequire("position")
	}
	if dto.Address == "" {
		return exceptions.ErrIsRequire("address")
	}
	if dto.Qualification == "" {
		return exceptions.ErrIsRequire("qualification")
	}
	if dto.JobDescription == "" {
		return exceptions.ErrIsRequire("job description")
	}

	return nil
}

func (g *Jobs) GetId() string {
	return g.id
}

func (g *Jobs) GetCompanyId() string {
	return g.companyId
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

func (g *Jobs) GetStatus() bool {
	return g.status
}

func (g *Jobs) GetQualificationi() string {
	return g.qualification
}

func (g *Jobs) GetJobDescription() string {
	return g.jobDescription
}

func (g *Jobs) GetCreatedAt() time.Time {
	return g.createdAt
}

func (g *Jobs) GetUpdateAt() time.Time {
	return g.updatedAt
}

func (g *Jobs) GetDeletedAt() time.Time {
	return g.deletedAt
}

func (g *Jobs) GetApplicant() int {
	return g.applicant
}

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
	category       string
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
	Category       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
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
		category:       dto.Category,
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

func (dto *JobsDTO) SetUpdate(data *JobsDTO) *Jobs {
	//if len(data.CompanyId) > 0 {
	//	dto.CompanyId = data.CompanyId
	//}
	//if len(data.Position) > 0 {
	//	dto.Position = data.Position
	//}
	//if len(data.Company) > 0 {
	//	dto.Company = data.Company
	//}
	//if len(data.Logo) > 0 {
	//	dto.Logo = data.Logo
	//}
	//if len(data.Address) > 0 {
	//	dto.Address = data.Address
	//}
	//if data.Status == true {
	//	dto.Status = data.Status
	//}
	//if len(data.Qualification) > 0 {
	//	dto.Qualification = data.Qualification
	//}
	//if len(data.JobDescription) > 0 {
	//	dto.JobDescription = data.JobDescription
	//}
	//if len(data.Category) > 0 {
	//	dto.Category = data.Category
	//}
	timeNow := time.Now()
	return &Jobs{
		id:             dto.Id,
		companyId:      data.CompanyId,
		position:       data.Position,
		company:        data.Company,
		logo:           data.Logo,
		address:        data.Address,
		status:         data.Status,
		qualification:  data.Qualification,
		jobDescription: data.JobDescription,
		category:       data.Category,
		createdAt:      dto.CreatedAt,
		updatedAt:      timeNow,
		deletedAt:      dto.DeletedAt,
	}
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

func (g *Jobs) GetCategory() string {
	return g.category
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

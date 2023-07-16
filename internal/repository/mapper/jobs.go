package mapper

import (
	"CareerCenter/domain/entity"
	"CareerCenter/internal/repository/models"
	"github.com/rocketlaunchr/dbq/v2"
)

func EntityToModelJobs(m *entity.Jobs) *models.JobsModel {
	return &models.JobsModel{
		Id:             m.GetId(),
		CompanyId:      m.GetCompanyId(),
		Position:       m.GetPosition(),
		Company:        m.GetCompany(),
		Logo:           m.GetLogo(),
		Address:        m.GetAddress(),
		Status:         m.GetStatus(),
		Qualification:  m.GetQualificationi(),
		JobDescription: m.GetJobDescription(),
		Category:       m.GetCategory(),
		CreatedAt:      m.GetCreatedAt(),
		UpdatedAt:      m.GetUpdateAt(),
		DeletedAt:      m.GetDeletedAt(),
	}
}

func ModelToEntityJobs(m *models.JobsModel) *entity.JobsDTO {
	return &entity.JobsDTO{
		Id:             m.Id,
		CompanyId:      m.CompanyId,
		Position:       m.Position,
		Company:        m.Company,
		Logo:           m.Logo,
		Address:        m.Address,
		Status:         m.Status,
		Qualification:  m.Qualification,
		JobDescription: m.JobDescription,
		Category:       m.Category,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
	}
}

func ModelToEntityListJobs(m []*models.JobsModel) []*entity.JobsDTO {
	listJobs := make([]*entity.JobsDTO, 0)
	for _, data := range m {
		jobs := ModelToEntityJobs(data)
		listJobs = append(listJobs, jobs)
	}
	return listJobs
}

func EntityJobToInterface(data *entity.Jobs) []interface{} {
	return dbq.Struct(EntityToModelJobs(data))
}

func DomainJobToInterface(domain *entity.Jobs) (dbqStruct []interface{}) {
	dbqStruct = append(dbqStruct, EntityJobToInterface(domain))
	return
}

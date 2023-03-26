package mapper

import (
	"CareerCenter/domain/entity"
	"CareerCenter/internal/repository/models"
)

func EntityToModelJobs(m *entity.Jobs) *models.JobsModel {
	return &models.JobsModel{
		Id:        m.GetId(),
		Position:  m.GetPosition(),
		Company:   m.GetCompany(),
		Logo:      m.GetLogo(),
		Address:   m.GetAddress(),
		CreatedAt: m.GetCreatedAt(),
	}
}

func ModelToEntityJobs(m *models.JobsModel) *entity.JobsDTO {
	return &entity.JobsDTO{
		Id:             m.Id,
		Position:       m.Position,
		Company:        m.Company,
		Logo:           m.Logo,
		Address:        m.Address,
		Status:         m.Status,
		SendDate:       m.SendDate,
		Qualification:  m.Qualification,
		JobDescription: m.JobDescription,
		Category:       m.Category,
		Description:    m.Description,
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

package mapper

import (
	"CareerCenter/domain/entity"
	"CareerCenter/domain/valueobject"
	"CareerCenter/internal/repository/models"
)

func ModelToEntityCompany(m *models.CompanyModel) *entity.CompanyDTO {
	typeCompany := valueobject.NewTypeCompany(m.TypeCompany)
	return &entity.CompanyDTO{
		Id:          m.Id,
		Name:        m.Name,
		TypeCompany: typeCompany,
		Address:     m.Address,
		Logo:        m.Logo,
		About: &entity.AboutDTO{
			Profile:  m.Profile,
			Website:  m.Website,
			Location: m.Location,
		},
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func ModelToEntityListCompany(m []*models.CompanyModel) []*entity.CompanyDTO {
	listJobs := make([]*entity.CompanyDTO, 0)
	for _, data := range m {
		jobs := ModelToEntityCompany(data)
		listJobs = append(listJobs, jobs)
	}
	return listJobs
}

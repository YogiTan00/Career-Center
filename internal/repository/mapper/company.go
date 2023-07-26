package mapper

import (
	"CareerCenter/domain/entity"
	"CareerCenter/domain/valueobject"
	"CareerCenter/internal/repository/models"

	"github.com/rocketlaunchr/dbq/v2"
)

func EntityToModelCompany(company *entity.Company) *models.CompanyModel {
	return &models.CompanyModel{
		Id:          company.GetId(),
		Email:       company.GetEmail(),
		Name:        company.GetName(),
		TypeCompany: company.GetTypeCompany().StringCompany(),
		Address:     company.GetAddress(),
		Logo:        company.GetLogo(),
		Profile:     company.About().GetProfile(),
		Website:     company.About().GetWebsite(),
		Location:    company.About().GetLocation(),
		CreatedAt:   company.GetCreatedAt(),
		UpdatedAt:   company.GetUpdateAt(),
	}
}

func ModelToEntityCompany(m *models.CompanyModel) *entity.CompanyDTO {
	typeCompany := valueobject.NewTypeCompanyFromString(m.TypeCompany)
	return &entity.CompanyDTO{
		Id:          m.Id,
		Email:       m.Email,
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

func EntityCompanyToInterface(data *entity.Company) []interface{} {
	return dbq.Struct(EntityToModelCompany(data))
}

func DomainCompanyToInterface(domain *entity.Company) (dbqStruct []interface{}) {
	dbqStruct = append(dbqStruct, EntityCompanyToInterface(domain))
	return
}

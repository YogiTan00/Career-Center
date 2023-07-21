package mapper

import (
	"CareerCenter/domain/entity"
	"CareerCenter/domain/valueobject"
	"CareerCenter/internal/repository/models"
	"github.com/rocketlaunchr/dbq/v2"
)

func EntityApplicationToModel(m *entity.Application) *models.ApplicationModel {
	return &models.ApplicationModel{
		Id:          m.GetId(),
		CompanyId:   m.GetCompanyId(),
		JobId:       m.GetJobId(),
		Email:       m.GetEmail(),
		Name:        m.GetName(),
		Skill:       m.GetSkill(),
		PhoneNumber: m.GetPhoneNumber(),
		CvResume:    m.GetCvResume(),
		Portofolio:  m.GetPortofolio(),
		CreatedAt:   m.GetCreatedAt(),
		UpdateAt:    m.GetUpdatedAt(),
	}
}

func ModelApplicationToEntity(m *models.ApplicationModel) *entity.ApplicationDTO {
	status := valueobject.NewTypeStatusApplicantFromStringPointer(m.Status)

	data := &entity.ApplicationDTO{
		Id:          m.Id,
		CompanyId:   m.CompanyId,
		JobId:       m.JobId,
		Email:       m.Email,
		Name:        m.Name,
		Skill:       m.Skill,
		PhoneNumber: m.PhoneNumber,
		CvResume:    m.CvResume,
		Portofolio:  m.Portofolio,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdateAt,
		DeletedAt:   m.DeletedAt,
		Status:      status,
	}
	return data
}

func EntityApplicationToInterface(data *entity.Application) []interface{} {
	return dbq.Struct(EntityApplicationToModel(data))
}

func DomainApplicationToInterface(domain *entity.Application) (dbqStruct []interface{}) {
	dbqStruct = append(dbqStruct, EntityApplicationToInterface(domain))
	return
}

func ModelApplicationListToEntity(m []*models.ApplicationModel) []*entity.ApplicationDTO {
	listApplicant := make([]*entity.ApplicationDTO, 0)
	for _, data := range m {
		apply := ModelApplicationToEntity(data)
		listApplicant = append(listApplicant, apply)
	}
	return listApplicant
}

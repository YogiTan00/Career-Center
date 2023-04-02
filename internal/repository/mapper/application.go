package mapper

import (
	"CareerCenter/domain/entity"
	"CareerCenter/internal/repository/models"
	"github.com/rocketlaunchr/dbq/v2"
)

func EntityApplicationToModel(m *entity.Application) *models.ApplicationModel {
	return &models.ApplicationModel{
		Id:          m.GetId(),
		CompanyId:   m.GetCompanyId(),
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
	data := &entity.ApplicationDTO{
		Id:          m.Id,
		CompanyId:   m.CompanyId,
		Email:       m.Email,
		Name:        m.Name,
		Skill:       m.Skill,
		PhoneNumber: m.PhoneNumber,
		CvResume:    m.CvResume,
		Portofolio:  m.Portofolio,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdateAt,
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

package profile

import (
	"CareerCenter/domain/entity/profile"
	profile2 "CareerCenter/internal/repository/models/profile"
	"CareerCenter/utils"
	"github.com/rocketlaunchr/dbq/v2"
)

func ModelProfileToEntity(m *profile2.ProfileModel) *profile.ProfileUserDTO {
	listAbility := utils.SplitTextToArray(m.Ability)
	listLanguage := utils.SplitTextToArray(m.Language)
	data := &profile.ProfileUserDTO{
		Id:          m.Id,
		Name:        m.Name,
		Photo:       m.Photo,
		Skill:       m.Skill,
		Email:       m.Email,
		PhoneNumber: m.PhoneNumber,
		Ability:     listAbility,
		Language:    listLanguage,
		CvResume:    m.CvResume,
		Portofolio:  m.Portofolio,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
	return data
}

func EntityProfileToModel(m *profile.ProfileUser) *profile2.ProfileModel {
	listAbility := utils.JoinTextFromArray(m.GetAbility())
	listLanguage := utils.JoinTextFromArray(m.GetLanguage())
	data := &profile2.ProfileModel{
		Id:          m.GetId(),
		Name:        m.GetName(),
		Photo:       m.GetEmail(),
		Skill:       m.GetSkill(),
		Email:       m.GetEmail(),
		PhoneNumber: m.GetPhoneNumber(),
		Ability:     listAbility,
		Language:    listLanguage,
		CvResume:    m.GetCvResume(),
		Portofolio:  m.GetPortofolio(),
		CreatedAt:   m.GetCreatedAt(),
		UpdatedAt:   m.GetUpdatedAt(),
	}
	return data
}

func EntityProfileToInterface(data *profile.ProfileUser) []interface{} {
	return dbq.Struct(EntityProfileToModel(data))
}

func DomainProfileToInterface(domain *profile.ProfileUser) (dbqStruct []interface{}) {
	dbqStruct = append(dbqStruct, EntityProfileToInterface(domain))
	return
}

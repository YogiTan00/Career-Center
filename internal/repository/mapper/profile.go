package mapper

import (
	"CareerCenter/domain/entity/profile"
	"CareerCenter/internal/repository/models"
	"CareerCenter/utils"
)

func ModelProfileToEntity(m *models.ProfileModel) (*profile.ProfileUserDTO, error) {
	listAbility := utils.SplitTextToArray(m.Ability)
	listLanguage := utils.SplitTextToArray(m.Language)
	data := &profile.ProfileUserDTO{
		Name:           m.Nama,
		Photo:          m.Photo,
		Skill:          m.Skill,
		Email:          m.Email,
		NoTlp:          m.NoTlp,
		WorkExperience: nil,
		Education:      nil,
		Ability:        listAbility,
		Language:       listLanguage,
		Cv:             m.Cv,
		Portofolio:     m.Portofolio,
		CreatedAt:      m.CreatedAt,
		UpdateAt:       m.UpdateAt,
	}
	return data, nil
}

package profile

import (
	"CareerCenter/domain/entity/profile"
	profile2 "CareerCenter/internal/repository/models/profile"
	"github.com/rocketlaunchr/dbq/v2"
)

func ModelWorkExperienceToEntity(m *profile2.WorkExperienceModel) *profile.WorkExperienceDTO {
	data := &profile.WorkExperienceDTO{
		Id:              m.Id,
		Email:           m.Email,
		SkillExperience: m.SkillExperience,
		Name:            m.Name,
		StillWorking:    m.StillWorking,
		DateRange: profile.DateRangeWorkDTO{
			Start: m.StartWork,
			End:   m.EndWork,
		},
		Description: m.Description,
	}
	return data
}

func EntityWorkExperienceToModel(m *profile.WorkExperience) *profile2.WorkExperienceModel {
	data := &profile2.WorkExperienceModel{
		Id:              m.GetId(),
		Email:           m.GetEmail(),
		SkillExperience: m.GetSkillExperience(),
		Name:            m.GetName(),
		StillWorking:    m.GetStillWorking(),
		StartWork:       m.GetStartWork(),
		EndWork:         m.GetEndWork(),
		Description:     m.GetDescription(),
	}

	return data
}

func ModelListWorkExperienceToEntity(m []*profile2.WorkExperienceModel) []*profile.WorkExperienceDTO {
	listWorkExperience := make([]*profile.WorkExperienceDTO, 0)
	for _, data := range m {
		workExperiencet := ModelWorkExperienceToEntity(data)
		listWorkExperience = append(listWorkExperience, workExperiencet)
	}
	return listWorkExperience
}

func EntityWorkExperienceToInterface(data *profile.WorkExperience) []interface{} {
	return dbq.Struct(EntityWorkExperienceToModel(data))
}

func DomainWorkExperienceToInterface(domain *profile.WorkExperience) (dbqStruct []interface{}) {
	dbqStruct = append(dbqStruct, EntityWorkExperienceToInterface(domain))
	return
}

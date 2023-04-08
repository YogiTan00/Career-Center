package profile

import (
	"CareerCenter/domain/entity/profile"
	"CareerCenter/domain/valueobject"
	profile2 "CareerCenter/internal/repository/models/profile"
	"github.com/rocketlaunchr/dbq/v2"
)

func ModelEducationToEntity(m *profile2.EducationModel) *profile.EducationDTO {
	level := valueobject.NewTypeLevelFromString(m.Level)
	data := &profile.EducationDTO{
		Id:             m.Id,
		Email:          m.Email,
		Level:          level,
		Name:           m.Name,
		Major:          m.Major,
		StillEducation: m.StillEducation,
		DateRange: profile.DateRangeEduDTO{
			Start: m.StartEdu,
			End:   m.EndEdu,
		},
		Description: m.Description,
	}
	return data
}

func EntityEducationToModel(m *profile.Education) *profile2.EducationModel {
	data := &profile2.EducationModel{
		Id:             m.GetId(),
		Email:          m.GetEmail(),
		Level:          m.GetLevel().StringLevel(),
		Name:           m.GetName(),
		Major:          m.GetMajor(),
		StillEducation: m.GetStillEducation(),
		StartEdu:       m.GetStartEducation(),
		EndEdu:         m.GetEndEducation(),
		Description:    m.GetDescription(),
	}

	return data
}

func ModelListEducationToEntity(m []*profile2.EducationModel) []*profile.EducationDTO {
	listEducation := make([]*profile.EducationDTO, 0)
	for _, data := range m {
		education := ModelEducationToEntity(data)
		listEducation = append(listEducation, education)
	}
	return listEducation
}

func EntityEducationeToInterface(data *profile.Education) []interface{} {
	return dbq.Struct(EntityEducationToModel(data))
}

func DomainEducationToInterface(domain *profile.Education) (dbqStruct []interface{}) {
	dbqStruct = append(dbqStruct, EntityEducationeToInterface(domain))
	return
}

package response

import (
	"CareerCenter/domain/entity/profile"
	"CareerCenter/utils"
)

type ProfileResponse struct {
	Email          string                    `json:"email"`
	Name           string                    `json:"name"`
	Photo          string                    `json:"photo"`
	Skill          string                    `json:"skill"`
	PhoneNumber    string                    `json:"phoneNumber"`
	WorkExperience []*WorkExperienceResponse `json:"workExperience"`
	Education      []*EducationResponse      `json:"education"`
	Ability        []string                  `json:"ability"`
	Language       []string                  `json:"language"`
	CvResume       string                    `json:"cvResume"`
	Portofolio     string                    `json:"portofolio"`
	CreatedAt      string                    `json:"createdAt"`
	UpdateAt       string                    `json:"updateAt"`
}

func GetProfileResponse(dto *profile.ProfileUserDTO) *ProfileResponse {
	listWorkExperiencet := make([]*WorkExperienceResponse, 0)
	for _, data := range dto.WorkExperience {
		workExperience := &WorkExperienceResponse{
			Id:              data.Id,
			SkillExperience: data.SkillExperience,
			Name:            data.Name,
			StillWorking:    data.StillWorking,
			DateRange: DateRange{
				Start: utils.ToOnlyDateResponse(data.DateRange.Start),
				End:   utils.ToOnlyDateResponse(data.DateRange.End),
			},
			Description: data.Description,
		}
		listWorkExperiencet = append(listWorkExperiencet, workExperience)
	}

	listEducation := make([]*EducationResponse, 0)
	for _, data := range dto.Education {
		education := &EducationResponse{
			Id:             data.Id,
			Level:          data.Level.StringLevel(),
			Name:           data.Name,
			Major:          data.Major,
			SkillEducation: data.StillEducation,
			DateRange: DateRange{
				Start: utils.ToOnlyDateResponse(data.DateRange.Start),
				End:   utils.ToOnlyDateResponse(data.DateRange.End),
			},
			Description: data.Description,
		}
		listEducation = append(listEducation, education)
	}

	return &ProfileResponse{
		Email:          dto.Email,
		Name:           dto.Name,
		Photo:          dto.Photo,
		Skill:          dto.Skill,
		PhoneNumber:    dto.PhoneNumber,
		WorkExperience: listWorkExperiencet,
		Education:      listEducation,
		Ability:        dto.Ability,
		Language:       dto.Language,
		CvResume:       dto.CvResume,
		Portofolio:     dto.Portofolio,
		CreatedAt:      utils.ToOnlyDateResponse(dto.CreatedAt),
		UpdateAt:       utils.ToOnlyDateResponse(dto.UpdatedAt),
	}
}

type WorkExperienceResponse struct {
	Id              string
	SkillExperience string
	Name            string
	StillWorking    bool
	DateRange       DateRange
	Description     string
}

type EducationResponse struct {
	Id             string
	Level          string
	Name           string
	Major          string
	SkillEducation bool
	DateRange      DateRange
	Description    string
}

type DateRange struct {
	Start string
	End   string
}

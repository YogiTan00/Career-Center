package response

import (
	"CareerCenter/domain/entity/profile"
	"time"
)

type ProfileResponse struct {
	Email          string                     `json:"email"`
	Name           string                     `json:"name"`
	Photo          string                     `json:"photo"`
	Skill          string                     `json:"skill"`
	NoTlp          string                     `json:"noTlp"`
	WorkExperience *profile.WorkExperienceDTO `json:"workExperience"`
	Education      *profile.EducationDTO      `json:"education"`
	Ability        []string                   `json:"ability"`
	Language       []string                   `json:"language"`
	Cv             string                     `json:"cv"`
	Portofolio     string                     `json:"portofolio"`
	CreatedAt      time.Time                  `json:"createdAt"`
	UpdateAt       time.Time                  `json:"updateAt"`
}

func GetProfileResponse(dto *profile.ProfileUserDTO) *ProfileResponse {
	return &ProfileResponse{
		Email:          dto.Email,
		Name:           dto.Name,
		Photo:          dto.Photo,
		Skill:          dto.Skill,
		NoTlp:          dto.NoTlp,
		WorkExperience: dto.WorkExperience,
		Education:      dto.Education,
		Ability:        dto.Ability,
		Language:       dto.Language,
		Cv:             dto.Cv,
		Portofolio:     dto.Portofolio,
		CreatedAt:      dto.CreatedAt,
		UpdateAt:       dto.UpdateAt,
	}
}

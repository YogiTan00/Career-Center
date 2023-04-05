package response

import (
	"CareerCenter/domain/entity/profile"
	"time"
)

type ProfileResponse struct {
	Email          string                       `json:"email"`
	Name           string                       `json:"name"`
	Photo          string                       `json:"photo"`
	Skill          string                       `json:"skill"`
	PhoneNumber    string                       `json:"phoneNumber"`
	WorkExperience []*profile.WorkExperienceDTO `json:"workExperience"`
	Education      []*profile.EducationDTO      `json:"education"`
	Ability        []string                     `json:"ability"`
	Language       []string                     `json:"language"`
	CvResume       string                       `json:"cvResume"`
	Portofolio     string                       `json:"portofolio"`
	CreatedAt      time.Time                    `json:"createdAt"`
	UpdateAt       time.Time                    `json:"updateAt"`
}

func GetProfileResponse(dto *profile.ProfileUserDTO) *ProfileResponse {
	return &ProfileResponse{
		Email:          dto.Email,
		Name:           dto.Name,
		Photo:          dto.Photo,
		Skill:          dto.Skill,
		PhoneNumber:    dto.PhoneNumber,
		WorkExperience: dto.WorkExperience,
		Education:      dto.Education,
		Ability:        dto.Ability,
		Language:       dto.Language,
		CvResume:       dto.CvResume,
		Portofolio:     dto.Portofolio,
		CreatedAt:      dto.CreatedAt,
		UpdateAt:       dto.UpdatedAt,
	}
}

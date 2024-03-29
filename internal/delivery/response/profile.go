package response

import (
	"CareerCenter/domain/entity/profile"
	"CareerCenter/pkg/config"
)

type ProfileResponse struct {
	Email          string                   `json:"email"`
	Name           string                   `json:"name"`
	Photo          string                   `json:"photo"`
	UrlPhotos      string                   `json:"urlPhoto"`
	Skill          string                   `json:"skill"`
	PhoneNumber    string                   `json:"phoneNumber"`
	WorkExperience []WorkExperienceResponse `json:"workExperience"`
	Education      []EducationResponse      `json:"education"`
	Ability        []string                 `json:"ability"`
	Language       []string                 `json:"language"`
	CvResume       string                   `json:"cvResume"`
	UrlCvResume    string                   `json:"urlCvResume"`
	Portofolio     string                   `json:"portofolio"`
	UrlPortofolio  string                   `json:"urlPortofolio"`
	CreatedAt      string                   `json:"createdAt"`
	UpdateAt       string                   `json:"updateAt"`
}

func GetProfileResponse(dto *profile.ProfileUserDTO, cfg config.Config) *ProfileResponse {
	if dto != nil {
		listWorkExperiencet := make([]WorkExperienceResponse, 0)
		for _, data := range dto.WorkExperience {
			workExperience := WorkExperienceResponse{
				Id:              data.Id,
				SkillExperience: data.SkillExperience,
				Name:            data.Name,
				StillWorking:    data.StillWorking,
				DateRange: DateRange{
					Start: data.DateRange.Start.Format("2006-01-02"),
					End:   data.DateRange.End.Format("2006-01-02"),
				},
				Description: data.Description,
			}
			listWorkExperiencet = append(listWorkExperiencet, workExperience)
		}

		listEducation := make([]EducationResponse, 0)
		for _, data := range dto.Education {
			education := EducationResponse{
				Id:             data.Id,
				Level:          data.Level.StringLevel(),
				Name:           data.Name,
				Major:          data.Major,
				SkillEducation: data.StillEducation,
				DateRange: DateRange{
					Start: data.DateRange.Start.Format("2006-01-02"),
					End:   data.DateRange.End.Format("2006-01-02"),
				},
				Description: data.Description,
			}
			listEducation = append(listEducation, education)
		}

		var urlPhotos, urlCvResume, urlPortofolio string
		if dto.Photo != "" {
			urlPhotos = cfg.DOMAIN + cfg.PATH_IMAGE_UPLOAD_META + dto.Photo
		}
		if dto.CvResume != "" {
			urlCvResume = cfg.DOMAIN + cfg.PATH_FILE_UPLOAD_META + dto.CvResume
		}
		if dto.Portofolio != "" {
			urlPortofolio = cfg.DOMAIN + cfg.PATH_FILE_UPLOAD_META + dto.Portofolio
		}

		return &ProfileResponse{
			Email:          dto.Email,
			Name:           dto.Name,
			Photo:          dto.Photo,
			UrlPhotos:      urlPhotos,
			Skill:          dto.Skill,
			PhoneNumber:    dto.PhoneNumber,
			WorkExperience: listWorkExperiencet,
			Education:      listEducation,
			Ability:        dto.Ability,
			Language:       dto.Language,
			CvResume:       dto.CvResume,
			UrlCvResume:    urlCvResume,
			Portofolio:     dto.Portofolio,
			UrlPortofolio:  urlPortofolio,
			CreatedAt:      dto.CreatedAt.Format("2006-01-02"),
			UpdateAt:       dto.UpdatedAt.Format("2006-01-02"),
		}
	}

	return nil
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

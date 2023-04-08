package request

import (
	"CareerCenter/domain/entity/profile"
	"CareerCenter/utils"
)

type RequestUpdateProfile struct {
	Nama        string `json:"name"`
	Skill       string `json:"skill"`
	PhoneNumber string `json:"phoneNumber"`
}

func NewUpdateProfileRequest(req *RequestUpdateProfile) *profile.ProfileUserDTO {
	return &profile.ProfileUserDTO{
		Name:        req.Nama,
		Skill:       req.Skill,
		PhoneNumber: req.PhoneNumber,
	}
}

type RequestWorkExperience struct {
	SkillExperience string `json:"skillExperience"`
	Name            string `json:"name"`
	StillWorking    bool   `json:"stillWorking"`
	StartWork       string `json:"startWork"`
	EndWork         string `json:"endWork"`
	Description     string `json:"description"`
}

func NewUpdateWorkExperience(req *RequestWorkExperience) (*profile.WorkExperienceDTO, error) {
	startWork, err := utils.ToDate(req.StartWork)
	if err != nil {
		return nil, err
	}
	endWork, err := utils.ToDate(req.EndWork)
	if err != nil {
		return nil, err
	}
	return &profile.WorkExperienceDTO{
		SkillExperience: req.SkillExperience,
		Name:            req.Name,
		StillWorking:    req.StillWorking,
		DateRange: profile.DateRangeWorkDTO{
			Start: startWork,
			End:   endWork,
		},
		Description: req.Description,
	}, nil
}

type RequestEducation struct {
	Level          string `json:"level"`
	Name           string `json:"name"`
	Major          string `json:"major"`
	StillEducation bool   `json:"stillEducation"`
	StartEdu       string `json:"startEducation"`
	EndEdu         string `json:"endEducation"`
	Description    string `json:"description"`
}

func NewUpdateEducation(req *RequestEducation) (*profile.EducationDTO, error) {
	startEdu, err := utils.ToDate(req.StartEdu)
	if err != nil {
		return nil, err
	}
	endEdu, err := utils.ToDate(req.EndEdu)
	if err != nil {
		return nil, err
	}
	return &profile.EducationDTO{
		Level:          req.Level,
		Name:           req.Name,
		Major:          req.Major,
		StillEducation: req.StillEducation,
		DateRange: profile.DateRangeEduDTO{
			Start: startEdu,
			End:   endEdu,
		},
		Description: req.Description,
	}, nil
}

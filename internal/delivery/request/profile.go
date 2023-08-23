package request

import (
	"CareerCenter/domain/entity/profile"
	"CareerCenter/domain/valueobject"
	"CareerCenter/utils"
	"time"
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
	var (
		startWork time.Time
		endWork   time.Time
		err       error
	)
	if len(req.StartWork) > 0 {
		startWork, err = utils.ToDate(req.StartWork)
		if err != nil {
			return nil, err
		}
	}
	if len(req.EndWork) > 0 {
		endWork, err = utils.ToDate(req.EndWork)
		if err != nil {
			return nil, err
		}
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
	var (
		startEdu time.Time
		endEdu   time.Time
		err      error
	)
	if len(req.StartEdu) > 0 {
		startEdu, err = utils.ToDate(req.StartEdu)
		if err != nil {
			return nil, err
		}
	}
	if len(req.EndEdu) > 0 {
		endEdu, _ = utils.ToDate(req.EndEdu)
		if err != nil {
			return nil, err
		}
	}

	level := valueobject.NewTypeLevelFromString(req.Level)

	return &profile.EducationDTO{
		Level:          level,
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

package entity

import (
	"CareerCenter/domain/entity/profile"
)

type HomePage struct {
	profile *profile.ProfileUser
	jobs    []*Jobs
}

type HomePageDTO struct {
	Profile *profile.ProfileUserDTO
	Jobs    []*JobsDTO
}

func NewHomePage(dto *HomePageDTO) (*HomePage, error) {
	err := dto.Validation()
	if err != nil {
		return nil, err
	}

	listJob := make([]*Jobs, 0)
	for _, data := range dto.Jobs {
		jobs := &JobsDTO{
			Id:        data.Id,
			Position:  data.Position,
			Company:   data.Company,
			Logo:      data.Logo,
			Address:   data.Address,
			CreatedAt: data.CreatedAt,
		}
		resultJob, err := NewJobs(jobs)
		if err != nil {
			return nil, err
		}
		listJob = append(listJob, resultJob)
	}

	profile, errProfile := profile.NewProfile(dto.Profile)
	if errProfile != nil {
		return nil, errProfile
	}
	return &HomePage{
		profile: profile,
		jobs:    listJob,
	}, nil
}

func (dto *HomePageDTO) Validation() error {
	return nil
}

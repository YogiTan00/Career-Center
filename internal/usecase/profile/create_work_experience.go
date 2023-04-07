package profile

import (
	"CareerCenter/domain/entity/profile"
	"context"
)

func (u UseCaseProfileInteractor) CreateWorkExperiencet(ctx context.Context, email string, workExp *profile.WorkExperienceDTO) error {
	data, err := profile.NewWorkExperience(workExp)
	if err != nil {
		return err
	}

	err = u.repoProfile.CreateWorkExperience(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

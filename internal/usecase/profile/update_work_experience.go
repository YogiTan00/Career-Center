package profile

import (
	"CareerCenter/domain/entity/profile"
	"context"
)

func (u UseCaseProfileInteractor) UpdateWorkExperience(ctx context.Context, id string, workExp *profile.WorkExperienceDTO) error {
	payload, err := profile.NewWorkExperience(workExp)
	if err != nil {
		return err
	}

	err = u.repoProfile.UpdateWorkExperience(ctx, id, payload)
	if err != nil {
		return err
	}
	return nil
}

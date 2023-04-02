package profile

import (
	"CareerCenter/domain/entity/profile"
	"context"
)

func (u UseCaseProfileInteractor) UpdateWorkExperience(ctx context.Context, email string, workExp *profile.WorkExperienceDTO) error {
	payload := profile.NewWorkExperience(*workExp)

	err := u.repoProfile.UpdateWorkExperience(ctx, email, payload)
	if err != nil {
		return err
	}
	return nil
}

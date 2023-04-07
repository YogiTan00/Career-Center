package profile

import (
	"CareerCenter/domain/entity/profile"
	"context"
	"github.com/google/uuid"
)

func (u UseCaseProfileInteractor) UpdateWorkExperience(ctx context.Context, id string, workExp *profile.WorkExperienceDTO) error {
	check, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	payload, err := profile.NewWorkExperience(workExp)
	if err != nil {
		return err
	}

	err = u.repoProfile.UpdateWorkExperience(ctx, check.String(), payload)
	if err != nil {
		return err
	}
	return nil
}

package profile

import (
	"CareerCenter/domain/entity/profile"
	"context"

	"github.com/google/uuid"
)

func (u UseCaseProfileInteractor) UpdateEducation(ctx context.Context, id string, education *profile.EducationDTO) error {
	check, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	payload, err := profile.NewEducation(education)
	if err != nil {
		return err
	}

	err = u.repoProfile.UpdateEducation(ctx, check.String(), payload)
	if err != nil {
		return err
	}
	return nil
}

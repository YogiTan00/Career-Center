package profile

import (
	"context"

	"github.com/google/uuid"
)

func (u UseCaseProfileInteractor) DeletedWorkExperience(ctx context.Context, id string) error {
	check, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	err = u.repoProfile.DeletedWorkExperience(ctx, check.String())
	if err != nil {
		return err
	}
	return nil
}

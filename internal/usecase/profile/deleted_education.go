package profile

import (
	"context"

	"github.com/google/uuid"
)

func (u UseCaseProfileInteractor) DeletedEducation(ctx context.Context, id string) error {
	check, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	err = u.repoProfile.DeletedEducation(ctx, check.String())
	if err != nil {
		return err
	}
	return nil
}

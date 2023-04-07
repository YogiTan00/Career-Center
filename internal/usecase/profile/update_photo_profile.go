package profile

import (
	"context"
)

func (u UseCaseProfileInteractor) UpdatePhotoProfile(ctx context.Context, email string, path string) error {
	err := u.repoProfile.UpdateOneColoum(ctx, email, path)
	if err != nil {
		return err
	}
	return nil
}

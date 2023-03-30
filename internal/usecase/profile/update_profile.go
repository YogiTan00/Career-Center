package profile

import (
	"CareerCenter/domain/entity/profile"
	"context"
)

func (u UseCaseProfileInteractor) UpdateProfile(ctx context.Context, email string, data *profile.ProfileUserDTO) error {
	payload, err := profile.NewProfile(data)
	if err != nil {
		return err
	}
	err = u.repoProfile.UpdateProfile(ctx, email, payload)
	if err != nil {
		return err
	}
	return nil
}

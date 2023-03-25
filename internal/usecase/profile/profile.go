package profile

import (
	"CareerCenter/domain/entity/profile"
	"context"
)

func (u UseCaseProfileInteractor) GetProfileByEmail(ctx context.Context, email string) (*profile.ProfileUserDTO, error) {
	data, err := u.repoProfile.GetProfileByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return data, nil
}

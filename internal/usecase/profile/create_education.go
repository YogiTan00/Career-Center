package profile

import (
	"CareerCenter/domain/entity/profile"
	"context"
)

func (u UseCaseProfileInteractor) CreateEducation(ctx context.Context, email string, education *profile.EducationDTO) error {
	data, err := profile.NewEducation(education)
	if err != nil {
		return err
	}
	data.SetEmail(email)
	err = u.repoProfile.CreateEducation(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

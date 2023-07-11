package application

import (
	"CareerCenter/domain/entity"
	"context"
)

func (u UseCaseApplicationInteractor) SendApplication(ctx context.Context, email string, apply *entity.ApplicationRequest) error {
	data, err := u.repoProfile.GetProfileByEmail(ctx, email)
	if err != nil {
		return err
	}
	toApplication, err := entity.NewApplicationFromProfile(data, apply)
	if err != nil {
		return err
	}

	err = u.repoApplication.SendApplication(ctx, toApplication)
	if err != nil {
		return err
	}
	return nil
}

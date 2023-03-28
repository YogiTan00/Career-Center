package application

import (
	"CareerCenter/domain/entity"
	"context"
)

func (u UseCaseApplicationInteractor) SendApplication(ctx context.Context, email string, companyId string) error {
	data, err := u.repoProfile.GetProfileByEmail(ctx, email)
	if err != nil {
		return err
	}
	toApplication, err := entity.NewApplicationFromProfile(data, companyId)
	if err != nil {
		return err
	}

	err = u.repoApplication.SendApplication(ctx, toApplication)
	if err != nil {
		return err
	}
	return nil
}

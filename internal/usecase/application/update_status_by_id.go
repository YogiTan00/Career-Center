package application

import (
	"CareerCenter/domain/entity"
	"context"
)

func (u UseCaseApplicationInteractor) UpdateStatusById(ctx context.Context, applicant *entity.StatusApplicantRequest) error {
	err := u.repoApplication.UpdateStatusById(ctx, applicant)
	if err != nil {
		return err
	}
	return nil
}

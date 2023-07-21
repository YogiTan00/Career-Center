package usecase

import (
	"CareerCenter/domain/entity"
	"context"
)

type UseCasApplication interface {
	SendApplication(ctx context.Context, email string, apply *entity.ApplicationRequest) error
	GetApplicantByJobId(ctx context.Context, jobId string) ([]*entity.ApplicationDTO, error)
	UpdateStatusById(ctx context.Context, applicant *entity.StatusApplicantRequest) error
}

package repository

import (
	"CareerCenter/domain/entity"
	"context"
)

type RepoApplication interface {
	SendApplication(ctx context.Context, application *entity.Application) error
	GetByEmail(ctx context.Context, email string) (*entity.ApplicationDTO, error)
	GetListApplication(ctx context.Context) ([]*entity.ApplicationDTO, error)
	GetByJobId(ctx context.Context, id string) ([]*entity.ApplicationDTO, error)
	GetListByEmail(ctx context.Context, email string) ([]*entity.ApplicationDTO, error)
	UpdateStatusById(ctx context.Context, applicant *entity.StatusApplicantRequest) error
}

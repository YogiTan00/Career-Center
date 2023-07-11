package repository

import (
	"CareerCenter/domain/entity"
	"context"
)

type RepoApplication interface {
	SendApplication(ctx context.Context, application *entity.Application) error
	GetByEmail(ctx context.Context, email string, companyId string) (*entity.ApplicationDTO, error)
	GetListApplication(ctx context.Context) ([]*entity.ApplicationDTO, error)
}

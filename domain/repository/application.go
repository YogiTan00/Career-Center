package repository

import (
	"CareerCenter/domain/entity"
	"context"
)

type RepoApplication interface {
	SendApplication(ctx context.Context, application *entity.Application) error
}

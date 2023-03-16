package usecase

import (
	"CareerCenter/domain/entity"
	"context"
)

type UseCaseJobs interface {
	GetListJobs(ctx context.Context, f *entity.FilterDTO) ([]*entity.JobsDTO, error)
}

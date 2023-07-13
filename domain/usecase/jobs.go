package usecase

import (
	"CareerCenter/domain/entity"
	"CareerCenter/domain/entity/filter"
	"context"
)

type UseCaseJobs interface {
	GetListJobs(ctx context.Context, f *filter.FilterDTO) ([]*entity.JobsDTO, int, error)
	GetJobById(ctx context.Context, email string, id string) (*entity.JobsDTO, error)
	CreateJob(ctx context.Context, dto *entity.JobsDTO) error
}

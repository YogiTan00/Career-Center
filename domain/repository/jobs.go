package repository

import (
	"CareerCenter/domain/entity"
	"CareerCenter/domain/entity/filter"
	"CareerCenter/domain/valueobject"
	"context"
)

type RepoJobs interface {
	GetListJobs(ctx context.Context, typeSearch *valueobject.TypeSearch, f *filter.Filter) ([]*entity.JobsDTO, error)
	GetJobById(ctx context.Context, id string) (*entity.JobsDTO, error)
	GetJobByCompanyId(ctx context.Context, id string) ([]*entity.JobsDTO, error)
	CreateJob(ctx context.Context, dto *entity.Jobs) error
}

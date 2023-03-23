package repository

import (
	"CareerCenter/domain/entity"
	"CareerCenter/domain/entity/filter"
	"context"
)

type RepoJobs interface {
	GetListJobs(ctx context.Context, f *filter.Filter) ([]*entity.JobsDTO, error)
}

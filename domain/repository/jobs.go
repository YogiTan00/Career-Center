package repository

import (
	"CareerCenter/domain/entity"
	"context"
)

type RepoJobs interface {
	GetListJobs(ctx context.Context, filter *entity.Filter) ([]*entity.JobsDTO, error)
}

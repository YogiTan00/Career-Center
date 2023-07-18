package usecase

import (
	"CareerCenter/domain/entity"
	"CareerCenter/domain/entity/filter"
	"CareerCenter/utils"
	"context"
)

type UseCaseJobs interface {
	GetListJobs(ctx context.Context, f *filter.FilterDTO) ([]*entity.JobsDTO, int, error)
	GetJobById(ctx context.Context, user *utils.User, id string) (*entity.JobsDTO, error)
	CreateJob(ctx context.Context, dto *entity.JobsDTO) error
	UpdateJob(ctx context.Context, id string, dto *entity.JobsDTO) error
	DeleteJob(ctx context.Context, id string) error
	GetListByEmail(ctx context.Context, user *utils.User) ([]*entity.JobsDTO, int, error)
}

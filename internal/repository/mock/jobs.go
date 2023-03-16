package mock

import (
	"CareerCenter/domain/entity"
	"context"
	"github.com/stretchr/testify/mock"
)

type RepoJobs struct {
	mock.Mock
}

func (r RepoAccount) GetListJobs(ctx context.Context, filter *entity.Filter) ([]*entity.JobsDTO, error) {
	args := r.Called(ctx, filter)
	return args.Get(0).([]*entity.JobsDTO), args.Error(1)
}

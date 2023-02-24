package mock

import (
	"CareerCenter/domain/entity"
	"context"
	"github.com/stretchr/testify/mock"
)

type RepoAccount struct {
	mock.Mock
}

func (r *RepoAccount) CreateAccount(ctx context.Context, data *entity.Account) error {
	args := r.Called(ctx, data)
	return args.Error(0)
}

func (r *RepoAccount) GetByEmail(ctx context.Context, email string) (*entity.AccountDTO, error) {
	args := r.Called(ctx, email)
	return args.Get(0).(*entity.AccountDTO), args.Error(1)
}

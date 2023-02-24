package repository

import (
	"CareerCenter/domain/entity"
	"context"
)

type RepoAccount interface {
	CreateAccount(ctx context.Context, data *entity.Account) error
	GetByEmail(ctx context.Context, email string) (*entity.AccountDTO, error)
}

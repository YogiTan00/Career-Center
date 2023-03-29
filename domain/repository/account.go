package repository

import (
	"CareerCenter/domain/entity/account"
	"context"
)

type RepoAccount interface {
	CreateAccount(ctx context.Context, data *account.Account) error
	GetByEmail(ctx context.Context, email string) (*account.AccountDTO, error)
	UpdatePassword(ctx context.Context, email string, password string) error
}

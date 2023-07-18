package repository

import (
	"CareerCenter/domain/entity/account"
	"context"
)

type RepoAccount interface {
	CreateAccount(ctx context.Context, data *account.Account) error
	GetByEmail(ctx context.Context, email string) (*account.AccountDTO, error)
	UpdatePassword(ctx context.Context, email string, password string) error
	UpdateRole(ctx context.Context, data *account.Account) error
	UpdateOTP(ctx context.Context, email string, otp string) error
	RemoveOTP(ctx context.Context, email string) error
	GetOTP(ctx context.Context, email string) (*account.CodeOTP, error)
}

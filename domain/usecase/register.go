package usecase

import (
	"CareerCenter/domain/entity"
	"context"
)

type UseCaseAccount interface {
	Register(ctx context.Context, data *entity.AccountDTO) error
	Login(ctx context.Context, email string, password string) (*entity.AccountDTO, error)
}

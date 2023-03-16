package usecase

import (
	"CareerCenter/domain/entity"
	"context"
	"net/http"
)

type UseCaseAccount interface {
	Register(ctx context.Context, data *entity.AccountDTO) error
	Login(ctx context.Context, email string, password string) (*http.Cookie, error)
}

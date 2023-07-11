package usecase

import (
	"CareerCenter/domain/entity"
	"context"
)

type UseCasApplication interface {
	SendApplication(ctx context.Context, email string, apply *entity.ApplicationRequest) error
}

package usecase

import (
	"context"
)

type UseCasApplication interface {
	SendApplication(ctx context.Context, email string, companyId string) error
}

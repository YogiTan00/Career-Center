package usecase

import (
	"CareerCenter/domain/entity"
	"CareerCenter/domain/entity/filter"
	"context"
)

type UseCaseHome interface {
	GetHome(ctx context.Context, f *filter.FilterHomeDTO, email string) (*entity.HomePageDTO, error)
}

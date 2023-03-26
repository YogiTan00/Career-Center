package usecase

import (
	"CareerCenter/domain/entity"
	"CareerCenter/domain/entity/filter"
	"context"
)

type UseCaseCompany interface {
	GetListCompany(ctx context.Context, f *filter.FilterDTO) ([]*entity.CompanyDTO, error)
	GetCompanyById(ctx context.Context, id string) (*entity.CompanyDTO, []*entity.JobsDTO, error)
}

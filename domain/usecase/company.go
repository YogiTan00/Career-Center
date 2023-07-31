package usecase

import (
	"CareerCenter/domain/entity"
	"CareerCenter/domain/entity/filter"
	"context"
)

type UseCaseCompany interface {
	CreateCompany(ctx context.Context, dto *entity.CompanyDTO) error
	GetListCompany(ctx context.Context, f *filter.FilterDTO) ([]*entity.CompanyDTO, error)
	GetCompanyById(ctx context.Context, id string) (*entity.CompanyDTO, []*entity.JobsDTO, error)
	UpdateCompanyById(ctx context.Context, id string, dto *entity.CompanyDTO) error
	DeleteCompanyById(ctx context.Context, companyId string) error
	UpdateLogoCompany(ctx context.Context, companyId string, path string) error
}

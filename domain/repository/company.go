package repository

import (
	"CareerCenter/domain/entity"
	"CareerCenter/domain/entity/filter"
	"CareerCenter/domain/valueobject"
	"context"
)

type RepoCompany interface {
	CreateCompany(ctx context.Context, company *entity.Company) error
	GetListCompany(ctx context.Context, typeSearch *valueobject.TypeSearch, f *filter.Filter) ([]*entity.CompanyDTO, error)
	GetCompanyById(ctx context.Context, id string) (*entity.CompanyDTO, error)
	UpdateCompanyById(ctx context.Context, companyId string, company *entity.Company) error
	DeleteCompanyById(ctx context.Context, companyId string) error
	UpdateLogoCompanyById(ctx context.Context, companyId string, path string) error
}

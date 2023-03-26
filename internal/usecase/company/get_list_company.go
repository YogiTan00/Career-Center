package company

import (
	"CareerCenter/domain/entity"
	"CareerCenter/domain/entity/filter"
	"CareerCenter/domain/valueobject"
	"context"
)

func (u UseCaseCompanyInteractor) GetListCompany(ctx context.Context, f *filter.FilterDTO) ([]*entity.CompanyDTO, error) {
	filter := filter.NewFilter(f)

	typeSearch := valueobject.NewTypeSearch(valueobject.COMPANY)
	data, err := u.repoCompany.GetListCompany(ctx, typeSearch, filter)
	if err != nil {
		return nil, err
	}

	return data, nil
}

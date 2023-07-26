package company

import (
	"CareerCenter/domain/entity"
	"CareerCenter/utils/exceptions"
	"context"
)

func (u UseCaseCompanyInteractor) UpdateCompanyById(ctx context.Context, id string, dto *entity.CompanyDTO) error {

	company, err := u.repoCompany.GetCompanyById(ctx, id)
	if err != nil {
		return err
	}

	if company == nil {
		return exceptions.ErrCantUpdate
	}

	err = u.repoCompany.UpdateCompanyById(ctx, id, company.SetUpdateCompany(dto))
	if err != nil {
		return err
	}

	return nil
}

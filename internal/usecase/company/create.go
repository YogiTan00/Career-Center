package company

import (
	"CareerCenter/domain/entity"
	"CareerCenter/utils/exceptions"
	"context"
)

func (u UseCaseCompanyInteractor) CreateCompany(ctx context.Context, dto *entity.CompanyDTO) error {
	check, err := u.repoProfile.GetProfileByEmail(ctx, dto.Email)
	if err != nil {
		return err
	}

	if check == nil {
		return exceptions.ErrCustomString("user not found")
	}

	dto.SetLogo(check)
	company, err := entity.NewCompany(dto)
	if err != nil {
		return err
	}

	err = u.repoCompany.CreateCompany(ctx, company)
	if err != nil {
		return err
	}

	return nil
}

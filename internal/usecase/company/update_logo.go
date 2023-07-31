package company

import "context"

func (u UseCaseCompanyInteractor) UpdateLogoCompany(ctx context.Context, companyId string, path string) error {
	err := u.repoCompany.UpdateLogoCompanyById(ctx, companyId, path)
	if err != nil {
		return err
	}

	return nil
}

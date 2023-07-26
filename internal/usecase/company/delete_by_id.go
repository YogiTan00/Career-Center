package company

import "context"

func (u UseCaseCompanyInteractor) DeleteCompanyById(ctx context.Context, companyId string) error {
	err := u.repoCompany.DeleteCompanyById(ctx, companyId)
	if err != nil {
		return err
	}
	return nil
}

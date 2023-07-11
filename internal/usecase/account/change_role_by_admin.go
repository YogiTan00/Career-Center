package account

import (
	"CareerCenter/domain/entity/account"
	"context"
)

func (u UseCaseAccountInteractor) ChangeRoleByAdmin(ctx context.Context, data *account.AccountDTO) error {
	dataRole, err := account.NewRole(data)
	if err != nil {
		return err
	}
	err = u.repoAccount.UpdateRole(ctx, dataRole)
	if err != nil {
		return err
	}
	return nil
}

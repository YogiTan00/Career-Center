package usecase

import (
	"CareerCenter/domain/entity"
	"CareerCenter/utils"
	"context"
)

func (u UseCaseAccountInteractor) Register(ctx context.Context, data *entity.AccountDTO) error {
	if data.Password != "" {
		password, _ := utils.HashPassword(data.Password)
		data.Password = password
	}
	register, errRegister := entity.NewAccount(data)
	if errRegister != nil {
		return errRegister
	}
	err := u.repoAccount.CreateAccount(ctx, register)
	if err != nil {
		return err
	}

	return nil
}

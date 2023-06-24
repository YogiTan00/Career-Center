package account

import (
	"CareerCenter/domain/entity/account"
	"CareerCenter/domain/entity/profile"
	"CareerCenter/utils"
	"context"
)

func (u UseCaseAccountInteractor) RegisterAdmin(ctx context.Context, email string, data *account.AccountDTO) error {
	var (
		err error
	)
	if data.Password != "" {
		password, err := utils.HashPassword(data.Password)
		if err != nil {
			return err
		}
		data.Password = password
	}

	register, errRegister := account.NewAccount(data)
	if errRegister != nil {
		return errRegister
	}

	register.SetRegisterBy(email)

	err = u.repoAccount.CreateAccount(ctx, register)
	if err != nil {
		return err
	}

	dataProfile, err := profile.NewProfileByRegister(register)
	if err != nil {
		return err
	}

	err = u.repoProfile.CreateProfile(ctx, dataProfile)
	if err != nil {
		return err
	}

	return nil
}

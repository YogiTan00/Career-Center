package account

import (
	"CareerCenter/domain/entity/account"
	"CareerCenter/domain/entity/profile"
	"CareerCenter/utils"
	"context"
	"time"
)

func (u UseCaseAccountInteractor) Register(ctx context.Context, data *account.AccountDTO) error {
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
	timeNow := time.Now()
	data.CreatedAt = timeNow
	data.UpdatedAt = timeNow
	register, errRegister := account.NewAccount(data)
	if errRegister != nil {
		return errRegister
	}
	err = u.repoAccount.CreateAccount(ctx, register)
	if err != nil {
		return err
	}
	create := &profile.ProfileUserDTO{
		Name:      data.Name,
		Email:     data.Email,
		CreatedAt: timeNow,
		UpdatedAt: timeNow,
	}

	dataProfile, err := profile.NewProfile(create)
	if err != nil {
		return err
	}
	err = u.repoProfile.CreateProfile(ctx, dataProfile)
	if err != nil {
		return err
	}

	return nil
}

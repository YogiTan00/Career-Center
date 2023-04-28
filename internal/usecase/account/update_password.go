package account

import (
	"CareerCenter/domain/entity/account"
	"CareerCenter/utils"
	"context"
	"errors"
)

func (u UseCaseAccountInteractor) UpdatePassword(ctx context.Context, email string, password *account.UpdatePasswordDTO) error {
	data, err := u.repoAccount.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	checkPw := utils.CheckPasswordHash(password.OldPassword, data.Password)
	if checkPw != true {
		return errors.New("wrong password")
	}

	checkSamePw := utils.CheckPasswordHash(password.NewPassword, data.Password)
	if checkSamePw != false {
		return errors.New("cannot be changed with the same password")
	}

	payload, err := account.NewPassword(password)
	if err != nil {
		return err
	}

	var newPassword string
	if payload.GetNewPassword() != "" {
		newPassword, err = utils.HashPassword(payload.GetNewPassword())
		if err != nil {
			return err
		}
	}

	err = u.repoAccount.UpdatePassword(ctx, email, newPassword)
	if err != nil {
		return err
	}

	return nil
}

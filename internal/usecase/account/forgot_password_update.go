package account

import (
	"CareerCenter/internal/delivery/request"
	"CareerCenter/utils"
	"context"
	"errors"
)

func (u UseCaseAccountInteractor) ForgetPasswordUpdate(ctx context.Context, body *request.RequestForgetPasswordUpdate) error {
	// check password match or no
	if body.NewPassword != body.ConfirmPassword {
		return errors.New("the password does not match")
	}
	// get user by mail
	user, err := u.repoAccount.GetByEmail(ctx, body.Email)
	if err != nil {
		return err
	}
	//
	var newPassword string
	if body.NewPassword != "" {
		newPassword, err = utils.HashPassword(body.NewPassword)
		if err != nil {
			return err
		}
	}
	// update password
	errOtpUpdate := u.repoAccount.UpdatePassword(ctx, user.Email, newPassword)
	if errOtpUpdate != nil {
		return errOtpUpdate
	}

	return nil
}

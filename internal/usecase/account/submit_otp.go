package account

import (
	"context"
	"errors"
	"time"
)

func (u UseCaseAccountInteractor) SubmitOtp(ctx context.Context, email string, otp string) error {
	// get user by mail
	user, err := u.repoAccount.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	// check otp expired or no
	if formattedTime > user.ExpiredOtp.Format("2006-01-02 15:04:05") {
		return errors.New("The verification code expired")
	}

	// check otp match or no
	if user.CodeOtp != otp {
		return errors.New("The verification code does not match")
	}

	// remove otp
	errRemove := u.repoAccount.RemoveOTP(ctx, email)
	if errRemove != nil {
		return errRemove
	}
	return nil
}

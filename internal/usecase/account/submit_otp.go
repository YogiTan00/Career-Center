package account

import (
	"context"
	"errors"
	"time"
)

func (u UseCaseAccountInteractor) SubmitOtp(ctx context.Context, email string, otp string) error {
	// get user by mail
	user, err := u.repoAccount.GetOTP(ctx, email)
	if err != nil {
		return err
	}

	currentTime := time.Now()
	// check otp match or no
	if user.Code != otp {
		return errors.New("the verification code does not match")
	}

	// check otp expired or no
	if currentTime.Unix() > user.Expired.Unix() {
		return errors.New("the verification code expired")
	}

	// remove otp
	errRemove := u.repoAccount.RemoveOTP(ctx, email)
	if errRemove != nil {
		return errRemove
	}
	return nil
}

package usecase

import (
	"CareerCenter/utils"
	"context"
	"errors"
)

func (u UseCaseAccountInteractor) Login(ctx context.Context, email string, password string) error {
	data, err := u.repoAccount.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	checkPw := utils.CheckPasswordHash(password, data.Password)
	if checkPw != true {
		return errors.New("wrong email or password")
	}
	return nil
}

package usecase

import (
	"CareerCenter/domain/entity"
	"CareerCenter/utils"
	"context"
	"errors"
)

func (u UseCaseAccountInteractor) Login(ctx context.Context, email string, password string) (*entity.AccountDTO, error) {
	data, err := u.repoAccount.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if data.Password != "" {
		errPw := utils.CheckPasswordHash(password, data.Password)
		if errPw != true {
			return nil, errors.New("error email or password")
		} else {

		}

	}
	return data, nil
}

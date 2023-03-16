package account

import (
	"CareerCenter/utils"
	"context"
	"errors"
	"net/http"
)

func (r UseCaseAccountInteractor) Login(ctx context.Context, email string, password string) (*http.Cookie, error) {
	data, err := r.repoAccount.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	checkPw := utils.CheckPasswordHash(password, data.Password)
	if checkPw != true {
		return nil, errors.New("wrong email or password")
	}

	token, errToken := utils.GenerateToken(email)
	if errToken != nil {
		return nil, errToken
	}
	return token, nil
}

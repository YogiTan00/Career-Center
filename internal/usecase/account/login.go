package account

import (
	"CareerCenter/utils"
	"CareerCenter/utils/exceptions"
	"context"
	"net/http"
)

func (r UseCaseAccountInteractor) Login(ctx context.Context, email string, password string) (*http.Cookie, error) {
	data, err := r.repoAccount.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	checkPw := utils.CheckPasswordHash(password, data.Password)
	if checkPw != true {
		return nil, exceptions.ErrorWrongEmailorPassword
	}

	token, errToken := utils.GenerateToken(email)
	if errToken != nil {
		return nil, errToken
	}
	return token, nil
}

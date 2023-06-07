package account

import (
	"CareerCenter/domain/entity/account"
	"CareerCenter/utils"
	"CareerCenter/utils/exceptions"
	"context"
	"net/http"
)

func (r UseCaseAccountInteractor) Login(ctx context.Context, data *account.AccountDTO) (*http.Cookie, *account.Login, error) {
	ac, err := r.repoAccount.GetByEmail(ctx, data.Email)
	if err != nil {
		return nil, nil, err
	}

	checkPw := utils.CheckPasswordHash(data.Password, ac.Password)
	if checkPw != true {
		return nil, nil, exceptions.ErrorWrongEmailorPassword
	}

	cookie, errToken := utils.GenerateToken(data.Email, ac.Role.StringRoles())
	if errToken != nil {
		return nil, nil, errToken
	}

	result := &account.Login{}
	result.SetLogin(cookie.Value, ac.Role.StringRoles())

	return cookie, result, nil
}

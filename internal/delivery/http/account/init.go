package account

import "CareerCenter/domain/usecase"

type AccountHandler struct {
	UCAccount usecase.UseCaseAccount
}

func NewUseCaseAccountHandler(UCAccount usecase.UseCaseAccount) *AccountHandler {
	return &AccountHandler{UCAccount: UCAccount}
}

package http

import "CareerCenter/domain/usecase"

type RegisterHandler struct {
	UCRegister usecase.UseCaseAccount
}

func NewUseCaseHandler(UCRegister usecase.UseCaseAccount) *RegisterHandler {
	return &RegisterHandler{UCRegister: UCRegister}
}

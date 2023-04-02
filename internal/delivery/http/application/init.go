package application

import "CareerCenter/domain/usecase"

type ApplicationHandler struct {
	UCApplication usecase.UseCasApplication
}

func NewUseCaseApplicationHandler(UCApplication usecase.UseCasApplication) *ApplicationHandler {
	return &ApplicationHandler{UCApplication: UCApplication}
}

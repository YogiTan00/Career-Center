package application

import (
	"CareerCenter/domain/usecase"
	"CareerCenter/pkg/config"
)

type ApplicationHandler struct {
	UCApplication usecase.UseCasApplication
	cfg           config.Config
}

func NewUseCaseApplicationHandler(UCApplication usecase.UseCasApplication, cfg config.Config) *ApplicationHandler {
	return &ApplicationHandler{
		UCApplication: UCApplication,
		cfg:           cfg,
	}
}

package profile

import (
	"CareerCenter/domain/usecase"
	"CareerCenter/pkg/config"
)

type ProfileHandler struct {
	UCProfile usecase.UseCaseProfile
	cfg       config.Config
}

func NewUseCaseProfileHandler(UCProfile usecase.UseCaseProfile, cfg config.Config) *ProfileHandler {
	return &ProfileHandler{UCProfile: UCProfile, cfg: cfg}
}

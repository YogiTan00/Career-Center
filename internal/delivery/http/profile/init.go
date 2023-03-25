package profile

import "CareerCenter/domain/usecase"

type ProfileHandler struct {
	UCProfile usecase.UseCaseProfile
}

func NewUseCaseProfileHandler(UCProfile usecase.UseCaseProfile) *ProfileHandler {
	return &ProfileHandler{UCProfile: UCProfile}
}

package home

import "CareerCenter/domain/usecase"

type HomeHandler struct {
	UCHome usecase.UseCaseHome
}

func NewUseCaseJobsHandler(UCHome usecase.UseCaseHome) *HomeHandler {
	return &HomeHandler{UCHome: UCHome}
}

package company

import (
	"CareerCenter/domain/usecase"
	"CareerCenter/pkg/config"
)

type CompanyHandler struct {
	UCCompany usecase.UseCaseCompany
	cfg       config.Config
}

func NewUseCaseCompanyHandler(UCCompany usecase.UseCaseCompany, cfg config.Config) *CompanyHandler {
	return &CompanyHandler{
		UCCompany: UCCompany,
		cfg:       cfg,
	}
}

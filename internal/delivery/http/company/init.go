package company

import "CareerCenter/domain/usecase"

type CompanyHandler struct {
	UCCompany usecase.UseCaseCompany
}

func NewUseCaseCompanyHandler(UCCompany usecase.UseCaseCompany) *CompanyHandler {
	return &CompanyHandler{UCCompany: UCCompany}
}

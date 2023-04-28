package company

import (
	"CareerCenter/domain/repository"
	"CareerCenter/domain/usecase"
)

type UseCaseCompanyInteractor struct {
	repoCompany repository.RepoCompany
	repoJobs    repository.RepoJobs
}

func NewCompanyUsecase(repoCompany repository.RepoCompany, repoJobs repository.RepoJobs) usecase.UseCaseCompany {
	return &UseCaseCompanyInteractor{
		repoCompany: repoCompany,
		repoJobs:    repoJobs,
	}
}

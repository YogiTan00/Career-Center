package company

import (
	"CareerCenter/domain/repository"
)

type UseCaseCompanyInteractor struct {
	repoCompany repository.RepoCompany
	repoJobs    repository.RepoJobs
}

func NewCompanyUsecase(repoCompany repository.RepoCompany, repoJobs repository.RepoJobs) *UseCaseCompanyInteractor {
	return &UseCaseCompanyInteractor{
		repoCompany: repoCompany,
		repoJobs:    repoJobs,
	}
}

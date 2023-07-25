package company

import (
	"CareerCenter/domain/repository"
	"CareerCenter/domain/usecase"
)

type UseCaseCompanyInteractor struct {
	repoCompany repository.RepoCompany
	repoJobs    repository.RepoJobs
	repoProfile repository.RepoProfile
}

func NewCompanyUsecase(repoCompany repository.RepoCompany, repoJobs repository.RepoJobs, repoProfile repository.RepoProfile) usecase.UseCaseCompany {
	return &UseCaseCompanyInteractor{
		repoCompany: repoCompany,
		repoJobs:    repoJobs,
		repoProfile: repoProfile,
	}
}

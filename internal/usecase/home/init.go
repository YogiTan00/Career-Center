package home

import "CareerCenter/domain/repository"

type UseCaseHomeInteractor struct {
	repoAccount repository.RepoAccount
	repoJobs    repository.RepoJobs
}

func NewHomeUseCase(repoAccount repository.RepoAccount, repoJobs repository.RepoJobs) *UseCaseHomeInteractor {
	return &UseCaseHomeInteractor{
		repoAccount: repoAccount,
		repoJobs:    repoJobs,
	}
}

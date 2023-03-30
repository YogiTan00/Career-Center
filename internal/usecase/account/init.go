package account

import (
	"CareerCenter/domain/repository"
)

type UseCaseAccountInteractor struct {
	repoAccount repository.RepoAccount
	repoProfile repository.RepoProfile
}

func NewAccountUsecase(repoAccount repository.RepoAccount, repoProfile repository.RepoProfile) *UseCaseAccountInteractor {
	return &UseCaseAccountInteractor{
		repoAccount: repoAccount,
		repoProfile: repoProfile,
	}
}

package account

import (
	"CareerCenter/domain/repository"
	"CareerCenter/domain/usecase"
)

type UseCaseAccountInteractor struct {
	repoAccount repository.RepoAccount
	repoProfile repository.RepoProfile
}

func NewAccountUsecase(repoAccount repository.RepoAccount, repoProfile repository.RepoProfile) usecase.UseCaseAccount {
	return &UseCaseAccountInteractor{
		repoAccount: repoAccount,
		repoProfile: repoProfile,
	}
}

package account

import (
	"CareerCenter/domain/repository"
)

type UseCaseAccountInteractor struct {
	repoAccount repository.RepoAccount
}

func NewAccountUsecase(repoAccount repository.RepoAccount) *UseCaseAccountInteractor {
	return &UseCaseAccountInteractor{
		repoAccount: repoAccount,
	}
}

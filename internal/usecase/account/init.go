package account

import (
	"CareerCenter/domain/repository"
	"CareerCenter/domain/usecase"
	"CareerCenter/pkg/config"
)

type UseCaseAccountInteractor struct {
	repoAccount repository.RepoAccount
	repoProfile repository.RepoProfile
	cfg         config.Config
}

func NewAccountUsecase(repoAccount repository.RepoAccount, repoProfile repository.RepoProfile, cfg config.Config) usecase.UseCaseAccount {
	return &UseCaseAccountInteractor{
		repoAccount: repoAccount,
		repoProfile: repoProfile,
		cfg:         cfg,
	}
}

package profile

import (
	"CareerCenter/domain/repository"
	"CareerCenter/domain/usecase"
	"CareerCenter/pkg/config"
)

type UseCaseProfileInteractor struct {
	repoProfile repository.RepoProfile
	cfg         config.Config
}

func NewProfileUsecase(repoProfile repository.RepoProfile, cfg config.Config) usecase.UseCaseProfile {
	return &UseCaseProfileInteractor{
		repoProfile: repoProfile,
		cfg:         cfg,
	}
}

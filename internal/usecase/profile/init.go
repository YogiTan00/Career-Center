package profile

import (
	"CareerCenter/domain/repository"
	"CareerCenter/domain/usecase"
)

type UseCaseProfileInteractor struct {
	repoProfile repository.RepoProfile
}

func NewProfileUsecase(repoProfile repository.RepoProfile) usecase.UseCaseProfile {
	return &UseCaseProfileInteractor{
		repoProfile: repoProfile,
	}
}

package application

import (
	"CareerCenter/domain/repository"
	"CareerCenter/domain/usecase"
)

type UseCaseApplicationInteractor struct {
	repoApplication repository.RepoApplication
	repoProfile     repository.RepoProfile
}

func NewApplicationUsecase(repoApplication repository.RepoApplication, repoProfile repository.RepoProfile) usecase.UseCasApplication {
	return &UseCaseApplicationInteractor{
		repoApplication: repoApplication,
		repoProfile:     repoProfile,
	}
}

package application

import (
	"CareerCenter/domain/repository"
)

type UseCaseApplicationInteractor struct {
	repoApplication repository.RepoApplication
	repoProfile     repository.RepoProfile
}

func NewApplicationUsecase(repoApplication repository.RepoApplication, repoProfile repository.RepoProfile) *UseCaseApplicationInteractor {
	return &UseCaseApplicationInteractor{
		repoApplication: repoApplication,
		repoProfile:     repoProfile,
	}
}

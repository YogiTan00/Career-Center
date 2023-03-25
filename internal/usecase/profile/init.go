package profile

import "CareerCenter/domain/repository"

type UseCaseProfileInteractor struct {
	repoProfile repository.RepoProfile
}

func NewProfileUsecase(repoProfile repository.RepoProfile) *UseCaseProfileInteractor {
	return &UseCaseProfileInteractor{
		repoProfile: repoProfile,
	}
}

package jobs

import (
	"CareerCenter/domain/repository"
	"CareerCenter/domain/usecase"
)

type UseCaseJobsInteractor struct {
	repoJobs        repository.RepoJobs
	repoApplication repository.RepoApplication
}

func NewJobsUsecase(repoJobs repository.RepoJobs, repoApplication repository.RepoApplication) usecase.UseCaseJobs {
	return &UseCaseJobsInteractor{
		repoJobs:        repoJobs,
		repoApplication: repoApplication,
	}
}

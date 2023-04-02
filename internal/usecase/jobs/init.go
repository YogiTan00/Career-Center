package jobs

import "CareerCenter/domain/repository"

type UseCaseJobsInteractor struct {
	repoJobs        repository.RepoJobs
	repoApplication repository.RepoApplication
}

func NewJobsUsecase(repoJobs repository.RepoJobs, repoApplication repository.RepoApplication) *UseCaseJobsInteractor {
	return &UseCaseJobsInteractor{
		repoJobs:        repoJobs,
		repoApplication: repoApplication,
	}
}

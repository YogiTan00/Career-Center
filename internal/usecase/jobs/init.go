package jobs

import "CareerCenter/domain/repository"

type UseCaseJobsInteractor struct {
	repoJobs repository.RepoJobs
}

func NewJobsUsecase(repoJobs repository.RepoJobs) *UseCaseJobsInteractor {
	return &UseCaseJobsInteractor{
		repoJobs: repoJobs,
	}
}

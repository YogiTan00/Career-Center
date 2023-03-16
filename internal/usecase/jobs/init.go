package jobs

import "CareerCenter/domain/repository"

type UseCaseJobstInteractor struct {
	repoJobs repository.RepoJobs
}

func NewJobsUsecase(repoJobs repository.RepoJobs) *UseCaseJobstInteractor {
	return &UseCaseJobstInteractor{
		repoJobs: repoJobs,
	}
}

package jobs

import "CareerCenter/domain/usecase"

type JobsHandler struct {
	UCJobs usecase.UseCaseJobs
}

func NewUseCaseJobsHandler(UCJobs usecase.UseCaseJobs) *JobsHandler {
	return &JobsHandler{UCJobs: UCJobs}
}

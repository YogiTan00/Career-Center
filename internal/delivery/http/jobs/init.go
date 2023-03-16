package jobs

import "CareerCenter/domain/usecase"

type JobsrHandler struct {
	UCJobs usecase.UseCaseJobs
}

func NewUseCaseJobsHandler(UCJobs usecase.UseCaseJobs) *JobsrHandler {
	return &JobsrHandler{UCJobs: UCJobs}
}

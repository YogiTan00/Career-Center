package jobs

import (
	"CareerCenter/domain/usecase"
	"CareerCenter/pkg/config"
)

type JobsHandler struct {
	UCJobs usecase.UseCaseJobs
	cfg    config.Config
}

func NewUseCaseJobsHandler(UCJobs usecase.UseCaseJobs, cfg config.Config) *JobsHandler {
	return &JobsHandler{
		UCJobs: UCJobs,
		cfg:    cfg,
	}
}

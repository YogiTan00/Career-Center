package jobs

import (
	"CareerCenter/domain/entity"
	"context"
)

func (u UseCaseJobsInteractor) CreateJob(ctx context.Context, dto *entity.JobsDTO) error {
	data, err := entity.NewJobs(dto)
	if err != nil {
		return err
	}

	err = u.repoJobs.CreateJob(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

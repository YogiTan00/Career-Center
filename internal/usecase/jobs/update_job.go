package jobs

import (
	"CareerCenter/domain/entity"
	"CareerCenter/utils/exceptions"
	"context"
)

func (u UseCaseJobsInteractor) UpdateJob(ctx context.Context, id string, dto *entity.JobsDTO) error {

	job, err := u.repoJobs.GetJobById(ctx, id)
	if err != nil {
		return err
	}
	if job == nil {
		return exceptions.ErrCustomString("cant update because data not found")
	}

	err = u.repoJobs.UpdateJobById(ctx, job.SetUpdate(dto))
	if err != nil {
		return err
	}

	return nil
}

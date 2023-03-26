package jobs

import (
	"CareerCenter/domain/entity"
	"CareerCenter/utils"
	"context"
)

func (u UseCaseJobsInteractor) GetJobById(ctx context.Context, id string) (*entity.JobsDTO, error) {
	uuid, err := utils.ValitUuId(id)
	if err != nil {
		return nil, err
	}
	data, err := u.repoJobs.GetJobById(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return data, nil
}

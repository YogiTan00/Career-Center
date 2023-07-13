package jobs

import (
	"CareerCenter/domain/entity"
	"CareerCenter/utils"
	"context"
)

func (u UseCaseJobsInteractor) GetJobById(ctx context.Context, email string, id string) (*entity.JobsDTO, error) {
	uuid, err := utils.ValitUuId(id)
	if err != nil {
		return nil, err
	}
	data, err := u.repoJobs.GetJobById(ctx, uuid)
	if err != nil {
		return nil, err
	}

	check, err := u.repoApplication.GetByEmail(ctx, email, id)
	if err != nil {
		return nil, err
	}

	if check != nil {
		data.Status = true
	}

	return data, nil
}

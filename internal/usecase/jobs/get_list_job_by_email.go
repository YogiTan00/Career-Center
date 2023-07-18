package jobs

import (
	"CareerCenter/domain/entity"
	"CareerCenter/utils"
	"context"
)

func (u UseCaseJobsInteractor) GetListByEmail(ctx context.Context, user *utils.User) ([]*entity.JobsDTO, int, error) {
	var result []*entity.JobsDTO
	data, err := u.repoApplication.GetListByEmail(ctx, user.Email)
	if err != nil {
		return nil, 0, err
	}
	if data != nil {
		for _, value := range data {
			job, _ := u.repoJobs.GetJobById(ctx, value.JobId)
			if job != nil {
				job.Status = true
				job.ApplyDate = value.CreatedAt
				result = append(result, job)
			}
		}
	}

	return result, len(result), nil
}

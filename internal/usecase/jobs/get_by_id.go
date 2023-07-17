package jobs

import (
	"CareerCenter/domain/entity"
	"CareerCenter/utils"
	"context"
)

func (u UseCaseJobsInteractor) GetJobById(ctx context.Context, user *utils.User, id string) (*entity.JobsDTO, error) {

	uuid, err := utils.ValitUuId(id)
	if err != nil {
		return nil, err
	}
	data, err := u.repoJobs.GetJobById(ctx, uuid)
	if err != nil {
		return nil, err
	}
	if user.Admin() {
		apply, _ := u.repoApplication.GetByJobId(ctx, uuid)
		data.Applicant = len(apply)
	} else {
		apply, _ := u.repoApplication.GetByEmail(ctx, user.Email)
		if apply != nil {
			data.SetForUser(apply)
		}
	}

	return data, nil
}

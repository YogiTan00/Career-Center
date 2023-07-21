package application

import (
	"CareerCenter/domain/entity"
	"CareerCenter/utils"
	"context"
)

func (u UseCaseApplicationInteractor) GetApplicantByJobId(ctx context.Context, jobId string) ([]*entity.ApplicationDTO, error) {
	uuid, err := utils.ValitUuId(jobId)
	if err != nil {
		return nil, err
	}

	applicant, err := u.repoApplication.GetByJobId(ctx, uuid)
	if err != nil {
		return nil, err
	}
	return applicant, nil
}

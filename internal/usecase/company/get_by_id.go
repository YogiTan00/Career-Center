package company

import (
	"CareerCenter/domain/entity"
	"CareerCenter/utils"
	"context"
)

func (u UseCaseCompanyInteractor) GetCompanyById(ctx context.Context, id string) (*entity.CompanyDTO, []*entity.JobsDTO, error) {
	uuid, err := utils.ValitUuId(id)
	if err != nil {
		return nil, nil, err
	}
	data, err := u.repoCompany.GetCompanyById(ctx, uuid)
	if err != nil {
		return nil, nil, err
	}

	dataJobs, err := u.repoJobs.GetJobByCompanyId(ctx, uuid)
	if err != nil {
		return nil, nil, err
	}

	return data, dataJobs, nil
}

package jobs

import (
	"CareerCenter/domain/entity"
	"context"
)

func (u UseCaseJobstInteractor) GetListJobs(ctx context.Context, f *entity.FilterDTO) ([]*entity.JobsDTO, error) {
	filter := entity.NewFilter(f)
	data, err := u.repoJobs.GetListJobs(ctx, filter)
	if err != nil {
		return nil, err
	}

	return data, nil
}

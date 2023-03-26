package jobs

import (
	"CareerCenter/domain/entity"
	"CareerCenter/domain/entity/filter"
	"context"
)

func (u UseCaseJobstInteractor) GetListJobs(ctx context.Context, f *filter.FilterDTO) ([]*entity.JobsDTO, error) {
	filter := filter.NewFilter(f)

	data, err := u.repoJobs.GetListJobs(ctx, filter)
	if err != nil {
		return nil, err
	}

	return data, nil
}

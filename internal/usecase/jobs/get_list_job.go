package jobs

import (
	"CareerCenter/domain/entity"
	"CareerCenter/domain/entity/filter"
	"CareerCenter/domain/valueobject"
	"context"
)

func (u UseCaseJobsInteractor) GetListJobs(ctx context.Context, f *filter.FilterDTO) ([]*entity.JobsDTO, error) {
	filter := filter.NewFilter(f)

	typeSearch := valueobject.NewTypeSearch(valueobject.JOBS)
	data, err := u.repoJobs.GetListJobs(ctx, typeSearch, filter)
	if err != nil {
		return nil, err
	}

	return data, nil
}

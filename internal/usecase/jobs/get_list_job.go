package jobs

import (
	"CareerCenter/domain/entity"
	"CareerCenter/domain/entity/filter"
	"CareerCenter/domain/valueobject"
	"context"
)

func (u UseCaseJobsInteractor) GetListJobs(ctx context.Context, f *filter.FilterDTO) ([]*entity.JobsDTO, int, error) {
	fil := filter.NewFilter(f)

	typeSearch := valueobject.NewTypeSearch(valueobject.JOBS)
	data, err := u.repoJobs.GetListJobs(ctx, typeSearch, fil)
	if err != nil {
		return nil, 0, err
	}

	applicant, err := u.repoApplication.GetListApplication(ctx)
	if err != nil {
		return nil, 0, err
	}

	count := 0
	for _, job := range data {
		for _, apply := range applicant {
			if job.Id == apply.JobId {
				count = count + 1
			}
		}
		job.Applicant = count
	}

	return data, len(data), nil
}

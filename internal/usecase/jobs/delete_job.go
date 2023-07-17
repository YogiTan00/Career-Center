package jobs

import "context"

func (u UseCaseJobsInteractor) DeleteJob(ctx context.Context, id string) error {
	err := u.repoJobs.DeleteJobById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

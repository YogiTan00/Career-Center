package profile

import "context"

func (u UseCaseProfileInteractor) UpdatePortofolio(ctx context.Context, email string, path string) error {
	err := u.repoProfile.UpdatePortofolio(ctx, email, path)
	if err != nil {
		return err
	}
	return nil
}

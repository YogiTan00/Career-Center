package profile

import "context"

func (u UseCaseProfileInteractor) UpdateCvResume(ctx context.Context, email string, path string) error {
	err := u.repoProfile.UpdateCvResume(ctx, email, path)
	if err != nil {
		return err
	}
	return nil
}

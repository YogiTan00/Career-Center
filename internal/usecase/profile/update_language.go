package profile

import (
	"context"
	"strings"
)

func (u UseCaseProfileInteractor) UpdateLanguage(ctx context.Context, email string, listLanguage []string) error {
	language := strings.Join(listLanguage, ",")

	err := u.repoProfile.UpdateLanguage(ctx, email, language)
	if err != nil {
		return err
	}
	return nil
}

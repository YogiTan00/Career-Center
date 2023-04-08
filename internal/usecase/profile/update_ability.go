package profile

import (
	"context"
	"strings"
)

func (u UseCaseProfileInteractor) UpdateAbility(ctx context.Context, email string, listAbility []string) error {
	ability := strings.Join(listAbility, ",")

	err := u.repoProfile.UpdateAbility(ctx, email, ability)
	if err != nil {
		return err
	}
	return nil
}

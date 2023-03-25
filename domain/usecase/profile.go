package usecase

import (
	"CareerCenter/domain/entity/profile"
	"context"
)

type UseCaseProfile interface {
	GetProfileByEmail(ctx context.Context, email string) (*profile.ProfileUserDTO, error)
}

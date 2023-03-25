package repository

import (
	"CareerCenter/domain/entity/profile"
	"context"
)

type RepoProfile interface {
	GetProfileByEmail(ctx context.Context, email string) (*profile.ProfileUserDTO, error)
}

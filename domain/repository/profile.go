package repository

import (
	"CareerCenter/domain/entity/profile"
	"context"
)

type RepoProfile interface {
	CreateProfile(ctx context.Context, data *profile.ProfileUser) error
	GetProfileByEmail(ctx context.Context, email string) (*profile.ProfileUserDTO, error)
	UpdateProfile(ctx context.Context, email string, data *profile.ProfileUser) error
	UpdateOneColoum(ctx context.Context, email string, coloum string, path string) error
	UpdateWorkExperience(ctx context.Context, email string, workExp *profile.WorkExperience) error
}

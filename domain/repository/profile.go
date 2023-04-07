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
	CreateWorkExperience(ctx context.Context, workExp *profile.WorkExperience) error
	GetListWorkExperience(ctx context.Context, email string) ([]*profile.WorkExperienceDTO, error)
	UpdateWorkExperience(ctx context.Context, id string, workExp *profile.WorkExperience) error
	DeletedWorkExperience(ctx context.Context, id string) error
}

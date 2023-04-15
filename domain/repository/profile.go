package repository

import (
	"CareerCenter/domain/entity/profile"
	"context"
)

type RepoProfile interface {
	CreateProfile(ctx context.Context, data *profile.ProfileUser) error
	GetProfileByEmail(ctx context.Context, email string) (*profile.ProfileUserDTO, error)
	UpdateProfile(ctx context.Context, email string, data *profile.ProfileUser) error
	UpdatePhotoProfile(ctx context.Context, email string, path string) error
	CreateWorkExperience(ctx context.Context, workExp *profile.WorkExperience) error
	GetListWorkExperience(ctx context.Context, email string) ([]*profile.WorkExperienceDTO, error)
	UpdateWorkExperience(ctx context.Context, id string, workExp *profile.WorkExperience) error
	DeletedWorkExperience(ctx context.Context, id string) error
	CreateEducation(ctx context.Context, education *profile.Education) error
	GetListEducation(ctx context.Context, email string) ([]*profile.EducationDTO, error)
	UpdateEducation(ctx context.Context, id string, education *profile.Education) error
	DeletedEducation(ctx context.Context, id string) error
	UpdateAbility(ctx context.Context, email string, ability string) error
	UpdateLanguage(ctx context.Context, email string, language string) error
	UpdateCvResume(ctx context.Context, email string, path string) error
	UpdatePortofolio(ctx context.Context, email string, path string) error
}

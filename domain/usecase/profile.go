package usecase

import (
	"CareerCenter/domain/entity/profile"
	"context"
)

type UseCaseProfile interface {
	GetProfileByEmail(ctx context.Context, email string) (*profile.ProfileUserDTO, error)
	UpdateProfile(ctx context.Context, email string, data *profile.ProfileUserDTO) error
	UpdatePhotoProfile(ctx context.Context, email string, path string) error
	CreateWorkExperiencet(ctx context.Context, email string, workExp *profile.WorkExperienceDTO) error
	UpdateWorkExperience(ctx context.Context, id string, workExp *profile.WorkExperienceDTO) error
	DeletedWorkExperience(ctx context.Context, id string) error
	CreateEducation(ctx context.Context, email string, education *profile.EducationDTO) error
	UpdateEducation(ctx context.Context, id string, education *profile.EducationDTO) error
	DeletedEducation(ctx context.Context, id string) error
}

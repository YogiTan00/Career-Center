package account

import (
	"errors"
	"time"
)

type UpdatePassword struct {
	oldPassword     string
	newPassword     string
	confirmPassword string
	updatedAt       time.Time
}

type UpdatePasswordDTO struct {
	OldPassword     string
	NewPassword     string
	ConfirmPassword string
	UpdatedAt       time.Time
}

func NewPassword(dto *UpdatePasswordDTO) (*UpdatePassword, error) {
	err := dto.Validation()
	if err != nil {
		return nil, err
	}
	return &UpdatePassword{
		oldPassword:     dto.OldPassword,
		newPassword:     dto.NewPassword,
		confirmPassword: dto.ConfirmPassword,
		updatedAt:       time.Now(),
	}, nil
}

func (dto *UpdatePasswordDTO) Validation() error {
	if dto.OldPassword == "" || dto.NewPassword == "" || dto.ConfirmPassword == "" {
		return errors.New("password cant be nil")
	}
	if dto.ConfirmPassword != dto.NewPassword {
		return errors.New("passwords are not the same")
	}
	return nil
}

func (data *UpdatePassword) GetOldPassword() string {
	return data.oldPassword
}
func (data *UpdatePassword) GetNewPassword() string {
	return data.newPassword
}
func (data *UpdatePassword) GetConfirmPassword() string {
	return data.confirmPassword
}
func (data *UpdatePassword) GetUpdateAt() time.Time {
	return data.updatedAt
}

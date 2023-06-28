package profile

import (
	"CareerCenter/domain/entity/profile"
	"CareerCenter/domain/mocks"
	"CareerCenter/testdata"
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestUseCaseProfileInteractor_GetProfileByEmail(t *testing.T) {
	type args struct {
		ctx   context.Context
		email string
	}
	ctx := context.TODO()
	data := testdata.TestDataProfile(3, 2)
	repoGetProfile := new(mocks.RepoProfile)
	repoGetProfile.On("GetProfileByEmail", mock.Anything, mock.Anything).
		Times(1).
		Return(data, nil)
	repoGetProfile.On("GetListWorkExperience", mock.Anything, mock.Anything).
		Times(1).
		Return(data.WorkExperience, nil)
	repoGetProfile.On("GetListEducation", mock.Anything, mock.Anything).
		Times(1).
		Return(data.Education, nil)

	tests := []struct {
		name    string
		u       UseCaseProfileInteractor
		args    args
		want    *profile.ProfileUserDTO
		wantErr bool
	}{
		{
			name: "sucess get profile",
			u:    UseCaseProfileInteractor{repoProfile: repoGetProfile},
			args: args{
				ctx:   ctx,
				email: data.Email,
			},
			want:    data,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.GetProfileByEmail(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCaseProfileInteractor.GetProfileByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCaseProfileInteractor.GetProfileByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

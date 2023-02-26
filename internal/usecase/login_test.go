package usecase

import (
	"CareerCenter/domain/entity"
	mock2 "CareerCenter/internal/repository/mock"
	"CareerCenter/utils"
	"context"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestUseCaseAccountInteractor_Login(t *testing.T) {
	type args struct {
		ctx      context.Context
		email    string
		password string
	}
	pw, _ := utils.HashPassword("abcd1234")
	data := &entity.AccountDTO{
		Email:    "yogi@gmail.com",
		Nama:     "yogi",
		Password: pw,
	}
	ctx := context.TODO()
	repoAccount := new(mock2.RepoAccount)
	repoAccount.On("GetByEmail", mock.Anything, mock.Anything).
		Times(1).
		Return(data, nil)

	tests := []struct {
		name    string
		u       UseCaseAccountInteractor
		args    args
		want    *entity.AccountDTO
		wantErr bool
	}{
		{
			name: "sucess log in",
			u:    UseCaseAccountInteractor{repoAccount},
			args: args{
				ctx:      ctx,
				email:    "yogi@gmail.com",
				password: "abcd1234",
			},
			want:    data,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.Login(tt.args.ctx, tt.args.email, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("UseCaseAccountInteractor.Login() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

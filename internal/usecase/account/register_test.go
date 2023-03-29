package account

import (
	"CareerCenter/domain/entity/account"
	"CareerCenter/domain/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestUseCaseRegisterInteractor_Register(t *testing.T) {
	type args struct {
		ctx  context.Context
		data *account.AccountDTO
	}
	ctx := context.TODO()
	data := &account.AccountDTO{
		Email:    "testing@gmail.com",
		Nama:     "testing",
		Password: "asdwas",
	}
	dataErr := &account.AccountDTO{
		Email:    "yogitest@gmail.com",
		Nama:     "testing",
		Password: "asdwas",
	}

	repoAccount := new(mocks.RepoAccount)
	repoAccount.On("CreateAccount", mock.Anything, mock.Anything).
		Times(1).
		Return(nil)

	repoRegisterErrRegister := new(mocks.RepoAccount)
	repoRegisterErrRegister.On("CreateAccount", mock.Anything, mock.Anything).
		Times(1).
		Return(errors.New("mock email telah terdaftar"))

	tests := []struct {
		name    string
		u       UseCaseAccountInteractor
		args    args
		wantErr bool
	}{
		{
			name: "sucess register",
			u:    UseCaseAccountInteractor{repoAccount: repoAccount},
			args: args{
				ctx:  ctx,
				data: data,
			},
			wantErr: false,
		},
		{
			name: "error register",
			u:    UseCaseAccountInteractor{repoAccount: repoRegisterErrRegister},
			args: args{
				ctx:  ctx,
				data: dataErr,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.u.Register(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UseCaseAccountInteractor.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

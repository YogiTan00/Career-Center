package account_test

import (
	account2 "CareerCenter/domain/entity/account"
	"CareerCenter/domain/mocks"
	"CareerCenter/domain/repository"
	"CareerCenter/internal/usecase/account"
	"CareerCenter/testdata"
	"CareerCenter/utils"
	"CareerCenter/utils/exceptions"
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestUseCaseAccountInteractor_UpdatePassword(t *testing.T) {
	type fields struct {
		repoAccount repository.RepoAccount
		repoProfile repository.RepoProfile
	}
	type args struct {
		ctx      context.Context
		email    string
		password *account2.UpdatePasswordDTO
	}
	ctx := context.TODO()
	data := testdata.TestDataAccountDTO()
	newPassword := testdata.TestDataPassword()

	repoAccount := new(mocks.RepoAccount)
	repoAccount.On("GetByEmail", mock.Anything, mock.Anything).
		Times(5).Return(data, nil)
	repoAccount.On("UpdatePassword", mock.Anything, mock.Anything, mock.Anything).
		Times(5).Return(nil)

	//Error get by email
	repoAccountErrEmail := new(mocks.RepoAccount)
	repoAccountErrEmail.On("GetByEmail", mock.Anything, mock.Anything).
		Times(5).Return(nil, exceptions.ErrCustomString("error mock get by email"))
	repoAccountErrEmail.On("UpdatePassword", mock.Anything, mock.Anything, mock.Anything).
		Times(5).Return(nil)

	//Error wrong password
	newPasswordErrWrong := testdata.TestDataPassword()
	newPasswordErrWrong.OldPassword = "awds"

	//Error cannot be changed with the same password
	newPasswordErrSamePw := testdata.TestDataPassword()
	newPasswordErrSamePw.NewPassword = newPassword.OldPassword
	newPasswordErrSamePw.ConfirmPassword = newPassword.OldPassword

	//Error cannot be changed with the same password
	newPasswordErrNewPw := testdata.TestDataPassword()
	newPasswordErrNewPw.NewPassword = ""
	newPasswordErrNewPw.ConfirmPassword = ""

	//Error password to long
	newPasswordErrHashPw := testdata.TestDataPassword()
	newPasswordErrHashPw.NewPassword = utils.RandomString(73)
	newPasswordErrHashPw.ConfirmPassword = newPasswordErrHashPw.NewPassword

	//Error update password
	repoAccountErrUpdate := new(mocks.RepoAccount)
	repoAccountErrUpdate.On("GetByEmail", mock.Anything, mock.Anything).
		Times(5).Return(data, nil)
	repoAccountErrUpdate.On("UpdatePassword", mock.Anything, mock.Anything, mock.Anything).
		Times(5).Return(exceptions.ErrCustomString("error mock update password"))

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "sucess update password",
			fields: fields{
				repoAccount: repoAccount,
				repoProfile: nil,
			},
			args: args{
				ctx:      ctx,
				email:    data.Email,
				password: newPassword,
			},
			wantErr: false,
		},
		{
			name: "error get by email",
			fields: fields{
				repoAccount: repoAccountErrEmail,
				repoProfile: nil,
			},
			args: args{
				ctx:      ctx,
				email:    data.Email,
				password: newPassword,
			},
			wantErr: true,
		},
		{
			name: "error wrong password",
			fields: fields{
				repoAccount: repoAccount,
				repoProfile: nil,
			},
			args: args{
				ctx:      ctx,
				email:    data.Email,
				password: newPasswordErrWrong,
			},
			wantErr: true,
		},
		{
			name: "error cannot be changed with the same password",
			fields: fields{
				repoAccount: repoAccount,
				repoProfile: nil,
			},
			args: args{
				ctx:      ctx,
				email:    data.Email,
				password: newPasswordErrSamePw,
			},
			wantErr: true,
		},
		{
			name: "error new password",
			fields: fields{
				repoAccount: repoAccount,
				repoProfile: nil,
			},
			args: args{
				ctx:      ctx,
				email:    data.Email,
				password: newPasswordErrNewPw,
			},
			wantErr: true,
		},
		{
			name: "error hash password",
			fields: fields{
				repoAccount: repoAccount,
				repoProfile: nil,
			},
			args: args{
				ctx:      ctx,
				email:    data.Email,
				password: newPasswordErrHashPw,
			},
			wantErr: true,
		},
		{
			name: "error update password",
			fields: fields{
				repoAccount: repoAccountErrUpdate,
				repoProfile: nil,
			},
			args: args{
				ctx:      ctx,
				email:    data.Email,
				password: newPassword,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := account.NewAccountUsecase(
				tt.fields.repoAccount,
				tt.fields.repoProfile,
			)
			if err := u.UpdatePassword(tt.args.ctx, tt.args.email, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("UpdatePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

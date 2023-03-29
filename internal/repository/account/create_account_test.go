package account_test

import (
	account2 "CareerCenter/domain/entity/account"
	"CareerCenter/internal/config/database"
	"CareerCenter/internal/repository/account"
	"CareerCenter/testdata"
	"context"
	"testing"
)

func TestRegisterMysqlInteractor_CreateRegister(t *testing.T) {
	type args struct {
		ctx  context.Context
		data *account2.Account
	}
	dbConn := database.InitMysqlDB()
	ctx := context.TODO()
	data := testdata.TestDataAccount()
	tests := []struct {
		name    string
		r       account.AccountMysqlInteractor
		args    args
		wantErr bool
	}{
		{
			name: "sucess register",
			r:    account.AccountMysqlInteractor{DbConn: dbConn},
			args: args{
				ctx:  ctx,
				data: data,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.CreateAccount(tt.args.ctx, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("AccountMysqlInteractor.CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

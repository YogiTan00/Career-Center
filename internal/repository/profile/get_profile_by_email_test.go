package profile

import (
	"CareerCenter/domain/entity/profile"
	"CareerCenter/internal/config/database"
	"context"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProfileMysqlInteractor_GetProfileByEmail(t *testing.T) {
	type args struct {
		ctx   context.Context
		email string
	}
	var (
		mysSqlConn  = database.InitMysqlDB()
		ctx         = context.TODO()
		repoProfile = NewProfileMysqlInteractor(mysSqlConn)
	)
	email := "yogi@gmail.com"
	wantSucess, _ := repoProfile.GetProfileByEmail(ctx, email)
	tests := []struct {
		name    string
		p       ProfileMysqlInteractor
		args    args
		want    *profile.ProfileUserDTO
		wantErr bool
	}{
		{
			name: "sucess get profile",
			p:    ProfileMysqlInteractor{DbConn: mysSqlConn},
			args: args{
				ctx:   ctx,
				email: email,
			},
			want:    wantSucess,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.p.GetProfileByEmail(tt.args.ctx, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProfileMysqlInteractor.GetProfileByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				assert.Equal(t, got, tt.want)
				t.Errorf("ProfileMysqlInteractor.GetProfileByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

package profile

import (
	"CareerCenter/domain/entity/profile"
	profile2 "CareerCenter/internal/repository/mapper/profile"
	profile3 "CareerCenter/internal/repository/models/profile"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (p ProfileMysqlInteractor) GetProfileByEmail(ctx context.Context, email string) (*profile.ProfileUserDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE email = ?`, profile3.GetTableNameProfile())
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: profile3.ProfileModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	result, err := dbq.Q(ctx, p.DbConn, stmt, opts, email)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}

	account := profile2.ModelProfileToEntity(result.(*profile3.ProfileModel))
	return account, nil
}

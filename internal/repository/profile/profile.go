package profile

import (
	"CareerCenter/domain/entity/profile"
	"CareerCenter/internal/repository/mapper"
	"CareerCenter/internal/repository/models"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (p ProfileMysqlInteractor) GetProfileByEmail(ctx context.Context, email string) (*profile.ProfileUserDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE email = ?`, models.GetTableNameProfile())
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: models.ProfileModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	result := dbq.MustQ(ctx, p.DbConn, stmt, opts, email)
	if result == nil {
		return nil, nil
	}

	account, errMap := mapper.ModelProfileToEntity(result.(*models.ProfileModel))
	if errMap != nil {
		return nil, errMap
	}
	return account, nil
}

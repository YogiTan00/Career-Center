package profile

import (
	"CareerCenter/domain/entity/profile"
	profile3 "CareerCenter/internal/repository/mapper/profile"
	profile2 "CareerCenter/internal/repository/models/profile"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (p ProfileMysqlInteractor) GetListWorkExperience(ctx context.Context, email string) ([]*profile.WorkExperienceDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE email = ? AND deleted_at IS NULL`, profile2.GetTableNameWorkExperience())

	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: profile2.WorkExperienceModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, p.DbConn, stmt, opts, email)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}

	workExprience := profile3.ModelListWorkExperienceToEntity(result.([]*profile2.WorkExperienceModel))

	return workExprience, nil
}

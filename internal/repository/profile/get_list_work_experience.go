package profile

import (
	"CareerCenter/domain/entity/profile"
	profile3 "CareerCenter/internal/repository/mapper/profile"
	profile2 "CareerCenter/internal/repository/models/profile"
	"context"
	"fmt"
	"time"

	"github.com/rocketlaunchr/dbq/v2"
)

func (p ProfileMysqlInteractor) GetListWorkExperience(ctx context.Context, email string) ([]*profile.WorkExperienceDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE email = ?`, profile2.GetTableNameWorkExperience())

	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: profile2.WorkExperienceModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result := dbq.MustQ(ctx, p.DbConn, stmt, opts, email)
	if result == nil {
		return nil, nil
	}

	workExprience := profile3.ModelListWorkExperienceToEntity(result.([]*profile2.WorkExperienceModel))

	return workExprience, nil
}

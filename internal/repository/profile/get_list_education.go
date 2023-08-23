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

func (p ProfileMysqlInteractor) GetListEducation(ctx context.Context, email string) ([]*profile.EducationDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE email = ? AND deleted_at IS NULL`, profile2.GetTableNameEducation())

	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: profile2.EducationModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result := dbq.MustQ(ctx, p.DbConn, stmt, opts, email)
	if result == nil {
		return nil, nil
	}

	education := profile3.ModelListEducationToEntity(result.([]*profile2.EducationModel))

	return education, nil
}

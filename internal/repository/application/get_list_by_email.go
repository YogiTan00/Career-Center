package application

import (
	"CareerCenter/domain/entity"
	"CareerCenter/internal/repository/mapper"
	"CareerCenter/internal/repository/models"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (a ApplicationMysqlInteractor) GetListByEmail(ctx context.Context, email string) ([]*entity.ApplicationDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE email = ?`, models.GetTableNameApplication())
	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: models.ApplicationModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	result := dbq.MustQ(ctx, a.DbConn, stmt, opts, email)
	if result == nil {
		return nil, nil
	}

	jobs := mapper.ModelApplicationListToEntity(result.([]*models.ApplicationModel))

	return jobs, nil
}
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

func (a ApplicationMysqlInteractor) GetListApplication(ctx context.Context) ([]*entity.ApplicationDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s`, models.GetTableNameApplication())
	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: models.ApplicationModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	result := dbq.MustQ(ctx, a.DbConn, stmt, opts)
	if result == nil {
		return nil, nil
	}

	job := mapper.ModelApplicationListToEntity(result.([]*models.ApplicationModel))

	return job, nil
}

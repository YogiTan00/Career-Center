package jobs

import (
	"CareerCenter/domain/entity"
	"CareerCenter/domain/entity/filter"
	"CareerCenter/domain/valueobject"
	"CareerCenter/internal/repository"
	"CareerCenter/internal/repository/mapper"
	"CareerCenter/internal/repository/models"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (j JobsMysqlInteractor) GetListJobs(ctx context.Context, typeSearch *valueobject.TypeSearch, f *filter.Filter) ([]*entity.JobsDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	fil := repository.TxQuery(typeSearch, f)
	stmt := fmt.Sprintf(`SELECT * FROM %s %s`, models.GetTableNameJobs(), fil)

	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: models.JobsModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result := dbq.MustQ(ctx, j.DbConn, stmt, opts)
	if result == nil {
		return nil, nil
	}

	jobs := mapper.ModelToEntityListJobs(result.([]*models.JobsModel))
	return jobs, nil
}

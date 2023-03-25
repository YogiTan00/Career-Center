package jobs

import (
	"CareerCenter/domain/entity"
	"CareerCenter/domain/entity/filter"
	"CareerCenter/internal/repository/mapper"
	"CareerCenter/internal/repository/models"
	"context"
	"errors"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (j JobsMysqlInteractor) GetListJobs(ctx context.Context, f *filter.Filter) ([]*entity.JobsDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	filter := txQuery(f)
	stmt := fmt.Sprintf(`SELECT * FROM %s %s`, models.GetTableNameJobs(), filter)

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
	if jobs == nil {
		return nil, errors.New("err mapper")
	}
	return jobs, nil
}

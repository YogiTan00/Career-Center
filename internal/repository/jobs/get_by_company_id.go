package jobs

import (
	"CareerCenter/domain/entity"
	"CareerCenter/internal/repository/mapper"
	"CareerCenter/internal/repository/models"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (j JobsMysqlInteractor) GetJobByCompanyId(ctx context.Context, id string) ([]*entity.JobsDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE company_id = ?`, models.GetTableNameJobs())
	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: models.JobsModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	result := dbq.MustQ(ctx, j.DbConn, stmt, opts, id)
	if result == nil {
		return nil, nil
	}

	job := mapper.ModelToEntityListJobs(result.([]*models.JobsModel))

	return job, nil
}

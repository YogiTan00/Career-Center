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

func (j JobsMysqlInteractor) GetJobById(ctx context.Context, id string) (*entity.JobsDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE id = ?`, models.GetTableNameJobs())
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: models.JobsModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	result, err := dbq.Q(ctx, j.DbConn, stmt, opts, id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}

	job := mapper.ModelToEntityJobs(result.(*models.JobsModel))

	return job, nil
}

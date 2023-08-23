package application

import (
	"CareerCenter/domain/entity"
	"CareerCenter/internal/repository/mapper"
	"CareerCenter/internal/repository/models"
	"CareerCenter/utils/exceptions"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (a ApplicationMysqlInteractor) GetByEmail(ctx context.Context, email string) (*entity.ApplicationDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE email = ?`, models.GetTableNameApplication())
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: models.ApplicationModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	result, err := dbq.Q(ctx, a.DbConn, stmt, opts, email)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, exceptions.ErrNotFound("email")
	}

	job := mapper.ModelApplicationToEntity(result.(*models.ApplicationModel))

	return job, nil
}

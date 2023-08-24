package company

import (
	"CareerCenter/domain/entity"
	"CareerCenter/internal/repository/mapper"
	"CareerCenter/internal/repository/models"
	"CareerCenter/utils/exceptions"
	"context"
	"fmt"
	"time"

	"github.com/rocketlaunchr/dbq/v2"
)

func (c CompanyMysqlInteractor) GetCompanyById(ctx context.Context, id string) (*entity.CompanyDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE id = ?`, models.GetTableNameCompany())
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: models.CompanyModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	result, err := dbq.Q(ctx, c.DbConn, stmt, opts, id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, exceptions.ErrNotFound("id")
	}

	job := mapper.ModelToEntityCompany(result.(*models.CompanyModel))

	return job, nil
}

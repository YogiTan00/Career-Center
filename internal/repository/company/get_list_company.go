package company

import (
	"CareerCenter/domain/entity"
	"CareerCenter/domain/entity/filter"
	"CareerCenter/domain/valueobject"
	"CareerCenter/internal/repository"
	"CareerCenter/internal/repository/mapper"
	"CareerCenter/internal/repository/models"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/rocketlaunchr/dbq/v2"
)

func (c CompanyMysqlInteractor) GetListCompany(ctx context.Context, typeSearch *valueobject.TypeSearch, f *filter.Filter) ([]*entity.CompanyDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	filter := repository.TxQuery(typeSearch, f)
	stmt := fmt.Sprintf(`SELECT * FROM %s %s`, models.GetTableNameCompany(), filter)

	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: models.CompanyModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result := dbq.MustQ(ctx, c.DbConn, stmt, opts)
	if result == nil {
		return nil, nil
	}

	jobs := mapper.ModelToEntityListCompany(result.([]*models.CompanyModel))
	if jobs == nil {
		return nil, errors.New("err mapper")
	}
	return jobs, nil
}

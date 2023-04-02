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

func (a ApplicationMysqlInteractor) GetByEmail(ctx context.Context, email string, companyId string) (*entity.ApplicationDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE email = ? AND company_id = ?`, models.GetTableNameApplication())
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: models.ApplicationModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	result := dbq.MustQ(ctx, a.DbConn, stmt, opts, email, companyId)
	if result == nil {
		return nil, nil
	}

	job := mapper.ModelApplicationToEntity(result.(*models.ApplicationModel))

	return job, nil
}

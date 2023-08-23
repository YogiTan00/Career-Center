package account

import (
	"CareerCenter/domain/entity/account"
	"CareerCenter/internal/repository/mapper"
	"CareerCenter/internal/repository/models"
	"CareerCenter/utils/exceptions"
	"context"
	"fmt"
	"time"

	"github.com/rocketlaunchr/dbq/v2"
)

func (l AccountMysqlInteractor) GetByEmail(ctx context.Context, email string) (*account.AccountDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE email = ?`, models.GetTableNameAccount())
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: models.AccountModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, l.DbConn, stmt, opts, email)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, exceptions.ErrNotFound("email")
	}

	accountMapper := mapper.ModelToEntity(result.(*models.AccountModel))

	return accountMapper, nil
}

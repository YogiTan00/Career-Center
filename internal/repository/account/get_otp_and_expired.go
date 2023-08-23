package account

import (
	"CareerCenter/domain/entity/account"
	"CareerCenter/internal/repository/mapper"
	"CareerCenter/internal/repository/models"
	"CareerCenter/utils/exceptions"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (r AccountMysqlInteractor) GetOTP(ctx context.Context, email string) (*account.CodeOTP, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE email = ?`, models.GetTableNameAccount())
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: models.CodeOTP{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, r.DbConn, stmt, opts, email)
	if err != nil {
		return nil, exceptions.ErrorWrongEmailorPassword
	}
	if result == nil {
		return nil, exceptions.ErrNotFound("email")
	}

	otp := mapper.ModelOTPToEntity(result.(*models.CodeOTP))

	return otp, nil
}

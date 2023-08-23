package account

import (
	"CareerCenter/internal/repository/models"
	"context"
	"fmt"
	"time"

	"github.com/rocketlaunchr/dbq/v2"
)

func (l AccountMysqlInteractor) RemoveOTP(ctx context.Context, email string) error {
	timeExp := time.Now().Add(time.Duration(-100) * time.Minute)
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf("UPDATE %s SET code_otp = ?, expired_otp = ? WHERE email = ? ", models.GetTableNameAccount())

	_, err := dbq.E(ctx, l.DbConn, query, nil, "", timeExp, email)

	if err != nil {
		return err
	}

	return nil
}

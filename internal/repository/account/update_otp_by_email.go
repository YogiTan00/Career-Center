package account

import (
	"CareerCenter/internal/repository/models"
	"context"
	"fmt"
	"time"

	"github.com/rocketlaunchr/dbq/v2"
)

func (l AccountMysqlInteractor) UpdateOTP(ctx context.Context, email string, otp string) error {
	timeExp := time.Now().Add(10 * time.Minute)
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf("UPDATE %s SET code_otp = ?, expired_otp = ? WHERE email = ? ", models.GetTableNameAccount())

	_, err := dbq.E(ctx, l.DbConn, query, nil, otp, timeExp, email)

	if err != nil {
		return err
	}

	return nil
}

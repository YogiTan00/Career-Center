package account

import (
	"CareerCenter/internal/repository/models"
	"context"
	"fmt"
	"time"

	"github.com/rocketlaunchr/dbq/v2"
)

func (l AccountMysqlInteractor) UpdateOTP(ctx context.Context, email string, otp string) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	currentTime := time.Now()
	tenMinutesLater := currentTime.Add(10 * time.Minute)
	formattedTime := tenMinutesLater.Format("2006-01-02 15:04:05")

	query := fmt.Sprintf("UPDATE %s SET code_otp = '%s', expired_otp = '%s' WHERE email = '%s' ", models.GetTableNameAccount(), otp, formattedTime, email)

	_, err := dbq.E(ctx, l.DbConn, query, nil)

	if err != nil {
		return err
	}

	return nil
}

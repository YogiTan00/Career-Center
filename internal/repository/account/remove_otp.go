package account

import (
	"CareerCenter/internal/repository/models"
	"context"
	"fmt"
	"time"

	"github.com/rocketlaunchr/dbq/v2"
)

func (l AccountMysqlInteractor) RemoveOTP(ctx context.Context, email string) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	currentTime := time.Now()
	later := currentTime.Add(time.Duration(-100) * time.Minute)
	formattedTime := later.Format("2006-01-02 15:04:05")

	query := fmt.Sprintf("UPDATE %s SET code_otp = '%s', expired_otp = '%s' WHERE email = '%s' ", models.GetTableNameAccount(), "", formattedTime, email)

	_, err := dbq.E(ctx, l.DbConn, query, nil)

	if err != nil {
		return err
	}

	return nil
}

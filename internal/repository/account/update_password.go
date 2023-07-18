package account

import (
	"CareerCenter/internal/repository/models"
	"context"
	"fmt"
	"time"

	"github.com/rocketlaunchr/dbq/v2"
)

func (l AccountMysqlInteractor) UpdatePassword(ctx context.Context, email string, password string) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	query := fmt.Sprintf("UPDATE %s SET password = '%s', updated_at = '%s' WHERE email = '%s' ", models.GetTableNameAccount(), password, formattedTime, email)

	_, err := dbq.E(ctx, l.DbConn, query, nil)

	if err != nil {
		return err
	}

	return nil
}

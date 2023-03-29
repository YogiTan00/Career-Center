package account

import (
	"CareerCenter/internal/repository/models"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (l AccountMysqlInteractor) UpdatePassword(ctx context.Context, email string, password string) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf("UPDATE %s SET password = '%s' WHERE email = '%s' ", models.GetTableNameAccount(), password, email)

	_, err := dbq.E(ctx, l.DbConn, query, nil)

	if err != nil {
		return err
	}

	return nil
}

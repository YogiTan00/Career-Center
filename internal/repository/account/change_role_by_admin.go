package account

import (
	"CareerCenter/domain/entity/account"
	"CareerCenter/internal/repository/models"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (r AccountMysqlInteractor) UpdateRole(ctx context.Context, data *account.Account) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf("UPDATE %s SET role = '%s' WHERE email = '%s' ", models.GetTableNameAccount(), data.GetRole().StringRoles(), data.GetEmail())

	_, err := dbq.E(ctx, r.DbConn, query, nil)

	if err != nil {
		return err
	}

	return nil
}

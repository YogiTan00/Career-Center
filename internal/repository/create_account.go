package repository

import (
	"CareerCenter/domain/entity"
	"CareerCenter/internal/repository/models"
	"context"
	"fmt"
	"time"
)

func (r AccountMysqlInteractor) CreateAccount(ctx context.Context, data *entity.Account) error {
	var (
		err error
	)
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	txQuery := fmt.Sprintf("INSERT INTO %s (email, nama, password, created_at, updated_at) values ('%s','%s','%s','%s','%s')",
		models.GetTableName(), data.GetEmail(), data.GetNama(), data.GetPassword(), data.GetCreatedAt(), data.GetUpdatedAt())
	_, err = r.DbConn.QueryContext(ctx, txQuery)
	if err != nil {
		return err
	}

	return err
}

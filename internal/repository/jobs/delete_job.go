package jobs

import (
	"CareerCenter/internal/repository/models"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (j JobsMysqlInteractor) DeleteJobById(ctx context.Context, id string) error {
	timeNow := time.Now()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf("UPDATE %s SET deleted_at = ? WHERE id = ?", models.GetTableNameJobs())

	_, err := dbq.E(ctx, j.DbConn, query, nil, timeNow, id)

	if err != nil {
		return err
	}

	return nil
}

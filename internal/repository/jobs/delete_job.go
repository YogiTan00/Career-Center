package jobs

import (
	"CareerCenter/internal/repository/models"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (j JobsMysqlInteractor) DeleteJobById(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", models.GetTableNameJobs())

	_, err := dbq.E(ctx, j.DbConn, query, nil, id)

	if err != nil {
		return err
	}

	return nil
}
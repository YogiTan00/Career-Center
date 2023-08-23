package profile

import (
	"CareerCenter/internal/repository/models/profile"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (p ProfileMysqlInteractor) DeletedEducation(ctx context.Context, id string) error {
	timeNow := time.Now()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf("UPDATE %s SET deleted_at = ? WHERE id = ?", profile.GetTableNameEducation())

	_, err := dbq.E(ctx, p.DbConn, query, nil, timeNow, id)

	if err != nil {
		return err
	}

	return nil
}

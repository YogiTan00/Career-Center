package profile

import (
	"CareerCenter/internal/repository/models/profile"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (p ProfileMysqlInteractor) UpdatePortofolio(ctx context.Context, email string, path string) error {
	timeNow := time.Now()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf("UPDATE %s SET portofolio = ?, updated_at = ? WHERE email = ? ",
		profile.GetTableNameProfile())

	_, err := dbq.E(ctx, p.DbConn, query, nil, path, timeNow, email)

	if err != nil {
		return err
	}

	return nil
}

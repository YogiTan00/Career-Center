package profile

import (
	"CareerCenter/internal/repository/models/profile"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (p ProfileMysqlInteractor) UpdatePhotoProfile(ctx context.Context, email string, path string) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf("UPDATE %s SET photo = ? WHERE email = ? ", profile.GetTableNameProfile())

	_, err := dbq.E(ctx, p.DbConn, query, nil, path, email)

	if err != nil {
		return err
	}

	return nil
}

package profile

import (
	"CareerCenter/internal/repository/models"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (p ProfileMysqlInteractor) UpdatePhotoProfile(ctx context.Context, email string, photo string) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf("UPDATE %s SET photo='%s' WHERE email = '%s' ",
		models.GetTableNameProfile(), photo, email)

	_, err := dbq.E(ctx, p.DbConn, query, nil)

	if err != nil {
		return err
	}

	return nil
}

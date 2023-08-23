package profile

import (
	"CareerCenter/internal/repository/models/profile"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (p ProfileMysqlInteractor) UpdateLanguage(ctx context.Context, email string, language string) error {
	timeNow := time.Now()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf("UPDATE %s SET language = ?, updated_at WHERE email = ? ",
		profile.GetTableNameProfile())

	_, err := dbq.E(ctx, p.DbConn, query, nil, language, timeNow, email)

	if err != nil {
		return err
	}

	return nil
}

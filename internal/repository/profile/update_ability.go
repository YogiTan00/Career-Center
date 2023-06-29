package profile

import (
	"CareerCenter/internal/repository/models/profile"
	"context"
	"fmt"
	"time"

	"github.com/rocketlaunchr/dbq/v2"
)

func (p ProfileMysqlInteractor) UpdateAbility(ctx context.Context, email string, ability string) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf("UPDATE %s SET ability ='%s' WHERE email = '%s' ",
		profile.GetTableNameProfile(), ability, email)

	_, err := dbq.E(ctx, p.DbConn, query, nil)

	if err != nil {
		return err
	}

	return nil
}

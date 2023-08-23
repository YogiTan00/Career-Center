package profile

import (
	"CareerCenter/domain/entity/profile"
	profile2 "CareerCenter/internal/repository/models/profile"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (p ProfileMysqlInteractor) UpdateProfile(ctx context.Context, email string, data *profile.ProfileUser) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf("UPDATE %s SET name = ?, skill = ?, phone_number = ?, updated_at = ? WHERE email = ? ", profile2.GetTableNameProfile())

	_, err := dbq.E(ctx, p.DbConn, query, nil, data.GetName(), data.GetSkill(), data.GetPhoneNumber(), data.GetUpdatedAt(), email)

	if err != nil {
		return err
	}

	return nil
}

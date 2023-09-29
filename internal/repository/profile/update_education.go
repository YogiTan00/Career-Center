package profile

import (
	"CareerCenter/domain/entity/profile"
	profile2 "CareerCenter/internal/repository/models/profile"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (p ProfileMysqlInteractor) UpdateEducation(ctx context.Context, id string, education *profile.Education) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	query := fmt.Sprintf("UPDATE %s SET level = ?, name = ?, major = ?, still_education = ?, start_education = ? , end_education = ?, description = ?, updated_at = ? WHERE id = ? ",
		profile2.GetTableNameEducation())

	_, err := dbq.E(ctx, p.DbConn, query, nil, education.GetLevel().StringLevel(), education.GetName(), education.GetMajor(), education.GetStillEducation(), education.GetStartEducation(), education.GetEndEducation(), education.GetDescription(), education.GetUpdatedAt(), id)

	if err != nil {
		return err
	}

	return nil
}

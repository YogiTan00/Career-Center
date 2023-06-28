package profile

import (
	"CareerCenter/domain/entity/profile"
	profile2 "CareerCenter/internal/repository/models/profile"
	"context"
	"fmt"
	"time"

	"github.com/rocketlaunchr/dbq/v2"
)

func (p ProfileMysqlInteractor) UpdateEducation(ctx context.Context, id string, education *profile.Education) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf("UPDATE %s SET level='%s', name= '%s', major= '%s', still_education = '%t', start_education= '%v' , end_education= '%v', description= '%s' WHERE id = '%s' ",
		profile2.GetTableNameEducation(), education.GetLevel(), education.GetName(), education.GetMajor(), education.GetStillEducation(), education.GetStartEducation(), education.GetEndEducation(), education.GetDescription(), id)

	_, err := dbq.E(ctx, p.DbConn, query, nil)

	if err != nil {
		return err
	}

	return nil
}

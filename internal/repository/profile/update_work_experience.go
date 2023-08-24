package profile

import (
	"CareerCenter/domain/entity/profile"
	profile2 "CareerCenter/internal/repository/models/profile"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (p ProfileMysqlInteractor) UpdateWorkExperience(ctx context.Context, id string, workExp *profile.WorkExperience) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	var strEndWork string
	if workExp.GetEndWork().IsZero() {
		strEndWork = "1000-10-10"
	} else {
		strEndWork = workExp.GetEndWork().String()
	}
	query := fmt.Sprintf("UPDATE %s SET skill_experience = ?, name = ?, still_working = ?, start_work = ? , end_work = ?, description = ?, updated_at = ? WHERE id = ? ", profile2.GetTableNameWorkExperience())

	_, err := dbq.E(ctx, p.DbConn, query, nil, workExp.GetSkillExperience(), workExp.GetName(), workExp.GetStillWorking(), workExp.GetStartWork(), strEndWork, workExp.GetDescription(), workExp.GetUpdatedAt(), id)

	if err != nil {
		return err
	}

	return nil
}

package profile

import (
	"CareerCenter/domain/entity/profile"
	profile2 "CareerCenter/internal/repository/models/profile"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (p ProfileMysqlInteractor) UpdateWorkExperience(ctx context.Context, email string, workExp *profile.WorkExperience) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf("UPDATE %s SET skill_experience='%s', name= '%s', still_working= '%t', start_work= '%v' , end_work= '%v', description= '%s' WHERE email = '%s' ",
		profile2.GetTableNameWorkExperience(), workExp.GetSkillExperience(), workExp.GetName(), workExp.GetStillWorking(), workExp.GetStartWork(), workExp.GetEndWork(), workExp.GetDescription(), email)

	_, err := dbq.E(ctx, p.DbConn, query, nil)

	if err != nil {
		return err
	}

	return nil
}

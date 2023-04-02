package profile

import (
	"CareerCenter/domain/entity/profile"
	"context"
)

func (p ProfileMysqlInteractor) UpdateWorkExperience(ctx context.Context, email string, workExp *profile.WorkExperience) error {
	//ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	//defer cancel()
	//
	//query := fmt.Sprintf("UPDATE %s SET photo='%s' WHERE email = '%s' ",
	//	models.GetTableNameProfile(), photo, email)
	//
	//_, err := dbq.E(ctx, p.DbConn, query, nil)
	//
	//if err != nil {
	//	return err
	//}

	return nil
}

package application

import (
	"CareerCenter/domain/entity"
	"CareerCenter/internal/repository/models"
	"context"
	"fmt"
	"time"
)

func (a ApplicationMysqlInteractor) SendApplication(ctx context.Context, application *entity.Application) error {
	var (
		err error
	)
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	txQuery := fmt.Sprintf("INSERT INTO %s (id, company_id, email, name , skill, phone_number,cv_resume, portofolio, created_at, updated_at) "+
		"values ('%s','%s','%s','%s','%s','%s','%s','%s','%v','%v')",
		models.GetTableNameApplication(),
		application.GetId(), application.GetCompanyId(), application.GetEmail(),
		application.GetName(), application.GetSkill(), application.GetPhoneNumber(),
		application.GetCvResume(), application.GetPortofolio(), application.GetCreatedAt(),
		application.GetUpdatedAt())
	_, err = a.DbConn.QueryContext(ctx, txQuery)
	if err != nil {
		return err
	}

	return err
}

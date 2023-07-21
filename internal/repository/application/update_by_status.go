package application

import (
	"CareerCenter/domain/entity"
	"CareerCenter/internal/repository/models"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (a ApplicationMysqlInteractor) UpdateStatusById(ctx context.Context, applicant *entity.StatusApplicantRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf("UPDATE %s SET status ='%s' WHERE id = '%s' ",
		models.GetTableNameApplication(), applicant.Status, applicant.Id)

	_, err := dbq.E(ctx, a.DbConn, query, nil)

	if err != nil {
		return err
	}

	return nil
}

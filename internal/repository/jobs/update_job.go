package jobs

import (
	"CareerCenter/domain/entity"
	"CareerCenter/internal/repository/mapper"
	"CareerCenter/internal/repository/models"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (j JobsMysqlInteractor) UpdateJobById(ctx context.Context, data *entity.Jobs) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf("UPDATE %s SET id=?, company_id = ?, position = ?,company = ?,logo = ?, address = ? ,status = ?,qualification = ?,job_description = ?,category = ?,created_at = ?,updated_at = ?,deleted_at = ? WHERE id = '%s' ", models.GetTableNameJobs(), data.GetId())

	_, err := dbq.E(ctx, j.DbConn, query, nil, mapper.DomainJobToInterface(data))

	if err != nil {
		return err
	}

	return nil
}

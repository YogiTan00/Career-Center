package company

import (
	"CareerCenter/internal/repository/models"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (c CompanyMysqlInteractor) DeleteCompanyById(ctx context.Context, companyId string) error {
	timeNow := time.Now()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf("UPDATE %s SET deleted_at=? WHERE id = ?", models.GetTableNameCompany())

	_, err := dbq.E(ctx, c.DbConn, query, nil, timeNow, companyId)

	if err != nil {
		return err
	}

	return nil
}

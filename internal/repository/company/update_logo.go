package company

import (
	"CareerCenter/internal/repository/models"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (c CompanyMysqlInteractor) UpdateLogoCompanyById(ctx context.Context, companyId string, path string) error {
	timeNow := time.Now()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	query := fmt.Sprintf("UPDATE %s SET logo = ?,updated_at = ? WHERE id = '%s' ", models.GetTableNameCompany(), companyId)

	_, err := dbq.E(ctx, c.DbConn, query, nil, path, timeNow)

	if err != nil {
		return err
	}

	return nil
}

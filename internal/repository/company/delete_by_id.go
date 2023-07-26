package company

import (
	"CareerCenter/internal/repository/models"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (c CompanyMysqlInteractor) DeleteCompanyById(ctx context.Context, companyId string) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf("DELETE FROM %s WHERE id = ?", models.GetTableNameCompany())

	_, err := dbq.E(ctx, c.DbConn, query, nil, companyId)

	if err != nil {
		return err
	}

	return nil
}

package company

import (
	"CareerCenter/domain/entity"
	"CareerCenter/internal/repository/models"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (c CompanyMysqlInteractor) UpdateCompanyById(ctx context.Context, companyId string, company *entity.Company) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	query := fmt.Sprintf("UPDATE %s SET email = ?, name = ?, type_company = ?, address = ?, logo = ? , profile = ?, website = ?, location = ?,updated_at = ?  WHERE id = ? ", models.GetTableNameCompany())

	_, err := dbq.E(ctx, c.DbConn, query, nil, company.GetEmail(), company.GetName(), company.GetTypeCompany().StringCompany(), company.GetAddress(), company.GetLogo(), company.About().GetProfile(), company.About().GetWebsite(), company.About().GetLocation(), company.GetUpdateAt(), companyId)

	if err != nil {
		return err
	}

	return nil
}

package company

import (
	"CareerCenter/domain/entity"
	"CareerCenter/internal/repository/mapper"
	"CareerCenter/internal/repository/models"
	"context"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (c CompanyMysqlInteractor) CreateCompany(ctx context.Context, company *entity.Company) error {
	var err error
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	err = dbq.Tx(ctx, c.DbConn, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		postModelStruct := mapper.DomainCompanyToInterface(company)

		stmt := dbq.INSERTStmt(models.GetTableNameCompany(), models.TableCompany(), len(postModelStruct), dbq.MySQL)

		_, errStore := E(ctx, stmt, nil, postModelStruct)

		if errStore != nil {
			err = errStore
		}

		_ = txCommit()
	})

	return err
}

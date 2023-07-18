package account

import (
	"CareerCenter/domain/entity/account"
	"CareerCenter/internal/repository/mapper"
	"CareerCenter/internal/repository/models"
	"context"
	"errors"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/rocketlaunchr/dbq/v2"
)

func (r AccountMysqlInteractor) CreateAccount(ctx context.Context, data *account.Account) error {
	var (
		checkErr error
	)
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	err := dbq.Tx(ctx, r.DbConn, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		postModelStruct := mapper.DomainToInterface(data)
		stmt := dbq.INSERTStmt(models.GetTableNameAccount(), models.TableAccount(), len(postModelStruct), dbq.MySQL)

		_, errStore := E(ctx, stmt, nil, postModelStruct)

		if errStore != nil {
			if mysqlErr, ok := errStore.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
				checkErr = errors.New("e-mail has been registered")
			} else {
				checkErr = errStore
			}

		}
		_ = txCommit()
	})
	if checkErr != nil {
		return checkErr
	}

	return err
}

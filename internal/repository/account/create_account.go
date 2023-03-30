package account

import (
	"CareerCenter/domain/entity/account"
	"CareerCenter/internal/repository/mapper"
	"CareerCenter/internal/repository/models"
	"CareerCenter/utils/exceptions"
	"context"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (r AccountMysqlInteractor) CreateAccount(ctx context.Context, data *account.Account) error {
	var (
		checkErr bool
	)
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	err := dbq.Tx(ctx, r.DbConn, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		postModelStruct := mapper.DomainToInterface(data)

		stmt := dbq.INSERTStmt(models.GetTableNameAccount(), models.TableAccount(), len(postModelStruct), dbq.MySQL)

		_, errStore := E(ctx, stmt, nil, postModelStruct)

		if errStore != nil {
			checkErr = true
		}

		_ = txCommit()
	})

	if checkErr == true {
		return exceptions.ErrCustomString("e-mail has been registered")
	}

	return err
}

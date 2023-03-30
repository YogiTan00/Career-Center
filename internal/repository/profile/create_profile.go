package profile

import (
	"CareerCenter/domain/entity/profile"
	"CareerCenter/internal/repository/mapper"
	"CareerCenter/internal/repository/models"
	"context"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (p ProfileMysqlInteractor) CreateProfile(ctx context.Context, data *profile.ProfileUser) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	err := dbq.Tx(ctx, p.DbConn, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		postModelStruct := mapper.DomainProfileToInterface(data)

		stmt := dbq.INSERTStmt(models.GetTableNameProfile(), models.TableProfile(), len(postModelStruct), dbq.MySQL)

		_, errStore := E(ctx, stmt, nil, postModelStruct)

		if errStore != nil {
			panic(errStore)
		}

		_ = txCommit()
	})

	return err
}

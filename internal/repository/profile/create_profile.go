package profile

import (
	"CareerCenter/domain/entity/profile"
	profile2 "CareerCenter/internal/repository/mapper/profile"
	profile3 "CareerCenter/internal/repository/models/profile"
	"context"
	"time"

	"github.com/rocketlaunchr/dbq/v2"
)

func (p ProfileMysqlInteractor) CreateProfile(ctx context.Context, data *profile.ProfileUser) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	err := dbq.Tx(ctx, p.DbConn, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		postModelStruct := profile2.DomainProfileToInterface(data)

		stmt := dbq.INSERTStmt(profile3.GetTableNameProfile(), profile3.TableProfile(), len(postModelStruct), dbq.MySQL)

		_, errStore := E(ctx, stmt, nil, postModelStruct)

		if errStore != nil {
			panic(errStore)
		}

		_ = txCommit()
	})

	return err
}

package profile

import (
	"CareerCenter/domain/entity/profile"
	profile2 "CareerCenter/internal/repository/mapper/profile"
	profile3 "CareerCenter/internal/repository/models/profile"
	"context"
	"fmt"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (p ProfileMysqlInteractor) CreateEducation(ctx context.Context, education *profile.Education) error {
	var err error
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()
	fmt.Println(education.GetCreatedAt(), education.GetUpdatedAt())
	err = dbq.Tx(ctx, p.DbConn, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		postModelStruct := profile2.DomainEducationToInterface(education)

		stmt := dbq.INSERTStmt(profile3.GetTableNameEducation(), profile3.TableEducation(), len(postModelStruct), dbq.MySQL)

		_, errStore := E(ctx, stmt, nil, postModelStruct)

		if errStore != nil {
			err = errStore
		}

		_ = txCommit()
	})

	return err
}

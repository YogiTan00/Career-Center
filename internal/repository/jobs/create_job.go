package jobs

import (
	"CareerCenter/domain/entity"
	"CareerCenter/internal/repository/mapper"
	"CareerCenter/internal/repository/models"
	"context"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

func (j JobsMysqlInteractor) CreateJob(ctx context.Context, data *entity.Jobs) error {
	var err error
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	err = dbq.Tx(ctx, j.DbConn, func(tx interface{}, Q dbq.QFn, E dbq.EFn, txCommit dbq.TxCommit) {
		postModelStruct := mapper.DomainJobToInterface(data)

		stmt := dbq.INSERTStmt(models.GetTableNameJobs(), models.TableJob(), len(postModelStruct), dbq.MySQL)

		_, errStore := E(ctx, stmt, nil, postModelStruct)

		if errStore != nil {
			err = errStore
		}

		_ = txCommit()
	})

	return err
}

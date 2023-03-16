package jobs

import (
	"CareerCenter/domain/repository"
	"database/sql"
)

type JobsMysqlInteractor struct {
	DbConn *sql.DB
}

func NewJobsMysqlInteractor(conndb *sql.DB) repository.RepoJobs {
	return &JobsMysqlInteractor{DbConn: conndb}
}

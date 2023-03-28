package application

import (
	"CareerCenter/domain/repository"
	"database/sql"
)

type ApplicationMysqlInteractor struct {
	DbConn *sql.DB
}

func NewApplicationMysqlInteractor(conndb *sql.DB) repository.RepoApplication {
	return &ApplicationMysqlInteractor{DbConn: conndb}
}

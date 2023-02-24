package repository

import (
	"CareerCenter/domain/repository"
	"database/sql"
)

type AccountMysqlInteractor struct {
	DbConn *sql.DB
}

func NewAccountMysqlInteractor(conndb *sql.DB) repository.RepoAccount {
	return &AccountMysqlInteractor{DbConn: conndb}
}

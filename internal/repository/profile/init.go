package profile

import (
	"CareerCenter/domain/repository"
	"database/sql"
)

type ProfileMysqlInteractor struct {
	DbConn *sql.DB
}

func NewProfileMysqlInteractor(conndb *sql.DB) repository.RepoProfile {
	return &ProfileMysqlInteractor{DbConn: conndb}
}

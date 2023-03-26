package company

import (
	"CareerCenter/domain/repository"
	"database/sql"
)

type CompanyMysqlInteractor struct {
	DbConn *sql.DB
}

func NewCompanyMysqlInteractor(conndb *sql.DB) repository.RepoCompany {
	return &CompanyMysqlInteractor{DbConn: conndb}
}

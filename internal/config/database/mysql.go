package database

import (
	"CareerCenter/logger"
	"CareerCenter/package/cfg"
	"database/sql"
	"fmt"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitMysqlDB(config cfg.Config) *sql.DB {
	var (
		errMysql error
		dbConn   *sql.DB
		log      = logger.NewLogger("Init Mysql")
	)

	dbHost := config.DB_HOST //localhost
	dbPort := config.DB_PORT //3306
	dbUser := config.DB_USER //kolaborasisalt_kolaborasisalt
	dbPass := config.DB_PASS //Ky4F-E*Yb^XT or KolaboraSalt
	dbName := config.DB_NAME //kolaborasisalt_career_center

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	log.Info(connection)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	dbConn, errMysql = sql.Open(`mysql`, dsn)

	if errMysql != nil {
		panic(errMysql)
	}

	dbConn.SetMaxOpenConns(300)
	dbConn.SetMaxIdleConns(25)
	dbConn.SetConnMaxLifetime(5 * time.Minute)

	return dbConn
}

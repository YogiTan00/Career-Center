package database

import (
	"CareerCenter/logger"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func InitMysqlDB() *sql.DB {
	var (
		errMysql error
		dbConn   *sql.DB
		log      = logger.NewLogger("Init Mysql")
	)

	dbHost := "localhost"     //localhost
	dbPort := "3306"          //3306
	dbUser := "root"          //kolaborasisalt_kolaborasisalt
	dbPass := ""              //Ky4F-E*Yb^XT or KolaboraSalt
	dbName := "career_center" //kolaborasisalt_career_center

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

	errDbConn := dbConn.Ping()
	if errDbConn != nil {
		if errors.Unwrap(errDbConn) != nil {
			log.Error(errors.Unwrap(errDbConn).Error())
		} else {
			log.Error(errDbConn.Error())
		}
	}

	dbConn.SetMaxOpenConns(300)
	dbConn.SetMaxIdleConns(25)
	dbConn.SetConnMaxLifetime(5 * time.Minute)

	return dbConn
}

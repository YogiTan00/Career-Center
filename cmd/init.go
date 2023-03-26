package main

import (
	"CareerCenter/internal/config/database"
	http2 "CareerCenter/internal/delivery/http/account"
	company3 "CareerCenter/internal/delivery/http/company"
	jobs3 "CareerCenter/internal/delivery/http/jobs"
	profile3 "CareerCenter/internal/delivery/http/profile"
	"CareerCenter/internal/repository/account"
	"CareerCenter/internal/repository/company"
	"CareerCenter/internal/repository/jobs"
	"CareerCenter/internal/repository/profile"
	account2 "CareerCenter/internal/usecase/account"
	company2 "CareerCenter/internal/usecase/company"
	jobs2 "CareerCenter/internal/usecase/jobs"
	profile2 "CareerCenter/internal/usecase/profile"
)

var (
	mysqlConn = database.InitMysqlDB()

	repoAccount    = account.NewAccountMysqlInteractor(mysqlConn)
	useCaseAccount = account2.NewAccountUsecase(repoAccount)
	handlerAccount = http2.NewUseCaseAccountHandler(useCaseAccount)

	repoJobs    = jobs.NewJobsMysqlInteractor(mysqlConn)
	useCaseJobs = jobs2.NewJobsUsecase(repoJobs)
	handlerJobs = jobs3.NewUseCaseJobsHandler(useCaseJobs)

	repoProfile    = profile.NewProfileMysqlInteractor(mysqlConn)
	useCaseProfile = profile2.NewProfileUsecase(repoProfile)
	handlerProfile = profile3.NewUseCaseProfileHandler(useCaseProfile)

	repoCompany    = company.NewCompanyMysqlInteractor(mysqlConn)
	useCaseCompany = company2.NewCompanyUsecase(repoCompany, repoJobs)
	handlerCompany = company3.NewUseCaseCompanyHandler(useCaseCompany)
)

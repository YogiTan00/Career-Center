package main

import (
	"CareerCenter/internal/config/database"
	http2 "CareerCenter/internal/delivery/http/account"
	home2 "CareerCenter/internal/delivery/http/home"
	jobs3 "CareerCenter/internal/delivery/http/jobs"
	"CareerCenter/internal/repository/account"
	"CareerCenter/internal/repository/jobs"
	account2 "CareerCenter/internal/usecase/account"
	"CareerCenter/internal/usecase/home"
	jobs2 "CareerCenter/internal/usecase/jobs"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	mysqlConn = database.InitMysqlDB()

	repoAccount    = account.NewAccountMysqlInteractor(mysqlConn)
	useCaseAccount = account2.NewAccountUsecase(repoAccount)

	repoJobs    = jobs.NewJobsMysqlInteractor(mysqlConn)
	useCaseJobs = jobs2.NewJobsUsecase(repoJobs)

	useCaseHome = home.NewHomeUseCase(repoAccount, repoJobs)

	handlerAccount = http2.NewUseCaseAccountHandler(useCaseAccount)
	handlerJobs    = jobs3.NewUseCaseJobsHandler(useCaseJobs)
	handlerHome    = home2.NewUseCaseJobsHandler(useCaseHome)
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", ParamHandlerWithoutInput).Methods(http.MethodGet)

	r.HandleFunc("/v1/register", handlerAccount.Register).Methods(http.MethodPost)
	r.HandleFunc("/v1/login", handlerAccount.Login).Methods(http.MethodPost)

	r.HandleFunc("/v1/list-jobs", handlerJobs.GetListJob).Methods(http.MethodGet)

	r.HandleFunc("/v1/home", handlerHome.GetHome).Methods(http.MethodGet)

	http.ListenAndServe(":8080", r)
}

func ParamHandlerWithoutInput(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "SUCCESS OK")
}

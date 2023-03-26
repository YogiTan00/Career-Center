package main

import (
	"CareerCenter/internal/config/database"
	http2 "CareerCenter/internal/delivery/http/account"
	jobs3 "CareerCenter/internal/delivery/http/jobs"
	profile3 "CareerCenter/internal/delivery/http/profile"
	"CareerCenter/internal/repository/account"
	"CareerCenter/internal/repository/jobs"
	"CareerCenter/internal/repository/profile"
	account2 "CareerCenter/internal/usecase/account"
	jobs2 "CareerCenter/internal/usecase/jobs"
	profile2 "CareerCenter/internal/usecase/profile"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
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
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", ParamHandlerWithoutInput).Methods(http.MethodGet)

	r.HandleFunc("/v1/register", handlerAccount.Register).Methods(http.MethodPost)
	r.HandleFunc("/v1/login", handlerAccount.Login).Methods(http.MethodPost)

	r.HandleFunc("/v1/list-jobs", handlerJobs.GetListJob).Methods(http.MethodGet)
	r.HandleFunc("/v1/job-detail/{job_id}", handlerJobs.GetJobById).Methods(http.MethodGet)

	r.HandleFunc("/v1/profile/people", handlerProfile.GetProfile).Methods(http.MethodGet)

	http.ListenAndServe(":8080", r)
}

func ParamHandlerWithoutInput(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "SUCCESS OK")
}

package main

import (
	"CareerCenter/internal/config/database"
	http2 "CareerCenter/internal/delivery/http"
	"CareerCenter/internal/repository"
	"CareerCenter/internal/usecase"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	ctx            = context.TODO()
	mysqlConn      = database.InitMysqlDB()
	repoAccount    = repository.NewAccountMysqlInteractor(mysqlConn)
	useCaseAccount = usecase.NewAccountUsecase(repoAccount)
)

func main() {
	r := mux.NewRouter()

	handlerAccount := http2.NewUseCaseHandler(useCaseAccount)

	r.HandleFunc("/", ParamHandlerWithoutInput).Methods(http.MethodGet)

	r.HandleFunc("/v1/register", handlerAccount.Register).Methods(http.MethodPost)
	r.HandleFunc("/v1/login", handlerAccount.Login).Methods(http.MethodPost)
	http.ListenAndServe(":8080", r)
}

func ParamHandlerWithoutInput(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "SUCCESS OK")
}

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
	ctx             = context.TODO()
	mysqlConn       = database.InitMysqlDB()
	repoRegister    = repository.NewAccountMysqlInteractor(mysqlConn)
	useCaseRegister = usecase.NewAccountUsecase(repoRegister)
)

func main() {
	r := mux.NewRouter()

	handlerRegister := http2.NewUseCaseHandler(useCaseRegister)

	r.HandleFunc("/", ParamHandlerWithoutInput).Methods(http.MethodGet)

	r.HandleFunc("/v1/register", handlerRegister.Register).Methods(http.MethodPost)
	http.ListenAndServe(":8080", r)
}

func ParamHandlerWithoutInput(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "SUCCESS OK")
}

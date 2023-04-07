package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", ParamHandlerWithoutInput).Methods(http.MethodGet)

	r.HandleFunc("/v1/register", handlerAccount.Register).Methods(http.MethodPost)
	r.HandleFunc("/v1/login", handlerAccount.Login).Methods(http.MethodPost)
	r.HandleFunc("/v1/change/password", handlerAccount.ChangePassword).Methods(http.MethodPost)

	r.HandleFunc("/v1/list-jobs", handlerJobs.GetListJob).Methods(http.MethodGet)
	r.HandleFunc("/v1/job-detail/{job_id}", handlerJobs.GetJobById).Methods(http.MethodGet)

	r.HandleFunc("/v1/profile", handlerProfile.GetProfile).Methods(http.MethodGet)
	r.HandleFunc("/v1/profile/update-profile", handlerProfile.UpdateProfile).Methods(http.MethodPost)
	r.HandleFunc("/v1/profile/update-photo", handlerProfile.UpdatePhotoProfile).Methods(http.MethodPost)
	r.HandleFunc("/v1/profile/add-work-experience", handlerProfile.CreateWorkExperience).Methods(http.MethodPost)
	r.HandleFunc("/v1/profile/update-work-experience/{work_experience_id}", handlerProfile.UpdateWorkExperience).Methods(http.MethodPost)

	r.HandleFunc("/v1/list-company", handlerCompany.GetListCompany).Methods(http.MethodGet)
	r.HandleFunc("/v1/company/{company_id}", handlerCompany.GetCompanyById).Methods(http.MethodGet)

	r.HandleFunc("/v1/job-aplication", handlerApplication.SendApplication).Methods(http.MethodPost)

	http.ListenAndServe(":8080", r)
}

func ParamHandlerWithoutInput(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "SUCCESS OK")
}

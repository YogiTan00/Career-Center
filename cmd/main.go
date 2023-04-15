package main

import (
	"CareerCenter/internal/delivery/http/general"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", general.ParamHandlerWithoutInput).Methods(http.MethodGet)
	r.HandleFunc("/v1/photo", general.GetImage).Methods(http.MethodGet)
	r.HandleFunc("/v1/pdf", general.GetPdf).Methods(http.MethodGet)

	r.HandleFunc("/v1/register", handlerAccount.Register).Methods(http.MethodPost)
	r.HandleFunc("/v1/login", handlerAccount.Login).Methods(http.MethodPost)
	r.HandleFunc("/v1/change/password", handlerAccount.ChangePassword).Methods(http.MethodPost)
	r.HandleFunc("/v1/log-out", handlerAccount.Logout).Methods(http.MethodPost)

	r.HandleFunc("/v1/list-jobs", handlerJobs.GetListJob).Methods(http.MethodGet)
	r.HandleFunc("/v1/job-detail/{job_id}", handlerJobs.GetJobById).Methods(http.MethodGet)

	r.HandleFunc("/v1/profile", handlerProfile.GetProfile).Methods(http.MethodGet)
	r.HandleFunc("/v1/profile/update-profile", handlerProfile.UpdateProfile).Methods(http.MethodPost)
	r.HandleFunc("/v1/profile/update-photo", handlerProfile.UpdatePhotoProfile).Methods(http.MethodPost)
	r.HandleFunc("/v1/profile/add-work-experience", handlerProfile.CreateWorkExperience).Methods(http.MethodPost)
	r.HandleFunc("/v1/profile/update-work-experience/{work_experience_id}", handlerProfile.UpdateWorkExperience).Methods(http.MethodPost)
	r.HandleFunc("/v1/profile/deleted-work-experience/{work_experience_id}", handlerProfile.DeletedWorkExperience).Methods(http.MethodPost)
	r.HandleFunc("/v1/profile/add-education/", handlerProfile.CreateEducation).Methods(http.MethodPost)
	r.HandleFunc("/v1/profile/update-education/{education_id}", handlerProfile.UpdateEducation).Methods(http.MethodPost)
	r.HandleFunc("/v1/profile/deleted-education/{education_id}", handlerProfile.DeletedEducation).Methods(http.MethodPost)
	r.HandleFunc("/v1/profile/update-ability", handlerProfile.UpdateAbility).Methods(http.MethodPost)
	r.HandleFunc("/v1/profile/update-language", handlerProfile.UpdateLanguage).Methods(http.MethodPost)
	r.HandleFunc("/v1/profile/update-cv-resume", handlerProfile.UpdateCvResume).Methods(http.MethodPost)
	r.HandleFunc("/v1/profile/update-portofolio", handlerProfile.UpdatePortofolio).Methods(http.MethodPost)

	r.HandleFunc("/v1/list-company", handlerCompany.GetListCompany).Methods(http.MethodGet)
	r.HandleFunc("/v1/company/{company_id}", handlerCompany.GetCompanyById).Methods(http.MethodGet)

	r.HandleFunc("/v1/job-aplication", handlerApplication.SendApplication).Methods(http.MethodPost)

	http.ListenAndServe(":8080", r)
}

package main

import (
	"CareerCenter/internal/config/database"
	account3 "CareerCenter/internal/delivery/http/account"
	application3 "CareerCenter/internal/delivery/http/application"
	company3 "CareerCenter/internal/delivery/http/company"
	handlerGenral "CareerCenter/internal/delivery/http/handler"
	jobs3 "CareerCenter/internal/delivery/http/jobs"
	profile3 "CareerCenter/internal/delivery/http/profile"
	"CareerCenter/internal/repository/account"
	"CareerCenter/internal/repository/application"
	"CareerCenter/internal/repository/company"
	"CareerCenter/internal/repository/jobs"
	"CareerCenter/internal/repository/profile"
	account2 "CareerCenter/internal/usecase/account"
	application2 "CareerCenter/internal/usecase/application"
	company2 "CareerCenter/internal/usecase/company"
	jobs2 "CareerCenter/internal/usecase/jobs"
	profile2 "CareerCenter/internal/usecase/profile"
	"CareerCenter/package/cfg"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	config = cfg.Config{
		ADDRESS:              goDotEnvVariable("ADDRESS"),
		DB_HOST:              goDotEnvVariable("DB_HOST"),
		DB_PORT:              goDotEnvVariable("DB_PORT"),
		DB_USER:              goDotEnvVariable("DB_USER"),
		DB_PASS:              goDotEnvVariable("DB_PASS"),
		DB_NAME:              goDotEnvVariable("DB_NAME"),
		CONFIG_SMTP_HOST:     goDotEnvVariable("CONFIG_SMTP_HOST"),
		CONFIG_SMTP_PORT:     goDotEnvVariable("CONFIG_SMTP_PORT"),
		CONFIG_SENDER_NAME:   goDotEnvVariable("CONFIG_SENDER_NAME"),
		CONFIG_AUTH_EMAIL:    goDotEnvVariable("CONFIG_AUTH_EMAIL"),
		CONFIG_AUTH_PASSWORD: goDotEnvVariable("CONFIG_AUTH_PASSWORD"),
	}
	mysqlConn = database.InitMysqlDB(config)

	repoAccount    = account.NewAccountMysqlInteractor(mysqlConn)
	useCaseAccount = account2.NewAccountUsecase(repoAccount, repoProfile)
	handlerAccount = account3.NewUseCaseAccountHandler(useCaseAccount)

	repoJobs    = jobs.NewJobsMysqlInteractor(mysqlConn)
	useCaseJobs = jobs2.NewJobsUsecase(repoJobs, repoApplication)
	handlerJobs = jobs3.NewUseCaseJobsHandler(useCaseJobs)

	repoProfile    = profile.NewProfileMysqlInteractor(mysqlConn)
	useCaseProfile = profile2.NewProfileUsecase(repoProfile)
	handlerProfile = profile3.NewUseCaseProfileHandler(useCaseProfile)

	repoApplication    = application.NewApplicationMysqlInteractor(mysqlConn)
	useCaseApplication = application2.NewApplicationUsecase(repoApplication, repoProfile)
	handlerApplication = application3.NewUseCaseApplicationHandler(useCaseApplication)

	repoCompany    = company.NewCompanyMysqlInteractor(mysqlConn)
	useCaseCompany = company2.NewCompanyUsecase(repoCompany, repoJobs)
	handlerCompany = company3.NewUseCaseCompanyHandler(useCaseCompany)
)

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	r := mux.NewRouter()
	//Handler General
	r.HandleFunc("/", handlerGenral.ParamHandlerWithoutInput).Methods(http.MethodGet)
	r.HandleFunc("/v1/photo", handlerGenral.GetImage).Methods(http.MethodGet)
	r.HandleFunc("/v1/pdf", handlerGenral.GetPdf).Methods(http.MethodGet)
	//Handler Account
	r.HandleFunc("/v1/register", handlerAccount.Register).Methods(http.MethodPost)
	r.HandleFunc("/v1/login", handlerAccount.Login).Methods(http.MethodPost)
	r.HandleFunc("/v1/change/password", handlerAccount.ChangePassword).Methods(http.MethodPost)
	r.HandleFunc("/v1/logout", handlerAccount.Logout).Methods(http.MethodPost)
	r.HandleFunc("/v1/forget-password", handlerAccount.ForgetPassword).Methods(http.MethodPost)
	//Handler Jobs
	r.HandleFunc("/v1/list-jobs", handlerJobs.GetListJob).Methods(http.MethodGet)
	r.HandleFunc("/v1/job-detail/{job_id}", handlerJobs.GetJobById).Methods(http.MethodGet)
	r.HandleFunc("/v1/job-aplication", handlerApplication.SendApplication).Methods(http.MethodPost)
	//Handler Profile
	r.HandleFunc("/v1/profile", handlerProfile.GetProfile).Methods(http.MethodGet)
	r.HandleFunc("/v1/profile/update-profile", handlerProfile.UpdateProfile).Methods(http.MethodPost)
	r.HandleFunc("/v1/profile/update-photo", handlerProfile.UpdatePhotoProfile).Methods(http.MethodPost)
	r.HandleFunc("/v1/profile/add-work-experience", handlerProfile.CreateWorkExperience).Methods(http.MethodPost)
	r.HandleFunc("/v1/profile/update-work-experience/{work_experience_id}", handlerProfile.UpdateWorkExperience).Methods(http.MethodPost)
	r.HandleFunc("/v1/profile/deleted-work-experience/{work_experience_id}", handlerProfile.DeletedWorkExperience).Methods(http.MethodPost)
	r.HandleFunc("/v1/profile/add-education", handlerProfile.CreateEducation).Methods(http.MethodPost)
	r.HandleFunc("/v1/profile/update-education/{education_id}", handlerProfile.UpdateEducation).Methods(http.MethodPost)
	r.HandleFunc("/v1/profile/deleted-education/{education_id}", handlerProfile.DeletedEducation).Methods(http.MethodPost)
	r.HandleFunc("/v1/profile/update-ability", handlerProfile.UpdateAbility).Methods(http.MethodPost)
	r.HandleFunc("/v1/profile/update-language", handlerProfile.UpdateLanguage).Methods(http.MethodPost)
	r.HandleFunc("/v1/profile/update-cv-resume", handlerProfile.UpdateCvResume).Methods(http.MethodPost)
	r.HandleFunc("/v1/profile/update-portofolio", handlerProfile.UpdatePortofolio).Methods(http.MethodPost)
	//Handler Company
	r.HandleFunc("/v1/list-company", handlerCompany.GetListCompany).Methods(http.MethodGet)
	r.HandleFunc("/v1/company/{company_id}", handlerCompany.GetCompanyById).Methods(http.MethodGet)
	//Handler Admin
	r.HandleFunc("/v1/admin/change-role", handlerAccount.ChangeRoleByAdmin).Methods(http.MethodPut)

	fmt.Println("Career Center Running....")
	originsOk := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Accept=Language", "Authorization", "X-Requested-With", "Ciam-Type",
		"X-Device", "X-App-Version", "Channel", "Device-Brand", "promo-config"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	credsOk := handlers.AllowCredentials()

	err := http.ListenAndServe(":9091", handlers.CORS(originsOk, headersOk, methodsOk, credsOk)(r))
	if err != nil {
		return
	}
}

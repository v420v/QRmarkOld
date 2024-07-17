package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/v420v/qrmarkapi/api/middlewares"
	"github.com/v420v/qrmarkapi/controllers"

	"github.com/v420v/qrmarkapi/services"
)

func NewRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	s := services.NewQrmarkAPIService(db)

	// Middlewares
	m := middlewares.NewMiddleware(s)

	// Controllers
	qrmarkController := controllers.NewQrmarkController(s)
	userController := controllers.NewUserController(s)
	schoolController := controllers.NewSchoolController(s)
	companyController := controllers.NewCompanyController(s)

	r.Use(m.LoggingMiddleware)

	r.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	}).Methods(http.MethodGet, http.MethodOptions)

	isAdmin := r.PathPrefix("/").Subrouter()
	{
		isAdmin.Use(m.AdminMiddleware)
		isAdmin.HandleFunc("/users", userController.GetUserListHandler).Methods(http.MethodGet, http.MethodOptions)
	}

	isAuthenticated := r.PathPrefix("/").Subrouter()
	{
		isAuthenticated.Use(m.AuthMiddleware)
		isAuthenticated.HandleFunc("/qrmarks", qrmarkController.GetQrmarkListHandler).Methods(http.MethodGet, http.MethodOptions)
		isAuthenticated.HandleFunc("/qrmarks", qrmarkController.PostQrmarkHandler).Methods(http.MethodPost, http.MethodOptions)
		isAuthenticated.HandleFunc("/users/current", userController.GetCurrentUserHandler).Methods(http.MethodGet, http.MethodOptions)
	}

	r.HandleFunc("/users/verify/{token:[0-9a-z]+}", userController.VerifyHandler).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/users/{id:[0-9|a-z]+}/points/total", qrmarkController.GetUserTotalPointsHandler).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/users", userController.PostUserHandler).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/users/{id:[0-9|a-z]+}", userController.GetUserDetailHandler).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/schools/{id:[0-9]+}/points", qrmarkController.GetSchoolPointsHandler).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/schools/{id:[0-9]+}", schoolController.GetSchoolDetailHandler).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/schools", schoolController.GetSchoolListHandler).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/schools/search", schoolController.GetSearchSchoolHandler).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/login", userController.LoginHandler).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/logout", userController.LogoutHandler).Methods(http.MethodDelete, http.MethodOptions)

	r.HandleFunc("/companys/list", companyController.GetCompanyListHandler).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/companys/{id:[0-9]+}", companyController.GetCompanyDetailHandler).Methods(http.MethodGet, http.MethodOptions)

	return r
}

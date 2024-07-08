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
		isAdmin.HandleFunc("/qrmark/list", qrmarkController.SelectQrmarkListHandler).Methods(http.MethodGet, http.MethodOptions)
		isAdmin.HandleFunc("/user/list", userController.SelectUserListHandler).Methods(http.MethodGet, http.MethodOptions)
	}

	isAuthenticated := r.PathPrefix("/").Subrouter()
	{
		isAuthenticated.Use(m.AuthMiddleware)
		isAuthenticated.HandleFunc("/user/{id:[0-9]+}/qrmark/list", qrmarkController.SelectUserQrmarkListHandler).Methods(http.MethodGet, http.MethodOptions)
		isAuthenticated.HandleFunc("/user/current", userController.CurrentUserHandler).Methods(http.MethodGet, http.MethodOptions)
		isAuthenticated.HandleFunc("/qrmark", qrmarkController.QrmarkHandler).Methods(http.MethodPost, http.MethodOptions)
	}

	// user
	r.HandleFunc("/verify/{token:[0-9a-z]+}", userController.VerifyHandler).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/user/{id:[0-9|a-z]+}/total_points", qrmarkController.SelectUserTotalPointsHandler).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/user", userController.InsertUserHandler).Methods(http.MethodPost, http.MethodOptions)

	// School
	r.HandleFunc("/school/{id:[0-9|a-z]+}/total_points", qrmarkController.SelectSchoolTotalPointsHandler).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/school/{id:[0-9]+}/points", qrmarkController.SelectSchoolPointsHandler).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/school/{id:[0-9]+}", schoolController.SelectSchoolDetailHandler).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/school/list", schoolController.SelectSchoolListHandler).Methods(http.MethodGet, http.MethodOptions)
	r.HandleFunc("/school/search", schoolController.SearchSchoolHandler).Methods(http.MethodGet, http.MethodOptions)

	r.HandleFunc("/login", userController.LoginHandler).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/logout", userController.LogoutHandler).Methods(http.MethodDelete, http.MethodOptions)
	r.HandleFunc("/company/list", companyController.SelectCompanyListHandler).Methods(http.MethodGet, http.MethodOptions)

	return r
}

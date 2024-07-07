package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/v420v/qrmarkapi/apierrors"
	"github.com/v420v/qrmarkapi/controllers/services"
)

type CompanyController struct {
	service services.CompanyServicer
}

func NewCompanyController(service services.CompanyServicer) *CompanyController {
	return &CompanyController{
		service: service,
	}
}

func (c *CompanyController) SelectCompanyDetailHandler(w http.ResponseWriter, req *http.Request) {
	schoolID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		apierrors.ErrorHandler(w, req, err)
		return
	}

	school, err := c.service.SelectCompanyDetailService(schoolID)
	if err != nil {
		apierrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(school)
}

func (c *CompanyController) SelectCompanyListHandler(w http.ResponseWriter, req *http.Request) {
	page := 0
	queryMap := req.URL.Query()

	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			err = apierrors.BadParam.Wrap(err, "page in query param must be number")
			apierrors.ErrorHandler(w, req, err)
			return
		}
	} else {
		page = 1
	}

	schoolList, err := c.service.SelectCompanyListService(page)
	if err != nil {
		apierrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(schoolList)
}

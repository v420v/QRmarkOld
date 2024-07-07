package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/v420v/qrmarkapi/apierrors"
	"github.com/v420v/qrmarkapi/controllers/services"
	"github.com/v420v/qrmarkapi/models"
)

type SchoolController struct {
	service services.SchoolServicer
}

func NewSchoolController(service services.SchoolServicer) *SchoolController {
	return &SchoolController{
		service: service,
	}
}

func (c *SchoolController) SelectSchoolDetailHandler(w http.ResponseWriter, req *http.Request) {
	schoolID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		apierrors.ErrorHandler(w, req, err)
		return
	}

	school, err := c.service.SelectSchoolDetailService(schoolID)
	if err != nil {
		apierrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(school)
}

func (c *SchoolController) SearchSchoolHandler(w http.ResponseWriter, req *http.Request) {
	var page int = 0
	var q string = ""

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

	if p, ok := queryMap["q"]; ok && len(p) > 0 {
		q = p[0]
	}

	if q == "" {
		json.NewEncoder(w).Encode(models.SchoolList{SchoolList: []models.School{}, HasNext: false, Page: page})
		return
	}

	schoolList, err := c.service.SearchSchoolService(q, page)
	if err != nil {
		apierrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(schoolList)
}

func (c *SchoolController) SelectSchoolListHandler(w http.ResponseWriter, req *http.Request) {
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

	schoolList, err := c.service.SelectSchoolListService(page)
	if err != nil {
		apierrors.ErrorHandler(w, req, err)
		return
	}

	json.NewEncoder(w).Encode(schoolList)
}

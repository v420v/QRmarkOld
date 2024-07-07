package services

import (
	"database/sql"
	"errors"

	"github.com/v420v/qrmarkapi/apierrors"
	"github.com/v420v/qrmarkapi/models"
	"github.com/v420v/qrmarkapi/repositories"
)

func (s *QrmarkAPIService) SelectSchoolDetailService(schoolID int) (models.School, error) {
	school, err := repositories.SelectSchoolDetail(s.DB, schoolID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apierrors.NAData.Wrap(err, "no data")
			return models.School{}, err
		}
		err = apierrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.School{}, err
	}

	return school, nil
}

func (s *QrmarkAPIService) SearchSchoolService(q string, page int) (models.SchoolList, error) {
	schoolList, hasNext, err := repositories.SearchSchool(s.DB, q, page)
	if err != nil {
		err = apierrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.SchoolList{}, err
	}

	return models.SchoolList{SchoolList: schoolList, HasNext: hasNext, Page: page}, nil

}

func (s *QrmarkAPIService) SelectSchoolListService(page int) (models.SchoolList, error) {
	schoolList, hasNext, err := repositories.SelectSchoolList(s.DB, page)
	if err != nil {
		err = apierrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.SchoolList{}, err
	}

	return models.SchoolList{SchoolList: schoolList, HasNext: hasNext, Page: page}, nil
}

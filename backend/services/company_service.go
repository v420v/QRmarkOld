package services

import (
	"database/sql"
	"errors"

	"github.com/v420v/qrmarkapi/apierrors"
	"github.com/v420v/qrmarkapi/models"
	"github.com/v420v/qrmarkapi/repositories"
)

func (s *QrmarkAPIService) SelectCompanyDetailService(schoolID int) (models.Company, error) {
	school, err := repositories.SelectCompanyDetail(s.DB, schoolID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apierrors.NAData.Wrap(err, "no data")
			return models.Company{}, err
		}
		err = apierrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.Company{}, err
	}

	return school, nil
}

func (s *QrmarkAPIService) SelectCompanyListService(page int) ([]models.Company, error) {
	schoolList, err := repositories.SelectCompanyList(s.DB, page)
	if err != nil {
		err = apierrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	return schoolList, nil
}

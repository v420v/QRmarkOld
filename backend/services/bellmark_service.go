package services

import (
	"database/sql"
	"errors"
	"time"

	"github.com/v420v/qrmarkapi/apierrors"
	"github.com/v420v/qrmarkapi/models"
	"github.com/v420v/qrmarkapi/repositories"
)

func (s *QrmarkAPIService) SelectUserQrmarkListService(userID int, page int) (models.QrmarkList, error) {
	qrmarkList, hasNext, err := repositories.SelectUserQrmarkList(s.DB, userID, page)
	if err != nil {
		err = apierrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.QrmarkList{}, err
	}

	return models.QrmarkList{QrmarkList: qrmarkList, HasNext: hasNext, Page: page}, nil
}

func (s *QrmarkAPIService) SelectQrmarkListService(page int) (models.QrmarkList, error) {
	qrmarkList, hasNext, err := repositories.SelectQrmarkList(s.DB, page)
	if err != nil {
		err = apierrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.QrmarkList{}, err
	}

	return models.QrmarkList{QrmarkList: qrmarkList, HasNext: hasNext, Page: page}, nil
}

func (s *QrmarkAPIService) SelectSchoolTotalPointsService(schoolID int) (models.TotalPoints, error) {
	totalPoints, err := repositories.SelectSchoolTotalPoints(s.DB, schoolID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.TotalPoints{Points: 0}, err
		}
		err = apierrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.TotalPoints{}, err
	}

	return totalPoints, nil
}

func (s *QrmarkAPIService) SelectUserTotalPointsService(userID int) (models.TotalPoints, error) {
	totalPoints, err := repositories.SelectUserTotalPoints(s.DB, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.TotalPoints{Points: 0}, err
		}
		err = apierrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.TotalPoints{}, err
	}

	return totalPoints, nil
}

func (s *QrmarkAPIService) SelectSchoolPointsService(schoolID int, year int, month int) ([]models.StaticPoint, error) {
	year_month_date := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
	schoolPoints, err := repositories.SelectSchoolPoints(s.DB, schoolID, year_month_date)
	if err != nil {
		err = apierrors.GetDataFailed.Wrap(err, "fail to get data")
		return nil, err
	}

	return schoolPoints, nil
}

func (s *QrmarkAPIService) UseQrmarkService(qrmarkInfo models.QrmarkInfo) error {
	err := repositories.UseQrmark(s.DB, qrmarkInfo)
	if err != nil {
		err = apierrors.InsertDataFailed.Wrap(err, "fail to insert data")
		return err
	}

	return nil
}

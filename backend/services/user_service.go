package services

import (
	"database/sql"
	"errors"

	"github.com/v420v/qrmarkapi/apierrors"
	"github.com/v420v/qrmarkapi/models"
	"github.com/v420v/qrmarkapi/repositories"
)

func (s *QrmarkAPIService) VerifyUserService(token string) error {
	err := repositories.VerifyUser(s.DB, token)
	if err != nil {
		return err
	}

	return nil
}

func (s *QrmarkAPIService) InsertVerificationTokenService(token models.VerificationToken) error {
	err := repositories.InsertVerificationToken(s.DB, token)
	if err != nil {
		return err
	}
	return nil
}

func (s *QrmarkAPIService) InsertUserService(user models.User) (models.User, error) {
	newUser, err := repositories.InsertUser(s.DB, user)
	if err != nil {
		err = apierrors.InsertDataFailed.Wrap(err, "fail to insert data")
		return models.User{}, err
	}

	return newUser, nil
}

func (s *QrmarkAPIService) SelectUserByIDService(userID int) (models.User, error) {
	user, err := repositories.SelectUserDetail(s.DB, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apierrors.NAData.Wrap(err, "no data")
			return models.User{}, err
		}
		err = apierrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.User{}, err
	}

	return user, nil
}

func (s *QrmarkAPIService) SelectUserByEmailService(email string) (models.User, error) {
	user, err := repositories.SelectUserByEmail(s.DB, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = apierrors.NAData.Wrap(err, "no data")
			return models.User{}, err
		}
		err = apierrors.GetDataFailed.Wrap(err, "fail to get data")
		return models.User{}, err
	}

	return user, nil
}

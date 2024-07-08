package services

import "github.com/v420v/qrmarkapi/models"

type QrmarkServicer interface {
	SelectUserTotalPointsService(userID int) (models.TotalPoints, error)
	SelectSchoolTotalPointsService(userID int) (models.TotalPoints, error)
	SelectSchoolPointsService(schoolID int, year int, month int) ([]models.StaticPoint, error)
	UseQrmarkService(qrmarkInfo models.QrmarkInfo) error
	SelectQrmarkListService(page int) (models.QrmarkList, error)
	SelectUserQrmarkListService(userID int, page int) (models.QrmarkList, error)
}

type UserServicer interface {
	InsertUserService(user models.User) (models.User, error)
	SelectUserByIDService(userID int) (models.User, error)
	VerifyUserService(token string) error
	SelectUserByEmailService(email string) (models.User, error)
	InsertVerificationTokenService(token models.VerificationToken) error
	SelectUserListService(page int) (models.UserList, error)
}

type SchoolServicer interface {
	SelectSchoolListService(page int) (models.SchoolList, error)
	SearchSchoolService(q string, page int) (models.SchoolList, error)
	SelectSchoolDetailService(id int) (models.School, error)
}

type CompanyServicer interface {
	SelectCompanyListService(page int) ([]models.Company, error)
	SelectCompanyDetailService(id int) (models.Company, error)
}

package repositories

import (
	"database/sql"
	"time"

	"github.com/v420v/qrmarkapi/models"
)

func SelectSchoolTotalPoints(db *sql.DB, schoolID int) (models.TotalPoints, error) {
	row := db.QueryRow(`SELECT s.name AS name, SUM(ssp.points) AS total_points FROM school_static_points ssp JOIN school s ON ssp.school_id = s.id WHERE s.id = ? and ssp.created_year_month = ?;`, schoolID)

	if err := row.Err(); err != nil {
		return models.TotalPoints{}, err
	}

	stp := models.TotalPoints{}

	totalPoints := 0

	err := row.Scan(&stp.Points)
	if err != nil {
		return models.TotalPoints{Points: totalPoints}, nil
	}

	return models.TotalPoints{Points: totalPoints}, nil
}

func SelectUserTotalPoints(db *sql.DB, userID int) (models.TotalPoints, error) {
	row := db.QueryRow(`SELECT SUM(points) FROM qrmarks WHERE user_id = ?;`, userID)

	if err := row.Err(); err != nil {
		return models.TotalPoints{}, err
	}

	totalPoints := 0

	err := row.Scan(&totalPoints)
	if err != nil {
		return models.TotalPoints{Points: totalPoints}, nil
	}

	return models.TotalPoints{Points: totalPoints}, nil
}

func SelectSchoolPoints(db *sql.DB, schoolID int, year_month_date time.Time) ([]models.StaticPoint, error) {
	rows, err := db.Query(`SELECT c.company_id, c.name, c.created_at, ssp.points, ssp.created_year_month FROM school_static_points ssp INNER JOIN companys c ON ssp.company_id = c.company_id WHERE ssp.school_id = ? and ssp.created_year_month = ?;`, schoolID, year_month_date)
	if err != nil {
		return []models.StaticPoint{}, err
	}

	list := []models.StaticPoint{}

	for rows.Next() {
		ssp := models.StaticPoint{}
		rows.Scan(&ssp.Company.ID, &ssp.Company.Name, &ssp.Company.CreatedAt, &ssp.Points, &ssp.CreatedYearMonth)
		list = append(list, ssp)
	}

	return list, nil
}

func SelectUserQrmarkList(db *sql.DB, userID int, page int) ([]models.Qrmark, bool, error) {
	sqlStr := `SELECT b.qrmark_id, b.user_id, s.name AS school_name, c.name AS company_name, b.points, b.created_at FROM qrmarks b JOIN schools s ON b.school_id = s.school_id JOIN companys c ON b.company_id = c.company_id where user_id = ? order by created_at desc LIMIT ? OFFSET ?;`

	limit := 5
	hasNext := false

	rows, err := db.Query(sqlStr, userID, limit+1, ((page - 1) * limit))
	if err != nil {
		return nil, hasNext, err
	}

	defer rows.Close()

	QrmarkList := make([]models.Qrmark, 0)
	for rows.Next() {
		var qrmark models.Qrmark
		rows.Scan(
			&qrmark.QrmarkID,
			&qrmark.UserID,
			&qrmark.SchoolName,
			&qrmark.CompanyName,
			&qrmark.Points,
			&qrmark.CreatedAt,
		)
		QrmarkList = append(QrmarkList, qrmark)
	}

	if len(QrmarkList) > limit {
		hasNext = true
		QrmarkList = QrmarkList[:limit]
	}

	return QrmarkList, hasNext, nil
}

func SelectQrmarkList(db *sql.DB, page int) ([]models.Qrmark, bool, error) {
	sqlStr := `SELECT b.qrmark_id, b.user_id, s.name AS school_name, c.name AS company_name, b.points, b.created_at FROM qrmarks b JOIN schools s ON b.school_id = s.school_id JOIN companys c ON b.company_id = c.company_id LIMIT ? OFFSET ?;`

	limit := 5
	hasNext := false

	rows, err := db.Query(sqlStr, limit+1, ((page - 1) * limit))
	if err != nil {
		return nil, hasNext, err
	}

	defer rows.Close()

	QrmarkList := make([]models.Qrmark, 0)
	for rows.Next() {
		var qrmark models.Qrmark
		rows.Scan(
			&qrmark.QrmarkID,
			&qrmark.UserID,
			&qrmark.SchoolName,
			&qrmark.CompanyName,
			&qrmark.Points,
			&qrmark.CreatedAt,
		)
		QrmarkList = append(QrmarkList, qrmark)
	}

	if len(QrmarkList) > limit {
		hasNext = true
		QrmarkList = QrmarkList[:limit]
	}

	return QrmarkList, hasNext, nil
}

func UseQrmark(db *sql.DB, qrmarkInfo models.QrmarkInfo) error {
	now := time.Now()
	year_month := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// school_static_points
	_, err = tx.Exec(
		`insert into school_static_points (school_id, company_id, points, created_year_month) values ((SELECT school_id FROM users WHERE user_id = ?), ?, ?, ?) ON DUPLICATE KEY UPDATE points = points + ?;`,
		qrmarkInfo.UserID, qrmarkInfo.CompanyID, qrmarkInfo.Point, year_month, qrmarkInfo.Point,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	// qrmarks
	_, err = tx.Exec(
		`
		insert into
		qrmarks (qrmark_id, user_id, school_id, company_id, points, created_at)
		values (?, ?, (SELECT school_id FROM users WHERE user_id = ?), ?, ?, ?);
		`,
		qrmarkInfo.QrmarkID, qrmarkInfo.UserID, qrmarkInfo.UserID, qrmarkInfo.CompanyID, qrmarkInfo.Point, now,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

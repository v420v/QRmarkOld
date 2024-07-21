package repositories

import (
	"database/sql"

	"github.com/v420v/qrmarkapi/models"
)

func SelectUserTotalPoints(db *sql.DB, userID int) (models.TotalPoints, error) {
	const sqlStr = `SELECT COALESCE(SUM(points), 0) AS total_points FROM qrmarks WHERE user_id = ?;`

	row := db.QueryRow(sqlStr, userID)

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

func SelectSchoolPoints(db *sql.DB, schoolID int) ([]models.StaticPoint, error) {
	const sql = `
	SELECT
	    combined.company_id,
	    c.name as company_name,
	    c.created_at as company_created_at,
	    sum(combined.total_points) as total_points
	FROM (
	    SELECT
	        company_id,
	        total_points
	    FROM
	        qrmark_snapshots
	    WHERE
	        school_id = ?
	        AND snapshot_date = (SELECT max(snapshot_date) from qrmark_snapshots WHERE school_id = ?)
	    UNION ALL
	    SELECT
	        company_id,
	        points
	    from
	        qrmarks
	    where
	        school_id = ?
	        AND created_at > coalesce((select max(snapshot_date) from qrmark_snapshots WHERE school_id = ?), '1970-01-01')
	) as combined
	INNER JOIN
	    companys c on combined.company_id = c.company_id
	GROUP BY
	    combined.company_id
	ORDER BY
	    combined.company_id;
	`

	rows, err := db.Query(sql, schoolID, schoolID, schoolID, schoolID)

	if err != nil {
		return []models.StaticPoint{}, err
	}

	list := []models.StaticPoint{}

	for rows.Next() {
		ssp := models.StaticPoint{}
		rows.Scan(&ssp.Company.ID, &ssp.Company.Name, &ssp.Company.CreatedAt, &ssp.Points)
		list = append(list, ssp)
	}

	return list, nil
}

func SelectUserQrmarkList(db *sql.DB, userID int, page int) ([]models.Qrmark, bool, error) {
	const sqlStr = `
	SELECT
		b.qrmark_id,
		b.user_id,
		s.name AS school_name,
		c.name AS company_name,
		b.points,
		b.created_at
	FROM
		qrmarks b
	JOIN
		schools s
	ON
		b.school_id = s.school_id
	JOIN
		companys c
	ON
		b.company_id = c.company_id
	WHERE
		user_id = ?
	ORDER BY
		created_at DESC LIMIT ? OFFSET ?;`

	limit := 10
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
	const sqlStr = `
	SELECT
		b.qrmark_id,
		b.user_id,
		s.name AS school_name,
		c.name AS company_name,
		b.points,
		b.created_at
	FROM
		qrmarks b
	JOIN
		schools s
	ON
		b.school_id = s.school_id
	JOIN
		companys c
	ON
		b.company_id = c.company_id
	ORDER BY
		created_at DESC LIMIT ? OFFSET ?;`

	limit := 10
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

func InsertQrmark(db *sql.DB, qrmarkInfo models.QrmarkInfo) error {
	const sqlStr = `INSERT INTO qrmarks (qrmark_id, user_id, school_id, company_id, points, created_at) values (?, ?, (SELECT school_id FROM users WHERE user_id = ?), ?, ?, now());`

	_, err := db.Exec(sqlStr, qrmarkInfo.QrmarkID, qrmarkInfo.UserID, qrmarkInfo.UserID, qrmarkInfo.CompanyID, qrmarkInfo.Point)

	if err != nil {
		return err
	}

	return nil
}

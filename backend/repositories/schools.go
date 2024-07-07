package repositories

import (
	"database/sql"

	"github.com/v420v/qrmarkapi/models"
)

func SelectSchoolDetail(db *sql.DB, schoolID int) (models.School, error) {
	var school models.School

	const sqlStr = `
	select * from schools where school_id = ?;
	`

	row := db.QueryRow(sqlStr, schoolID)

	if err := row.Err(); err != nil {
		return models.School{}, err
	}

	err := row.Scan(
		&school.ID,
		&school.Name,
		&school.CreatedAt,
	)
	if err != nil {
		return models.School{}, err
	}

	return school, nil
}

func SearchSchool(db *sql.DB, q string, page int) ([]models.School, bool, error) {
	const sqlStr = `
	select * from schools where name like CONCAT('%', ?, '%') limit ? offset ?;
	`

	limit := 10
	hasNext := false

	rows, err := db.Query(sqlStr, q, limit+1, ((page - 1) * limit))
	if err != nil {
		return nil, hasNext, err
	}

	defer rows.Close()

	schoolList := make([]models.School, 0)
	for rows.Next() {
		var school models.School
		rows.Scan(
			&school.ID,
			&school.Name,
			&school.CreatedAt,
		)
		schoolList = append(schoolList, school)
	}

	if len(schoolList) > limit {
		hasNext = true
		schoolList = schoolList[:limit]
	}

	return schoolList, hasNext, nil
}

func SelectSchoolList(db *sql.DB, page int) ([]models.School, bool, error) {
	const sqlStr = `
	select * from schools limit ? offset ?;
	`

	limit := 10
	hasNext := false

	rows, err := db.Query(sqlStr, limit+1, ((page - 1) * limit))
	if err != nil {
		return nil, hasNext, err
	}

	defer rows.Close()

	schoolList := make([]models.School, 0)
	for rows.Next() {
		var school models.School
		rows.Scan(
			&school.ID,
			&school.Name,
			&school.CreatedAt,
		)
		schoolList = append(schoolList, school)
	}

	if len(schoolList) > limit {
		hasNext = true
		schoolList = schoolList[:limit]
	}

	return schoolList, hasNext, nil
}

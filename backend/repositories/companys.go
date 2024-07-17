package repositories

import (
	"database/sql"

	"github.com/v420v/qrmarkapi/models"
)

func SelectCompanyDetail(db *sql.DB, companyID int) (models.Company, error) {
	var company models.Company

	const sqlStr = `select * from companys where company_id = ?;`

	row := db.QueryRow(sqlStr, companyID)

	if err := row.Err(); err != nil {
		return models.Company{}, err
	}

	err := row.Scan(
		&company.ID,
		&company.Name,
		&company.CreatedAt,
	)
	if err != nil {
		return models.Company{}, err
	}

	return company, nil
}

func SelectCompanyList(db *sql.DB, page int) ([]models.Company, error) {
	const sqlStr = `select * from companys limit ? offset ?;`

	rows, err := db.Query(sqlStr, 10, ((page - 1) * 10))
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	companyList := make([]models.Company, 0)
	for rows.Next() {
		var company models.Company
		rows.Scan(
			&company.ID,
			&company.Name,
			&company.CreatedAt,
		)
		companyList = append(companyList, company)
	}

	return companyList, nil
}

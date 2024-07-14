package repositories

import (
	"database/sql"
	"errors"
	"time"

	"github.com/v420v/qrmarkapi/apierrors"
	"github.com/v420v/qrmarkapi/models"
)

func InsertVerificationToken(db *sql.DB, verification_token models.VerificationToken) error {
	sqlStr := `insert into verification_tokens (user_id, token, expired_at) values (?, ?, ?);`

	_, err := db.Exec(sqlStr, verification_token.UserID, verification_token.Token, verification_token.ExpiredAt)
	if err != nil {
		return err
	}

	return nil
}

func VerifyUser(db *sql.DB, token string) error {
	sqlStr := `update users
	join verification_tokens on users.user_id = verification_tokens.user_id
	set users.verified = true
	where verification_tokens.token = ?;`

	result, err := db.Exec(sqlStr, token)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return apierrors.NAData.Wrap(errors.New("data not found"), "data not found")
	}

	return nil
}

func InsertUser(db *sql.DB, user models.User) (models.User, error) {
	tx, err := db.Begin()
	if err != nil {
		return models.User{}, err
	}

	deleteUnverifiedUserSql := "delete from users where email = ? and verified = false"

	_, err = tx.Exec(deleteUnverifiedUserSql, user.Email)
	if err != nil {
		tx.Rollback()
		return models.User{}, err
	}

	createdAt := time.Now()

	const insertUserSqlStr = `
	insert into users (name, email, password, role, school_id, verified, created_at) values (?, ?, ?, ?, ?, ?, ?);
	`

	var newUser models.User = models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
		SchoolID: user.SchoolID,
		Verified: user.Verified,
	}

	insertUserResult, err := tx.Exec(insertUserSqlStr, newUser.Name, newUser.Email, newUser.Password, newUser.Role, newUser.SchoolID, newUser.Verified, createdAt)
	if err != nil {
		tx.Rollback()
		return models.User{}, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return models.User{}, err
	}

	id, err := insertUserResult.LastInsertId()
	if err != nil {
		return models.User{}, err
	}

	newUser.ID = int(id)
	newUser.CreatedAt = createdAt

	return newUser, nil
}

func SelectUserByEmail(db *sql.DB, email string) (models.User, error) {
	var user models.User

	const sqlStr = `
	select user_id, name, email, password, role, school_id, verified, created_at from users where email = ?;
	`

	row := db.QueryRow(sqlStr, email)

	if err := row.Err(); err != nil {
		return models.User{}, err
	}

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.SchoolID,
		&user.Verified,
		&user.CreatedAt,
	)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func SelectUserDetail(db *sql.DB, userID int) (models.UserRes, error) {
	var user models.UserRes

	const sqlStr = `
	select 
	    u.user_id, 
	    u.name, 
	    u.email, 
	    u.role, 
	    s.school_id,
		s.name,
		s.created_at,
	    u.verified, 
	    u.created_at
	from 
	    users u
	inner join 
	    schools s on u.school_id = s.school_id
	where 
	    u.user_id = ?;
	`

	row := db.QueryRow(sqlStr, userID)

	if err := row.Err(); err != nil {
		return models.UserRes{}, err
	}

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Role,
		&user.School.ID,
		&user.School.Name,
		&user.School.CreatedAt,
		&user.Verified,
		&user.CreatedAt,
	)
	if err != nil {
		return models.UserRes{}, err
	}

	return user, nil
}

func SelectUserList(db *sql.DB, page int) ([]models.UserRes, bool, error) {
	const sqlStr = `
	select users.user_id, users.name, users.email, users.role, users.verified, schools.school_id, schools.name, schools.created_at, users.created_at from users join schools ON users.school_id = schools.school_id order by users.created_at desc limit ? offset ?;
	`

	limit := 10
	hasNext := false

	rows, err := db.Query(sqlStr, limit+1, ((page - 1) * limit))
	if err != nil {
		return nil, hasNext, err
	}

	defer rows.Close()

	userList := make([]models.UserRes, 0)
	for rows.Next() {
		var user models.UserRes
		rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Role,
			&user.Verified,
			&user.School.ID,
			&user.School.Name,
			&user.School.CreatedAt,
			&user.CreatedAt,
		)
		userList = append(userList, user)
	}

	if len(userList) > limit {
		hasNext = true
		userList = userList[:limit]
	}

	return userList, hasNext, nil
}

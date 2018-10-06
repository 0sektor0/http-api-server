package restapi

import (
	"database/sql"
	m "projects/http-api-server/models"
)

func ScanUserFromRows(rows *sql.Rows) (*m.User, error) {
	user := new(m.User)
	err := rows.Scan(&user.About, &user.Email, &user.Fullname, &user.Nickname)

	return user, err
}

package restapi

import (
	"database/sql"
	"log"
	"net/http"

	m "projects/http-api-server/models"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
)

type UsersStorage struct {
	db *sql.DB
}

func (s *UsersStorage) AddUser(nickname string, user *m.User) *ApiResponse { //(*m.User, []*m.User, *m.Error) {
	_, err := s.db.Exec("INSERT INTO fuser (about, email, ci_email, fullname, nickname, ci_nickname) VALUES($1, $2, LOWER($2), $3, $4, LOWER($4))",
		user.About,
		user.Email,
		user.Fullname,
		nickname)

	//пользователь успешно добавлен
	if err == nil {
		user.Nickname = nickname
		log.Println(user)
		return &ApiResponse{Code: http.StatusCreated, Response: user}
	}

	rows, err := s.db.Query(`SELECT id, about, email, fullname, nickname 
		FROM fuser 
		WHERE ci_nickname=LOWER($1) OR ci_email=LOWER($2)`,
		nickname,
		user.Email)

	if err != nil {
		log.Fatalln(err)
		return &ApiResponse{Code: http.StatusInternalServerError, Response: err}
	}

	existingUsers := make([]*m.User, 0)
	for rows.Next() {
		user, err := ScanUserFromRow(rows)

		if err != nil {
			log.Fatalln(err)
			return &ApiResponse{Code: http.StatusInternalServerError, Response: err}
		}

		existingUsers = append(existingUsers, user)
	}

	if err = rows.Err(); err != nil {
		log.Fatalln(err)
		return &ApiResponse{Code: http.StatusInternalServerError, Response: err}
	}

	log.Println(existingUsers)
	return &ApiResponse{Code: http.StatusConflict, Response: existingUsers}
}

func (s *UsersStorage) UpdateUser(nickname string, update *m.User) *ApiResponse { //(*m.User, *m.Error) {
	row := s.db.QueryRow (`UPDATE fuser
	SET about=$1, email=$2, fullname=$3
	WHERE ci_nickname=LOWER($4)
	REtuRNING id, about, email, fullname, nickname`,
		update.About,
		update.Email,
		update.Fullname,
		nickname)

	user, err := ScanUserFromRow(row)
	if err != nil {
		return &ApiResponse{Code: http.StatusConflict, Response: err}
	}

	return &ApiResponse{Code: http.StatusOK, Response: user}
}

func (s *UsersStorage) GetUserDetails(nickname string) *ApiResponse { //(*m.User, *m.Error) {
	row := s.db.QueryRow(`SELECT id, about, email, fullname, nickname 
	FROM fuser 
	WHERE ci_nickname=LOWER($1)`,
		nickname)

	user, err := ScanUserFromRow(row)
	if err != nil {
		return &ApiResponse{Code: http.StatusNotFound, Response: err}
	}

	return &ApiResponse{Code: http.StatusOK, Response: user}
}

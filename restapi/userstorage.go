package restapi

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	m "projects/http-api-server/models"
)

type UsersStorage struct {
	db *sql.DB
}

func (s *UsersStorage) AddUser(user *m.User) *ApiResponse { //(*m.User, []*m.User, *m.Error) {
	_, err := s.db.Exec("INSERT INTO fuser (about, email, ci_email, fullname, nickname, ci_nickname) VALUES($1, $2, LOWER($2), $3, $4, LOWER($4))",
		user.About,
		user.Email,
		user.Fullname,
		user.Nickname)

	//пользователь успешно добавлен
	if err == nil {
		log.Println(user)
		return &ApiResponse{Code: http.StatusCreated, Response: user}
	}

	rows, err := s.db.Query(`SELECT id, about, email, fullname, nickname 
		FROM fuser 
		WHERE ci_nickname=LOWER($1) OR ci_email=LOWER($2)`,
		user.Nickname,
		user.Email)

	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusInternalServerError, Response: err}
	}

	existingUsers := make([]*m.User, 0)
	for rows.Next() {
		user, err := ScanUserFromRow(rows)

		if err != nil {
			log.Println(err)
			return &ApiResponse{Code: http.StatusInternalServerError, Response: err}
		}

		existingUsers = append(existingUsers, user)
	}

	if err = rows.Err(); err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusInternalServerError, Response: err}
	}

	log.Println(existingUsers)
	return &ApiResponse{Code: http.StatusConflict, Response: existingUsers}
}

//очень плохой код
func (s *UsersStorage) UpdateUser(update *m.User) *ApiResponse { //(*m.User, *m.Error) {
	row := s.db.QueryRow(`
	SELECT id, about, email, fullname, nickname 
	FROM fuser 
	WHERE ci_nickname=LOWER($1)`,
		update.Nickname)

	oldUser, err := ScanUserFromRow(row)
	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusNotFound, Response: err}
	}

	if update.About == "" {
		update.About = oldUser.About
	}
	if update.Email == "" {
		update.Email = oldUser.Email
	}
	if update.Fullname == "" {
		update.Fullname = oldUser.Fullname
	}

	_, err = s.db.Exec(`UPDATE fuser
	SET about=$1, email=$2, ci_email=LOWER($2), fullname=$3
	WHERE ci_nickname=LOWER($4)`,
		update.About,
		update.Email,
		update.Fullname,
		update.Nickname)

	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusConflict, Response: err}
	}

	return &ApiResponse{Code: http.StatusOK, Response: update}
}

func (s *UsersStorage) GetUserDetails(nickname string) *ApiResponse { //(*m.User, *m.Error) {
	row := s.db.QueryRow(`SELECT id, about, email, fullname, nickname 
	FROM fuser 
	WHERE ci_nickname=LOWER($1)`,
		nickname)

	user, err := ScanUserFromRow(row)
	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusNotFound, Response: err}
	}

	return &ApiResponse{Code: http.StatusOK, Response: user}
}

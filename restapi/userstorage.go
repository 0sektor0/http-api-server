package restapi

import (
	"database/sql"
	"log"

	m "projects/http-api-server/models"

	_ "github.com/mattn/go-sqlite3"
)

type UsersStorage struct {
	db *sql.DB
}

func (s *UsersStorage) AddUser(nickname string, user *m.User) *ApiResponse { //(*m.User, []*m.User, *m.Error) {
	_, err := s.db.Exec("INSERT INTO user (about, email, fullname, nickname) VALUES($1, $2, $3, $4)",
		user.About,
		user.Email,
		user.Fullname,
		nickname)

	if err == nil {
		user.Nickname = nickname
		log.Println(user)
		return &ApiResponse{Code: 201, Response: user}
	}

	rows, err := s.db.Query(`SELECT about, email, fullname, nickname 
		FROM user 
		WHERE nickname=$1 OR email=$2`,
		nickname,
		user.Email)

	if err != nil {
		log.Fatalln(err)
		return &ApiResponse{Code: 500, Response: err}
	}

	existingUsers := make([]*m.User, 0)
	for rows.Next() {
		user, err := m.ScanUserFromRows(rows)

		if err != nil {
			log.Fatalln(err)
			return &ApiResponse{Code: 500, Response: err}
		}

		existingUsers = append(existingUsers, user)
	}

	if err = rows.Err(); err != nil {
		log.Fatalln(err)
		return &ApiResponse{Code: 500, Response: err}
	}

	log.Println(existingUsers)
	return &ApiResponse{Code: 409, Response: existingUsers}
}

func (s *UsersStorage) UpdateUser(nickname string, update *m.UserUpdate) *ApiResponse { //(*m.User, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *UsersStorage) GetUserDetails(nickname string) *ApiResponse { //(*m.User, *m.Error) {
	panic("unemplimented function")
	return nil
}

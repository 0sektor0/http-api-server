package restapi

import (
	"bytes"
	"database/sql"
	"fmt"
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

func (s *UsersStorage) GetUserByNickname(nickname string) (*m.User, error) {
	row := s.db.QueryRow(`SELECT id, about, email, fullname, nickname 
	FROM fuser 
	WHERE ci_nickname=LOWER($1)`,
		nickname)

	return ScanUserFromRow(row)
}

func (s *UsersStorage) GetUserByEmail(email string) (*m.User, error) {
	row := s.db.QueryRow(`SELECT id, about, email, fullname, nickname 
	FROM fuser 
	WHERE ci_email=LOWER($1)`,
		email)

	return ScanUserFromRow(row)
}

func (s *UsersStorage) UpdateUser(nickname string, update *m.User) *ApiResponse { //(*m.User, *m.Error) {
	user, err := s.GetUserByNickname(nickname)
	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusNotFound, Response: err}
	}
	
	queryBytes := bytes.Buffer{}
	var queryParams []interface{}

	queryBytes.WriteString(`UPDATE fuser SET id=id`)

	if update.Nickname != "" {
		user.Nickname = nickname
		len := len(queryParams) + 1

		queryBytes.WriteString(fmt.Sprintf(`, nickname=$%v, ci_nickname=LOWER($%v)`, len, len))
		queryParams = append(queryParams, update.Nickname)
	}

	if update.Fullname != "" {
		user.Fullname = update.Fullname
		queryBytes.WriteString(fmt.Sprintf(`, fullname=$%v`, len(queryParams)+1))
		queryParams = append(queryParams, update.Fullname)
	}

	if update.Email != "" {
		user.Email = update.Email
		len := len(queryParams) + 1

		queryBytes.WriteString(fmt.Sprintf(`, email=$%v, ci_email=LOWER($%v)`, len, len))
		queryParams = append(queryParams, update.Email)
	}

	if update.About != "" {
		user.About = update.About
		queryBytes.WriteString(fmt.Sprintf(`, about=$%v`, len(queryParams)+1))
		queryParams = append(queryParams, update.About)
	}

	queryParams = append(queryParams, nickname)
	queryBytes.WriteString(fmt.Sprintf(` WHERE ci_nickname=LOWER($%v)`, len(queryParams)))

	query := queryBytes.String()
	log.Println(query)

	_, err = s.db.Exec(query, queryParams...)
	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusConflict, Response: err}
	}

	return &ApiResponse{Code: http.StatusOK, Response: user}
}

func (s *UsersStorage) GetUserDetails(nickname string) *ApiResponse { //(*m.User, *m.Error) {
	user, err := s.GetUserByNickname(nickname)

	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusNotFound, Response: err}
	}

	return &ApiResponse{Code: http.StatusOK, Response: user}
}

func (s *UsersStorage) GetUsersByForum(slug string, limit int, since string, desc bool) *ApiResponse { //([]*m.User, *m.Error) {
	panic("unemplimented function")
	return nil
}

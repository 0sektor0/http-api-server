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

func GetUserByNickname(db *sql.DB, nickname string) (*m.User, error) {
	row := db.QueryRow(`SELECT id, about, email, fullname, nickname 
	FROM fuser 
	WHERE ci_nickname=LOWER($1)`,
		nickname)

	return ScanUserFromRow(row)
}

func GetUserByEmail(db *sql.DB, email string) (*m.User, error) {
	row := db.QueryRow(`SELECT id, about, email, fullname, nickname 
	FROM fuser 
	WHERE ci_email=LOWER($1)`,
		email)

	return ScanUserFromRow(row)
}

func GetUserById(db *sql.DB, id int) (*m.User, error) {
	row := db.QueryRow(`SELECT id, about, email, fullname, nickname 
	FROM fuser 
	WHERE id=$1`,
		id)

	return ScanUserFromRow(row)
}

func (s *UsersStorage) UpdateUser(nickname string, update *m.User) *ApiResponse { //(*m.User, *m.Error) {
	user, err := GetUserByNickname(s.db, nickname)
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
	user, err := GetUserByNickname(s.db, nickname)

	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusNotFound, Response: err}
	}

	return &ApiResponse{Code: http.StatusOK, Response: user}
}

func ReadUsersArray(rows *sql.Rows) ([]*m.User, error) {
	users := make([]*m.User, 0)

	for rows.Next() {
		user, err := ScanUserFromRow(rows)

		if err != nil {
			log.Println(err)
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (s *UsersStorage) GetUsersByForum(slug string, limit int, since string, desc bool) *ApiResponse { //([]*m.User, *m.Error) {
	row := s.db.QueryRow(`SELECT id, 0, slug, 0, title, '' 
	FROM forum 
	WHERE ci_slug=LOWER($1)` ,
		slug,
	)

	forum, err := ScanForumFromRow(row)
	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusNotFound, Response: err}
	}

	queryBytes := bytes.Buffer{}
	var queryParams []interface{}

	queryBytes.WriteString(`WITH f AS (
		SELECT id, admin_id, slug, title
		FROM forum
		WHERE id=$1
	),
	t AS (
		SELECT id, author_id
		FROM thread
		WHERE forum_id=(SELECT id FROM f)
	),
	fu AS (
		SELECT p.user_id AS id
		FROM post AS p
		JOIN t ON t.id=p.thread_id
		UNION (SELECT author_id AS id FROM t)
	)
	SELECT id, about, email, fullname, nickname 
	FROM  fuser
	JOIN fu USING(id)`)
	queryParams = append(queryParams, forum.Id)

	if desc {
		if since != "" {
			queryBytes.WriteString(fmt.Sprintf(` WHERE convert_to(ci_nickname,'SQL_ASCII')<convert_to(LOWER($%v),'SQL_ASCII')`, len(queryParams)+1))
			queryParams = append(queryParams, since)
		}

		queryBytes.WriteString(` ORDER BY convert_to(ci_nickname,'SQL_ASCII') DESC`)
	} else {
		if since != "" {
			queryBytes.WriteString(` WHERE convert_to(ci_nickname,'SQL_ASCII')>convert_to(LOWER($2),'SQL_ASCII')`)
			queryParams = append(queryParams, since)
		}

		queryBytes.WriteString(` ORDER BY convert_to(ci_nickname,'SQL_ASCII') ASC`)
	}

	if limit > 1 {
		queryBytes.WriteString(fmt.Sprintf(` LIMIT $%v`, len(queryParams)+1))
		queryParams = append(queryParams, limit)
	}

	query := queryBytes.String()
	log.Println(query)

	rows, err := s.db.Query(query, queryParams...)
	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusNotFound, Response: err}
	}

	users, err := ReadUsersArray(rows)
	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusInternalServerError, Response: err}
	}

	return &ApiResponse{Code: http.StatusOK, Response: users}
}

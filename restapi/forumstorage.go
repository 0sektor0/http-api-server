package restapi

import (
	"database/sql"
	"net/http"
	m "projects/http-api-server/models"

	pq "github.com/lib/pq"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

const (
	notFoundError = "23502"
)

type ForumsStorage struct {
	db *sql.DB
}

func (s *ForumsStorage) AddForum(forum *m.Forum) *ApiResponse { //(*m.Forum, *m.Error) {}
	_, err := s.db.Exec(`INSERT INTO forum (title, slug, ci_slug, admin_id) 
	VALUES ($1, $2, LOWER($2), (SELECT id FROM fuser WHERE ci_nickname=LOWER($3)))
	RETURNING id`,
		forum.Title,
		forum.Slug,
		forum.User)

	//форум успешно добавлен
	if err == nil {
		return &ApiResponse{Code: http.StatusCreated, Response: forum}
	}
	
	pgErr := err.(*pq.Error)
	if pgErr.Code == notFoundError {
		return &ApiResponse{Code: http.StatusNotFound, Response: err}
	}

	row := s.db.QueryRow(`WITH f AS (
		SELECT id, admin_id, title, slug
		FROM forum 
		WHERE ci_slug=LOWER($1)
	)
	SELECT f.id, 0 AS posts, f.slug, 0 AS threads, f.title, u.nickname
	FROM f
	LEFT JOIN fuser AS u ON u.id=f.admin_id `,
		forum.Slug)

	oldForum, err := ScanForumFromRow(row)
	return &ApiResponse{Code: http.StatusConflict, Response: oldForum}
}

func (s *ForumsStorage) GetForumDetails(slug string) *ApiResponse { //(*m.Forum, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ForumsStorage) GetForumThreads(slug string, limit int, since string, desc bool) *ApiResponse { //([]*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ForumsStorage) GetForumUsers(slug string, limit int, since string, desc bool) *ApiResponse { //([]*m.User, *m.Error) {
	panic("unemplimented function")
	return nil
}

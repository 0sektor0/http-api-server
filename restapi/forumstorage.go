package restapi

import (
	"database/sql"
	pq "github.com/lib/pq"
	"net/http"
	m "projects/http-api-server/models"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type ForumsStorage struct {
	db *sql.DB
}

func (s *ForumsStorage) AddForum(forum *m.Forum) *ApiResponse { //(*m.Forum, *m.Error) {}
	row := s.db.QueryRow(`WITH u AS (
		SELECT id, nickname
		FROM fuser 
		WHERE ci_nickname=LOWER($3)
	)
	INSERT INTO forum (title, slug, ci_slug, admin_id) 
	VALUES ($1, $2, LOWER($2), (SELECT id FROM u))
	RETURNING id, 0, slug, 0, title, (SELECT nickname FROM u)`,
		forum.Title,
		forum.Slug,
		forum.User,
	)

	addedForum, err := ScanForumFromRow(row)
	if err == nil {
		return &ApiResponse{Code: http.StatusCreated, Response: addedForum}
	}

	pgErr := err.(*pq.Error)
	if pgErr.Code == notFoundError {
		return &ApiResponse{Code: http.StatusNotFound, Response: err}
	}

	row = s.db.QueryRow(`WITH f AS (
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
	row := s.db.QueryRow(`SELECT f.id, SUM(coalesce(p.id,0)) AS posts, f.slug, COUNT(DISTINCT t.id) AS threads, f.title, u.nickname
	FROM (
		SELECT * 
		FROM forum
		WHERE ci_slug=LOWER($1)
	) AS f
	LEFT JOIN thread AS t ON t.forum_id=f.id
	LEFT JOIN post AS p ON p.thread_id=t.id
	LEFT JOIN fuser AS u ON u.id=f.admin_id
	GROUP BY f.id, f.admin_id, f.slug, f.title, u.nickname`,
		slug)

	forum, err := ScanForumFromRow(row)
	if err != nil {
		return &ApiResponse{Code: http.StatusNotFound, Response: err}
	}

	return &ApiResponse{Code: http.StatusOK, Response: forum}
}

func (s *ForumsStorage) GetForumThreads(slug string, limit int, since string, desc bool) *ApiResponse { //([]*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ForumsStorage) GetForumUsers(slug string, limit int, since string, desc bool) *ApiResponse { //([]*m.User, *m.Error) {
	panic("unemplimented function")
	return nil
}

package restapi

import (
	"database/sql"
	pq "github.com/lib/pq"
	"log"
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
		log.Println(err)
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
	row := s.db.QueryRow(`WITH f AS (
		SELECT id, admin_id, slug, title
		FROM forum
		WHERE ci_slug=LOWER($1)
	),
	t AS (
		SELECT id
		FROM thread
		WHERE forum_id=(SELECT id FROM f)
	),
	p AS (
		SELECT COUNT(p.id) AS posts
		FROM post AS p
		JOIN t ON t.id=p.thread_id
	)
	
	SELECT f.id, (SELECT posts FROM p), slug, (SELECT COUNT(t.id) FROM t), title, u.nickname 
	FROM f 
	JOIN fuser AS u ON u.id=f.admin_id `,
		slug)

	forum, err := ScanForumFromRow(row)
	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusNotFound, Response: err}
	}

	return &ApiResponse{Code: http.StatusOK, Response: forum}
}
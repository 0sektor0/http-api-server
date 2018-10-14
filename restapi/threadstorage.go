package restapi

import (
	"database/sql"
	pq "github.com/lib/pq"
	"net/http"
	m "projects/http-api-server/models"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type ThreadsStorage struct {
	db *sql.DB
}

func (s *ThreadsStorage) AddThread(thread *m.Thread) *ApiResponse { //*m.Thread, *m.Error) {
	row := s.db.QueryRow(`WITH u AS (
		SELECT id, nickname
		FROM fuser
		WHERE ci_nickname=LOWER($2)
	), f AS (
		SELECT id, slug
		FROM forum 
		WHERE ci_slug=LOWER($3) 
	)
	INSERT INTO thread (id, author_id, forum_id, title, message, slug, ci_slug, created)
	VALUES ($1, (SELECT id FROM u), (SELECT id FROM f), $4, $5, '$6', LOWER($6), $7)
	RETURNING id, slug, title, 0, (SELECT slug FROM f), (SELECT nickname FROM u), created, message`,
		thread.Id,
		thread.Author,
		thread.Forum,
		thread.Title,
		thread.Message,
		thread.Slug,
		thread.Created,
	)

	addedThread, err := ScanThreadFromRow(row)
	if err == nil {
		return &ApiResponse{Code: http.StatusCreated, Response: addedThread}
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
		thread.Slug)

	oldThread, err := ScanForumFromRow(row)
	return &ApiResponse{Code: http.StatusConflict, Response: oldThread}	
}

func (s *ThreadsStorage) GetThreadDetails(slug string) *ApiResponse { //(*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ThreadsStorage) GetThreadPosts(slug string, limit int, since int, sort string, desc bool) *ApiResponse { //([]*m.Post, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ThreadsStorage) UpdateThread(thread *m.Thread) *ApiResponse { //(*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ThreadsStorage) VoteForThread(slug string, vote *m.Vote) *ApiResponse { //(*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil
}

package restapi

import (
	"time"
	"log"
	"database/sql"
	"fmt"
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
	if thread.Created == "" {
		thread.Created = fmt.Sprintf("%v", time.Now().Format(time.RFC3339)) 
	}
	
	row := s.db.QueryRow(`WITH u AS (
		SELECT id, nickname
		FROM fuser
		WHERE ci_nickname=LOWER($1)
	), f AS (
		SELECT id, slug
		FROM forum 
		WHERE ci_slug=LOWER($2) 
	)
	INSERT INTO thread (author_id, forum_id, title, message, slug, ci_slug, created)
	VALUES ((SELECT id FROM u), (SELECT id FROM f), $3, $4, $5, LOWER($5), $6)
	RETURNING id, slug, title, 0, (SELECT slug FROM f), (SELECT nickname FROM u), created, message`,
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

	log.Println(err)
	pgErr := err.(*pq.Error)
	if pgErr.Code == notFoundError {
		return &ApiResponse{Code: http.StatusNotFound, Response: err}
	}

	row = s.db.QueryRow(`WITH t AS (
		SELECT t.id, t.slug, t.title, f.slug AS forum_slug, u.nickname, t.created, t.message
		FROM thread AS t
		LEFT JOIN forum AS f ON f.id=t.forum_id
		LEFT JOIN fuser AS u ON u.id=t.author_id
		WHERE u.ci_nickname=LOWER($1) AND f.ci_slug=LOWER($2) AND t.title=$3 AND t.message=$4 AND t.created=$5 AND t.ci_slug=LOWER($6)
	)
	SELECT t.id, t.slug, t.title, SUM(coalesce(v.voice, 0)), t.forum_slug, t.nickname, t.created, t.message
	FROM t
	LEFT JOIN vote AS v ON t.id=v.thread_id
	GROUP BY t.id, t.title, t.slug, t.nickname, t.created, t.message`,
		thread.Author,
		thread.Forum,
		thread.Title,
		thread.Message,
		thread.Created,
		thread.Slug,
	)

	oldThread, err := ScanForumFromRow(row)
	log.Println(err)
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

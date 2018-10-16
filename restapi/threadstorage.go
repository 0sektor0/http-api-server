package restapi

import (
	"database/sql"
	"fmt"
	pq "github.com/lib/pq"
	"log"
	"net/http"
	m "projects/http-api-server/models"
	"strconv"
	"time"

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

	pgErr := err.(*pq.Error)
	if pgErr.Code == notFoundError {
		log.Println(err)
		return &ApiResponse{Code: http.StatusNotFound, Response: err}
	}

	row = s.db.QueryRow(`WITH t AS (
		SELECT t.id, t.slug, t.title, f.slug AS forum_slug, u.nickname, t.created, t.message
		FROM thread AS t
		LEFT JOIN forum AS f ON f.id=t.forum_id
		LEFT JOIN fuser AS u ON u.id=t.author_id
		WHERE t.ci_slug=LOWER($1)
	)
	SELECT t.id, t.slug, t.title, SUM(coalesce(v.voice, 0)), t.forum_slug, t.nickname, t.created, t.message
	FROM t
	LEFT JOIN vote AS v ON t.id=v.thread_id
	GROUP BY t.id, t.title, t.slug, t.nickname, t.created, t.message, t.forum_slug`,
		thread.Slug,
	)

	oldThread, err := ScanThreadFromRow(row)
	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusInternalServerError, Response: err}
	}

	return &ApiResponse{Code: http.StatusConflict, Response: oldThread}
}

func (s *ThreadsStorage) GetThreadDetails(slug string) *ApiResponse { //(*m.Thread, *m.Error) {
	threadId, err := strconv.Atoi(slug)
	if err != nil {
		threadId = 0
	}

	row := s.db.QueryRow(`WITH t AS (
		SELECT t.id, t.slug, t.title, f.slug AS forum_slug, u.nickname, t.created, t.message
		FROM thread AS t
		LEFT JOIN forum AS f ON f.id=t.forum_id
		LEFT JOIN fuser AS u ON u.id=t.author_id
		WHERE t.ci_slug=LOWER($1) OR t.id=$2
	)
	SELECT t.id, t.slug, t.title, SUM(coalesce(v.voice, 0)), t.forum_slug, t.nickname, t.created, t.message
	FROM t
	LEFT JOIN vote AS v ON t.id=v.thread_id
	GROUP BY t.id, t.title, t.slug, t.nickname, t.created, t.message, t.forum_slug`,
		slug,
		threadId,
	)

	thread, err := ScanThreadFromRow(row)
	if err != nil {
		return &ApiResponse{Code: http.StatusNotFound, Response: err}
	}

	return &ApiResponse{Code: http.StatusOK, Response: thread}
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
	threadId, err := strconv.Atoi(slug)
	if err != nil {
		threadId = 0
	}

	row := s.db.QueryRow(`SELECT t.id, t.slug, t.title, 0, f.slug, u.nickname, t.created, t.message
		FROM thread AS t
		LEFT JOIN fuser AS u ON u.id=t.author_id
		LEFT JOIN forum AS f ON f.id=t.forum_id
		WHERE t.id=$1 OR t.ci_slug=LOWER($2)`,
		threadId,
		slug,
	)

	thread, err := ScanThreadFromRow(row)
	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusNotFound, Response: err}
	}

	row = s.db.QueryRow(`SELECT id, about, email, fullname, nickname 
	FROM fuser 
	WHERE ci_nickname=LOWER($1)`,
		vote.NickName,
	)

	user, err := ScanUserFromRow(row)
	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusNotFound, Response: err}
	}

	res, err := s.db.Exec(`UPDATE vote 
	SET voice=$1
	WHERE thread_id=$2 AND user_id=$3`,
		vote.Voice,
		thread.Id,
		user.Id,
	)

	rowsAffcted, err := res.RowsAffected()
	if rowsAffcted == 0 || err != nil {
		_, err = s.db.Exec(`INSERT INTO vote (user_id, thread_id, voice)
		VALUES ((SELECT id FROM fuser WHERE ci_nickname=LOWER($1)), $2, $3)
		RETURNING id`,
			vote.NickName,
			thread.Id,
			vote.Voice,
		)

		if err != nil {
			log.Println(err)
			return &ApiResponse{Code: http.StatusNotFound, Response: err}
		}
	}

	row = s.db.QueryRow(`SELECT SUM(coalesce(voice, 0)) AS votes
	FROM vote 
	WHERE thread_id=$1`,
		thread.Id,
	)

	err = row.Scan(&thread.Votes)
	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusInternalServerError, Response: err}
	}

	return &ApiResponse{Code: http.StatusOK, Response: thread}
}

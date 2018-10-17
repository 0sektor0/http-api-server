package restapi

import (
	"bytes"
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

func ReadPostsArray(rows *sql.Rows) ([]*m.Post, error) {
	posts := make([]*m.Post, 0)
	for rows.Next() {
		post, err := ScanPostFromRow(rows)

		if err != nil {
			log.Println(err)
			return posts, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (s *ThreadsStorage) GetThreadPostsFlat(thread *m.Thread, limit int, since int, desc bool) *ApiResponse {
	var queryParams []interface{}
	queryParams = append(queryParams, thread.Id)

	queryBytes := bytes.Buffer{}
	if since != -1 {
		whereFilter := ">"
		if desc {
			whereFilter = "<"
		}

		queryBytes.WriteString(fmt.Sprintf(`WITH p AS (
			SELECT p.id, p.parent_id, p.thread_id, p.message, p.created, p.edited, p.user_id
			FROM post AS p
			WHERE thread_id=$1 AND id%v$`, whereFilter))
		queryBytes.WriteString(strconv.Itoa(len(queryParams) + 1))
		queryParams = append(queryParams, since)
	} else {
		queryBytes.WriteString(`WITH p AS (
			SELECT p.id, p.parent_id, p.thread_id, p.message, p.created, p.edited, p.user_id
			FROM post AS p
			WHERE thread_id=$1`)
	}

	queryBytes.WriteString(fmt.Sprintf(`)
	SELECT p.id, '%v', p.parent_id, p.thread_id, u.nickname, p.message, p.created, p.edited
	FROM p
	LEFT JOIN fuser AS u ON u.id=p.user_id`, thread.Forum))

	if desc {
		queryBytes.WriteString(` ORDER BY (p.id, p.created) DESC`)
	} else {
		queryBytes.WriteString(` ORDER BY (p.id, p.created) ASC`)
	}

	if limit != -1 {
		queryBytes.WriteString(` LIMIT $`)
		queryBytes.WriteString(strconv.Itoa(len(queryParams) + 1))
		queryParams = append(queryParams, limit)
	}

	query := queryBytes.String()
	log.Println(query)
	rows, err := s.db.Query(query, queryParams...)

	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusNotFound, Response: err}
	}

	posts, err := ReadPostsArray(rows)
	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusNotFound, Response: err}
	}

	return &ApiResponse{Code: http.StatusOK, Response: posts}
}

func (s *ThreadsStorage) GetThreadPostsTree(thread *m.Thread, limit int, since int, desc bool) *ApiResponse {
	var queryParams []interface{}
	queryParams = append(queryParams, thread.Id)

	queryBytes := bytes.Buffer{}
	queryBytes.WriteString(fmt.Sprintf(`WITH RECURSIVE tree AS (
		SELECT id, user_id, thread_id, parent_id, message, edited, created, ARRAY[]::INTEGER[] || id AS path, id AS root
		FROM post 
		WHERE parent_id IS NULL AND thread_id=$1
	
		UNION ALL
	
		SELECT p.id, p.user_id, p.thread_id, p.parent_id, p.message, p.edited, p.created, t.path || p.id, t.id
		FROM post AS p, tree AS t
		WHERE p.parent_id = t.id
   	) 
   SELECT t.id, '%v', coalesce(t.parent_id, 0), t.thread_id, u.nickname, t.message, t.created, t.edited
   FROM tree AS t
   LEFT JOIN fuser AS u ON u.id=user_id`, thread.Forum))

	if since != -1 {
		if desc {
			queryBytes.WriteString(` WHERE t.path < (SELECT path FROM tree WHERE id=$`)
			queryBytes.WriteString(strconv.Itoa(len(queryParams) + 1))
			queryBytes.WriteString(`) ORDER BY t.path DESC`)
		} else {
			queryBytes.WriteString(` WHERE t.path > (SELECT path FROM tree WHERE id=$`)
			queryBytes.WriteString(strconv.Itoa(len(queryParams) + 1))
			queryBytes.WriteString(`) ORDER BY t.path ASC`)
		}
		queryParams = append(queryParams, since)
	} else {
		if desc {
			queryBytes.WriteString(` ORDER BY t.path DESC`)
		} else {
			queryBytes.WriteString(` ORDER BY t.path ASC`)
		}
	}

	if limit != -1 {
		queryBytes.WriteString(` LIMIT $`)
		queryBytes.WriteString(strconv.Itoa(len(queryParams) + 1))
		queryParams = append(queryParams, limit)
	}

	query := queryBytes.String()
	log.Println(thread, limit, since, desc, query)

	rows, err := s.db.Query(queryBytes.String(), queryParams...)
	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusInternalServerError, Response: err}
	}

	posts, err := ReadPostsArray(rows)
	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusInternalServerError, Response: err}
	}

	return &ApiResponse{Code: http.StatusOK, Response: posts}
}

func (s *ThreadsStorage) GetThreadPostsParentTree(thread *m.Thread, limit int, since int, desc bool) *ApiResponse {
	var queryParams []interface{}
	queryParams = append(queryParams, thread.Id)
	queryBytes := bytes.Buffer{}

	if since == -1 {
		if limit != -1 {
			queryBytes.WriteString(`WITH RECURSIVE tree AS (
				(SELECT id, user_id, thread_id, parent_id, message, edited, created, ARRAY[]::INTEGER[] || id AS path, id AS root
				FROM post 
			   WHERE parent_id IS NULL AND thread_id=$1
			   ORDER BY id
			   LIMIT $2)`)
			queryParams = append(queryParams, limit)
		} else {
			queryBytes.WriteString(`WITH RECURSIVE tree AS (
				SELECT id, user_id, thread_id, parent_id, message, edited, created, ARRAY[]::INTEGER[] || id AS path, id AS root
				FROM post 
			   WHERE parent_id IS NULL AND thread_id=$1`)
		}
	} else {
		queryBytes.WriteString(`WITH RECURSIVE tree AS (
			SELECT id, user_id, thread_id, parent_id, message, edited, created, ARRAY[]::INTEGER[] || id AS path, id AS root
			FROM post 
		   WHERE parent_id IS NULL AND thread_id=$1`)
	}

	queryBytes.WriteString(fmt.Sprintf(` UNION ALL
	   SELECT p.id, p.user_id, p.thread_id, p.parent_id, p.message, p.edited, p.created, t.path || p.id, t.id
		 FROM post AS p, tree AS t
		 WHERE p.parent_id = t.id
	   ) 
	   SELECT t.id, '%v', coalesce(t.parent_id, 0), t.thread_id, u.nickname, t.message, t.created, t.edited
	   FROM tree AS t
	   LEFT JOIN fuser AS u ON u.id=user_id`, thread.Forum))

	if since != -1 {
		if desc {
			queryBytes.WriteString(` WHERE t.path < (SELECT path FROM tree WHERE id=$`)
			queryBytes.WriteString(strconv.Itoa(len(queryParams) + 1))
			queryBytes.WriteString(`) ORDER BY t.root DESC, t.path ASC`)
		} else {
			queryBytes.WriteString(` WHERE t.path > (SELECT path FROM tree WHERE id=$`)
			queryBytes.WriteString(strconv.Itoa(len(queryParams) + 1))
			queryBytes.WriteString(`) ORDER BY t.path ASC`)
		}

		queryParams = append(queryParams, since)
		if limit != -1 {
			queryBytes.WriteString(` LIMIT $`)
			queryBytes.WriteString(strconv.Itoa(len(queryParams) + 1))
			queryParams = append(queryParams, limit)
		}

	} else {
		if desc {
			queryBytes.WriteString(` ORDER BY t.root DESC, t.path ASC`)
		} else {
			queryBytes.WriteString(` ORDER BY t.path ASC`)
		}
	}

	query := queryBytes.String()
	log.Println(thread, limit, since, desc, query)

	rows, err := s.db.Query(queryBytes.String(), queryParams...)
	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusInternalServerError, Response: err}
	}

	posts, err := ReadPostsArray(rows)
	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusInternalServerError, Response: err}
	}

	return &ApiResponse{Code: http.StatusOK, Response: posts}
}

func (s *ThreadsStorage) GetThreadPosts(slug string, limit int, since int, sort string, desc bool) *ApiResponse { //([]*m.Post, *m.Error) {
	threadId, err := strconv.Atoi(slug)
	if err != nil {
		threadId = 0
	}

	row := s.db.QueryRow(`SELECT t.id, t.slug, t.title, 0, f.slug, '', t.created, t.message
		FROM thread AS t
		LEFT JOIN forum AS f ON f.id=t.forum_id
		WHERE t.ci_slug=LOWER($1) OR t.id=$2`,
		slug,
		threadId,
	)

	thread, err := ScanThreadFromRow(row)
	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusNotFound, Response: err}
	}

	switch sort {
	case "tree":
		return s.GetThreadPostsTree(thread, limit, since, desc)
	case "parent_tree":
		return s.GetThreadPostsParentTree(thread, limit, since, desc)
	default:
		return s.GetThreadPostsFlat(thread, limit, since, desc)
	}

	return &ApiResponse{Code: http.StatusBadRequest, Response: nil}
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

func (s *ThreadsStorage) UpdateThread(thread *m.Thread) *ApiResponse { //(*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil
}

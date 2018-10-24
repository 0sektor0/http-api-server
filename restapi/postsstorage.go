package restapi

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	pq "github.com/lib/pq"
	m "projects/http-api-server/models"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type PostsStorage struct {
	db *sql.DB
}

func (s *PostsStorage) AddPosts(slug string, posts []*m.Post) *ApiResponse { //([]*m.Post, *m.Error) {
	threadId, err := strconv.Atoi(slug)
	if err != nil {
		threadId = 0
	}

	row := s.db.QueryRow(`SELECT t.id, t.slug, t.title, 0, f.slug, '', t.created, t.message
		FROM thread AS t
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

	newPosts := make([]*m.Post, 0)
	timeNow := time.Now().Format(time.RFC3339)

	for _, post := range posts {
		if post.Parent != nil && *post.Parent != 0 {
			parent, err := GetPostDetails(s.db, *post.Parent)

			if err != nil {
				log.Println(err)
				return &ApiResponse{Code: http.StatusConflict, Response: err}
			}

			if parent.ThreadId != thread.Id {
				log.Println(err)
				return &ApiResponse{Code: http.StatusConflict, Response: err}
			}
		}

		row = s.db.QueryRow(`WITH u AS (
			SELECT id, nickname
			FROM fuser
			WHERE ci_nickname=LOWER($1)
		)
		INSERT INTO post (user_id, parent_id, thread_id, message, created)
		VALUES((SELECT id FROM u), $2, $3, $4, $7)
		RETURNING id, $5, coalesce(parent_id, 0), $6, (SELECT nickname from u), message, created, edited `,
			post.Author,
			post.Parent,
			thread.Id,
			post.Message,
			thread.Forum,
			thread.Id,
			timeNow,
		)

		post, err := ScanPostFromRow(row)
		if err != nil {
			log.Println(err)
			pgErr := err.(*pq.Error)

			switch pgErr.Code {
			case notFoundError:
				return &ApiResponse{Code: http.StatusNotFound, Response: err}
			default:
				return &ApiResponse{Code: http.StatusConflict, Response: err}
			}
		}

		newPosts = append(newPosts, post)
	}

	return &ApiResponse{Code: http.StatusCreated, Response: newPosts}
}

func GetPostDetails(db *sql.DB, id int) (*m.PostFull, error) {
	row := db.QueryRow(`WITH p AS (
		SELECT *
		FROM post 
		WHERE id=$1
	)
	SELECT p.id, u.nickname, p.created, f.slug, p.edited, p.message, coalesce(p.parent_id, 0), t.id, u.id, f.id
	FROM p
	JOIN fuser AS u ON u.id=p.user_id
	JOIN thread AS t ON t.id=p.thread_id
	JOIN forum AS f ON f.id=t.forum_id`,
		id,
	)

	return ScanPostDetailsFromRow(row)
}

func (s *PostsStorage) GetPostDetails(id int, relates []string) *ApiResponse { //(*m.PostFull, *m.Error) {
	post, err := GetPostDetails(s.db, id)
	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusNotFound, Response: err}
	}

	details := &m.Details{Post: post}

	for _, related := range relates {
		switch related {
		case "user":
			user, _ := GetUserById(s.db, post.AuthorId)
			details.User = user
		case "thread":
			thread, _ := GetThreadDetails(s.db, strconv.Itoa(post.ThreadId))
			details.Thread = thread
		case "forum":
			forum, _ := GetForumDetails(s.db, post.Forum)
			details.Forum = forum
		default:
			break
		}
	}

	return &ApiResponse{Code: http.StatusOK, Response: details}
}

func (s *PostsStorage) UpdatePost(id int, update *m.PostUpdate) *ApiResponse { //(*m.Post, *m.Error) {
	post, err := GetPostDetails(s.db, id)
	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusNotFound, Response: err}
	}

	queryBytes := bytes.Buffer{}
	var queryParams []interface{}

	queryBytes.WriteString(`UPDATE post SET edited='true'`)

	if update.Message == "" || update.Message == post.Message {
		return &ApiResponse{Code: http.StatusOK, Response: post}
	}

	queryBytes.WriteString(fmt.Sprintf(`, message=$%v`, len(queryParams)+1))
	queryParams = append(queryParams, update.Message)

	queryBytes.WriteString(fmt.Sprintf(` WHERE id=$%v`, len(queryParams)+1))
	queryParams = append(queryParams, id)

	query := queryBytes.String()
	log.Println(query)

	_, err = s.db.Exec(query, queryParams...)
	if err != nil {
		log.Println(err)
		return &ApiResponse{Code: http.StatusBadRequest, Response: err}
	}

	post.Message = update.Message
	post.IsEdited = true

	return &ApiResponse{Code: http.StatusOK, Response: post}
}

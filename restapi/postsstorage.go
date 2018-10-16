package restapi

import (
	"time"
	"log"
	"database/sql"
	"net/http"
	m "projects/http-api-server/models"
	"strconv"

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
		row = s.db.QueryRow(`WITH u AS (
			SELECT id, nickname
			FROM fuser
			WHERE ci_nickname=LOWER($1)
		)
		INSERT INTO post (user_id, parent_id, thread_id, message, created)
		VALUES((SELECT id FROM u), $2, $3, $4, $7)
		RETURNING id, $5, parent_id, $6, (SELECT nickname from u), message, created, edited `,
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
			return &ApiResponse{Code: http.StatusConflict, Response: err}
		}

		newPosts = append(newPosts, post)
	}

	return &ApiResponse{Code: http.StatusCreated, Response: newPosts}
}

func (s *PostsStorage) AddThread(slug string, thread *m.Thread) *ApiResponse { //*m.Thread, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *PostsStorage) GetPostDetails(id int32, related []string) *ApiResponse { //(*m.PostFull, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *PostsStorage) UpdatePost(id int64, update *m.PostUpdate) *ApiResponse { //(*m.Post, *m.Error) {
	panic("unemplimented function")
	return nil
}

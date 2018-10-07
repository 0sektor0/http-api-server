package restapi

import (
	"net/http"
	"database/sql"
	m "projects/http-api-server/models"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
)

type ForumsStorage struct {
	db *sql.DB
}

func (s *ForumsStorage) AddForum(forum *m.Forum) *ApiResponse { //(*m.Forum, *m.Error) {}
	_, err := s.db.Exec(`INSERT INTO forum (title, slug, admin_id) 
	VALUES ($1, $2, (SELECT id FROM fuser WHERE nickname=$3))`,
		forum.Title,
		forum.Slug,
		forum.User)

	//форум успешно добавлен
	if err==nil {
		return &ApiResponse{Code: http.StatusCreated, Response: forum}
	}

	return &ApiResponse{Code: http.StatusOK, Response: new(m.Forum)}
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

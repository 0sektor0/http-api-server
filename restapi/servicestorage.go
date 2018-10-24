package restapi

import (
	"net/http"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
)

type ServiceStorage struct {
	db *sql.DB
}

func (s *ServiceStorage) GetServiceStatus() *ApiResponse { //(*m.Status, *m.Error) {
	row := s.db.QueryRow(`SELECT 
	(SELECT COUNT(*) FROM fuser), 
	(SELECT COUNT(*) FROM thread), 
	(SELECT COUNT(*) FROM post), 
	(SELECT COUNT(*) FROM forum)`)

	status, err := ScanStatusFromRow(row)
	if err != nil {
		return &ApiResponse{Code: http.StatusInternalServerError, Response: err}
	}

	return &ApiResponse{Code: http.StatusOK, Response: status}
}

func (s *ServiceStorage) VipeServiceStatus() *ApiResponse { //*m.Error {
	s.db.Exec(`truncate table vote, post, thread, forum, fuser`)
	return &ApiResponse{Code: http.StatusOK, Response: "OK"}
}

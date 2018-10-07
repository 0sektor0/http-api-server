package restapi

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
)

type ServiceStorage struct {
	db *sql.DB
}

func (s *ServiceStorage) GetServiceStatus() *ApiResponse { //(*m.Status, *m.Error) {
	panic("unemplimented function")
	return nil
}

func (s *ServiceStorage) VipeServiceStatus() *ApiResponse { //*m.Error {
	panic("unemplimented function")
	return nil
}

package restapi

import "database/sql"

type ApiService struct {
	Users   IUsersStorage
	Forums  IForumsStorage
	Threads IThreadsStorage
	Posts   IPostsStorage
	Service IServiceStorege
}

func NewApiService(connector string, connection string) (*ApiService, error) {
	db, err := sql.Open(connector, connection)
	if err != nil {
		return nil, err
	}

	service := &ApiService{
		Users:   &UsersStorage{db},
		Forums:  &ForumsStorage{db},
		Threads: &ThreadsStorage{db},
		Posts:   &PostsStorage{db},
		Service: &ServiceStorage{db},
	}

	return service, nil
}

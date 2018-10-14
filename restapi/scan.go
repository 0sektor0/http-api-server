package restapi

import (
	m "projects/http-api-server/models"
)

type IRow interface {
	Scan(dest ...interface{}) error
}

//SELECT id, about, email, fullname, nickname
func ScanUserFromRow(row IRow) (*m.User, error) {
	user := new(m.User)
	err := row.Scan(&user.Id, &user.About, &user.Email, &user.Fullname, &user.Nickname)

	return user, err
}

//SELECT id, posts, slug, threads, title, userNickname
func ScanForumFromRow(row IRow) (*m.Forum, error) {
	f := new(m.Forum)
	err := row.Scan(&f.Id, &f.Posts, &f.Slug, &f.Threads, &f.Title, &f.User)

	return f, err
}

//SELECT id, slug, title, votes, forum, author, created, message
func ScanThreadFromRow(row IRow) (*m.Thread, error) {
	t := new(m.Thread)
	err := row.Scan(&t.Id, &t.Slug, &t.Title, &t.Votes, &t.Forum, &t.Author, &t.Created, &t.Message)

	return t, err
}

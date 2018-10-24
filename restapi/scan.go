package restapi

import (
	m "http-api-server/models"
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

//SELECT id, forum, parent, thread, author, message, created, isEdited
func ScanPostFromRow(row IRow) (*m.Post, error) {
	p := new(m.Post)
	err := row.Scan(&p.Id, &p.Forum, &p.Parent, &p.Thread, &p.Author, &p.Message, &p.Created, &p.Isedited)

	return p, err
}

//p.id, u.nickname, p.created, f.slug, p.edited, p.message, coalesce(p.parent_id, 0), t.id, f.id
func ScanPostDetailsFromRow(row IRow) (*m.PostFull, error) {
	p := new(m.PostFull)
	err := row.Scan(&p.Id, &p.Author, &p.Created, &p.Forum, &p.IsEdited, &p.Message, &p.Parent, &p.ThreadId, &p.AuthorId, &p.ForumId)

	return p, err
}

func ScanStatusFromRow(row IRow) (*m.Status, error) {
	s := new(m.Status)
	err := row.Scan(&s.User, &s.Thread, &s.Post, &s.Forum)

	return s, err
}

package models

type PostFull struct {
	Id       int    `json:"id"`
	Author   string `json:"author"`
	Created  string `json:"created"`
	Forum    string `json:"forum"`
	IsEdited bool   `json:"isEdited"`
	Message  string `json:"message"`
	Parent   int    `json:"parent"`
	ThreadId int    `json:"thread"`
	AuthorId int    `json:"-"`
	ForumId  int    `json:"-"`
}

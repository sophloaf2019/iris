package social

import "iris/domain/types/base"

type Post struct {
	base.Entity
	Title   string `json:"title"`
	UserID  int    `json:"userID"`
	Message string `json:"message"`
}

func (p *Post) SetAuthor(id int) {
	p.UserID = id
}

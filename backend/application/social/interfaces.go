package social

import "iris/domain/types/social"

type PostRepo interface {
	Get(id int) (*social.Post, error)
	GetMany() ([]*social.Post, error)
	Create(post *social.Post) error
	Update(post *social.Post) error
	Delete(id int) error
}

type CommentRepo interface {
	Get(id int) (*social.Comment, error)
	Create(comment *social.Comment) error
	Update(comment *social.Comment) error
	Delete(id int) error
	For(id int, contentType string) ([]*social.Comment, error)
}

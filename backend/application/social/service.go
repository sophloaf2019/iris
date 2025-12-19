package social

import (
	"errors"
	"iris/domain/types/auth"
	"iris/domain/types/social"
	"sort"
)

// GetPost
// MakePost
// SavePost
// DeletePost
// GetComment
// MakeComment
// SaveComment
// DeleteComment

type Service struct {
	postRepo    PostRepo
	commentRepo CommentRepo
}

func NewService(postRepo PostRepo, commentRepo CommentRepo) *Service {
	return &Service{postRepo: postRepo, commentRepo: commentRepo}
}

func (s *Service) GetPost(ctx auth.Context, id int) (*social.Post, error) {
	return s.postRepo.Get(id)
}

func (s *Service) MakePost(ctx auth.Context, post *social.Post) error {
	if !auth.Can(ctx.Clearance, auth.ActionCreate, auth.ContentPost, ctx.UserID == post.UserID) {
		return errors.New("access denied")
	}
	if post.UserID == 0 {
		post.UserID = ctx.UserID
	}
	return s.postRepo.Create(post)
}

func (s *Service) SavePost(ctx auth.Context, post *social.Post) error {
	_, err := s.GetPost(ctx, post.ID)
	if err != nil {
		return err
	}
	if !auth.Can(ctx.Clearance, auth.ActionEdit, auth.ContentPost, ctx.UserID == post.UserID) {
		return errors.New("access denied")
	}
	return s.postRepo.Update(post)
}

func (s *Service) DeletePost(ctx auth.Context, id int) error {
	post, err := s.GetPost(ctx, id)
	if err != nil {
		return err
	}
	if !auth.Can(ctx.Clearance, auth.ActionDelete, auth.ContentPost, ctx.UserID == post.UserID) {
		return errors.New("access denied")
	}
	return s.postRepo.Delete(id)
}

func (s *Service) GetPosts(ctx auth.Context) ([]*social.Post, error) {
	posts, err := s.postRepo.GetMany()
	if err != nil {
		return nil, err
	}
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].CreatedAt.After(posts[j].CreatedAt)
	})
	return posts, nil
}

func (s *Service) GetComment(ctx auth.Context, id int) (*social.Comment, error) {
	return s.commentRepo.Get(id)
}

func (s *Service) MakeComment(ctx auth.Context, comment *social.Comment) error {
	if !auth.Can(ctx.Clearance, auth.ActionCreate, auth.ContentComment, ctx.UserID == comment.UserID) {
		return errors.New("access denied")
	}
	if comment.UserID == 0 {
		comment.UserID = ctx.UserID
	}
	return s.commentRepo.Create(comment)
}

func (s *Service) SaveComment(ctx auth.Context, comment *social.Comment) error {
	_, err := s.GetComment(ctx, comment.ID)
	if err != nil {
		return err
	}
	if !auth.Can(ctx.Clearance, auth.ActionEdit, auth.ContentComment, ctx.UserID == comment.UserID) {
		return errors.New("access denied")
	}
	return s.commentRepo.Update(comment)
}

func (s *Service) DeleteComment(ctx auth.Context, id int) error {
	comment, err := s.GetComment(ctx, id)
	if err != nil {
		return err
	}
	if !auth.Can(ctx.Clearance, auth.ActionDelete, auth.ContentComment, ctx.UserID == comment.UserID) {
		return errors.New("access denied")
	}
	return s.commentRepo.Delete(id)
}

func (s *Service) GetCommentsFor(ctx auth.Context, parentID int, contentType string) ([]*social.Comment, error) {
	cm, err := s.commentRepo.For(parentID, contentType)
	if err != nil {
		return nil, err
	}
	sort.Slice(cm, func(i, j int) bool {
		return cm[i].CreatedAt.After(cm[j].CreatedAt)
	})
	return cm, nil
}

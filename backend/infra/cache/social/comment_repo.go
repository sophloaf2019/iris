package social

import (
	"errors"
	"iris/domain/types/social"
	"sync"
	"time"
)

type CommentRepo struct {
	cache    map[int]*social.Comment
	mu       sync.Mutex
	lastUsed int
}

func (p *CommentRepo) For(id int, contentType string) ([]*social.Comment, error) {
	results := make([]*social.Comment, 0)
	for _, comment := range p.cache {
		if comment.ParentID == id && comment.ParentType == contentType {
			results = append(results, comment)
		}
	}
	return results, nil
}

func (p *CommentRepo) Get(id int) (*social.Comment, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	comment, ok := p.cache[id]
	if !ok {
		return nil, errors.New("comment not found")
	}
	if comment.DeletedAt.IsZero() == false {
		return nil, errors.New("comment not found")
	}
	return comment, nil
}

func (p *CommentRepo) Create(comment *social.Comment) error {
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()
	comment.DeletedAt = time.Time{}
	comment.ID = p.lastUsed + 1
	p.mu.Lock()
	defer p.mu.Unlock()
	p.cache[comment.ID] = comment
	p.lastUsed++
	return nil
}

func (p *CommentRepo) Update(comment *social.Comment) error {
	comment.UpdatedAt = time.Now()
	p.mu.Lock()
	defer p.mu.Unlock()
	if _, ok := p.cache[comment.ID]; !ok {
		return errors.New("comment not found")
	}
	p.cache[comment.ID] = comment
	return nil
}

func (p *CommentRepo) Delete(id int) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.cache, id)
	p.lastUsed = id
	return nil
}

func NewCommentRepo() *CommentRepo {
	return &CommentRepo{
		cache: make(map[int]*social.Comment),
	}
}

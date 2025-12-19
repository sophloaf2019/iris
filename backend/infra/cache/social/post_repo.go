package social

import (
	"errors"
	"iris/domain/types/social"
	"sync"
	"time"
)

type PostRepo struct {
	cache    map[int]*social.Post
	mu       sync.Mutex
	lastUsed int
}

func (p *PostRepo) GetMany() ([]*social.Post, error) {
	results := make([]*social.Post, 0)
	for _, v := range p.cache {
		results = append(results, v)
	}
	return results, nil
}

func (p *PostRepo) Get(id int) (*social.Post, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	post, ok := p.cache[id]
	if !ok {
		return nil, errors.New("post not found")
	}
	if post.DeletedAt.IsZero() == false {
		return nil, errors.New("post not found")
	}
	return post, nil
}

func (p *PostRepo) Create(post *social.Post) error {
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	post.DeletedAt = time.Time{}
	post.ID = p.lastUsed + 1
	p.mu.Lock()
	defer p.mu.Unlock()
	p.cache[post.ID] = post
	p.lastUsed++
	return nil
}

func (p *PostRepo) Update(post *social.Post) error {
	post.UpdatedAt = time.Now()
	p.mu.Lock()
	defer p.mu.Unlock()
	if _, ok := p.cache[post.ID]; !ok {
		return errors.New("post not found")
	}
	p.cache[post.ID] = post
	return nil
}

func (p *PostRepo) Delete(id int) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	delete(p.cache, id)
	p.lastUsed = id
	return nil
}

func NewPostRepo() *PostRepo {
	return &PostRepo{
		cache: make(map[int]*social.Post),
	}
}

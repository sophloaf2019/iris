package fieldwork

import (
	"errors"
	"iris/domain/types/fieldwork"
	"sync"
	"time"
)

type GOIRepo struct {
	cache      map[int]*fieldwork.GOI
	mu         sync.Mutex
	lastUsedID int
}

func (d *GOIRepo) GetMany() ([]*fieldwork.GOI, error) {
	results := make([]*fieldwork.GOI, 0)
	d.mu.Lock()
	defer d.mu.Unlock()
	for _, v := range d.cache {
		results = append(results, v)
	}
	return results, nil
}

func (d *GOIRepo) Slug(slug string) (*fieldwork.GOI, error) {
	for _, v := range d.cache {
		if v.Slug == slug {
			return v, nil
		}
	}
	return nil, errors.New("not found")
}

func (d *GOIRepo) Get(id int) (*fieldwork.GOI, error) {
	goi, ok := d.cache[id]
	if !ok {
		return nil, errors.New("debrief not found")
	}
	if !goi.DeletedAt.IsZero() {
		return nil, errors.New("debrief not found")
	}
	return goi, nil
}

func (d *GOIRepo) Create(goi *fieldwork.GOI) error {
	goi.ID = d.lastUsedID + 1
	goi.CreatedAt = time.Now()
	goi.UpdatedAt = time.Now()
	goi.DeletedAt = time.Time{}
	d.mu.Lock()
	defer d.mu.Unlock()
	d.cache[goi.ID] = goi
	return nil
}

func (d *GOIRepo) Update(goi *fieldwork.GOI) error {
	goi.ID = d.lastUsedID + 1
	goi.UpdatedAt = time.Now()
	d.mu.Lock()
	defer d.mu.Unlock()
	d.cache[goi.ID] = goi
	d.lastUsedID = goi.ID
	return nil
}

func (d *GOIRepo) Delete(id int) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.cache, id)
	return nil
}

func NewGOIRepo() *GOIRepo {
	return &GOIRepo{
		cache: make(map[int]*fieldwork.GOI),
	}
}

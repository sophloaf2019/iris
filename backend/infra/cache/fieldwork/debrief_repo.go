package fieldwork

import (
	"errors"
	"iris/domain/types/fieldwork"
	"sync"
	"time"
)

type DebriefRepo struct {
	cache      map[int]*fieldwork.Debrief
	mu         sync.Mutex
	lastUsedID int
}

func (d *DebriefRepo) GetMany() ([]*fieldwork.Debrief, error) {
	results := make([]*fieldwork.Debrief, 0)
	d.mu.Lock()
	defer d.mu.Unlock()
	for _, v := range d.cache {
		results = append(results, v)
	}
	return results, nil
}

func (d *DebriefRepo) Slug(slug string) (*fieldwork.Debrief, error) {
	for _, v := range d.cache {
		if v.Slug == slug {
			return v, nil
		}
	}
	return nil, errors.New("not found")
}

func (d *DebriefRepo) Get(id int) (*fieldwork.Debrief, error) {
	dbf, ok := d.cache[id]
	if !ok {
		return nil, errors.New("debrief not found")
	}
	if !dbf.DeletedAt.IsZero() {
		return nil, errors.New("debrief not found")
	}
	return dbf, nil
}

func (d *DebriefRepo) Create(dbf *fieldwork.Debrief) error {
	dbf.ID = d.lastUsedID + 1
	dbf.CreatedAt = time.Now()
	dbf.UpdatedAt = time.Now()
	dbf.DeletedAt = time.Time{}
	d.mu.Lock()
	defer d.mu.Unlock()
	d.cache[dbf.ID] = dbf
	return nil
}

func (d *DebriefRepo) Update(dbf *fieldwork.Debrief) error {
	dbf.ID = d.lastUsedID + 1
	dbf.UpdatedAt = time.Now()
	d.mu.Lock()
	defer d.mu.Unlock()
	d.cache[dbf.ID] = dbf
	return nil
}

func (d *DebriefRepo) Delete(id int) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.cache, id)
	d.lastUsedID = id
	return nil
}

func NewDebriefRepo() *DebriefRepo {
	return &DebriefRepo{
		cache: make(map[int]*fieldwork.Debrief),
	}
}

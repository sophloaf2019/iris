package fieldwork

import (
	"errors"
	"iris/domain/types/fieldwork"
	"sync"
	"time"
)

type MissionRepo struct {
	cache      map[int]*fieldwork.Mission
	mu         sync.Mutex
	lastUsedID int
}

func (d *MissionRepo) GetMany() ([]*fieldwork.Mission, error) {
	var result []*fieldwork.Mission
	d.mu.Lock()
	defer d.mu.Unlock()
	for _, v := range d.cache {
		result = append(result, v)
	}
	return result, nil
}

func (d *MissionRepo) Slug(slug string) (*fieldwork.Mission, error) {
	for _, v := range d.cache {
		if v.Slug == slug {
			return v, nil
		}
	}
	return nil, errors.New("not found")
}

func (d *MissionRepo) Get(id int) (*fieldwork.Mission, error) {
	msn, ok := d.cache[id]
	if !ok {
		return nil, errors.New("debrief not found")
	}
	if !msn.DeletedAt.IsZero() {
		return nil, errors.New("debrief not found")
	}
	return msn, nil
}

func (d *MissionRepo) Create(msn *fieldwork.Mission) error {
	msn.ID = d.lastUsedID + 1
	msn.CreatedAt = time.Now()
	msn.UpdatedAt = time.Now()
	msn.DeletedAt = time.Time{}
	d.mu.Lock()
	defer d.mu.Unlock()
	d.cache[msn.ID] = msn
	return nil
}

func (d *MissionRepo) Update(msn *fieldwork.Mission) error {
	msn.ID = d.lastUsedID + 1
	msn.UpdatedAt = time.Now()
	d.mu.Lock()
	defer d.mu.Unlock()
	d.cache[msn.ID] = msn
	return nil
}

func (d *MissionRepo) Delete(id int) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	delete(d.cache, id)
	d.lastUsedID = id
	return nil
}

func NewMissionRepo() *MissionRepo {
	return &MissionRepo{
		cache: make(map[int]*fieldwork.Mission),
	}
}

package fieldwork

import "iris/domain/types/fieldwork"

type DebriefRepo interface {
	Get(id int) (*fieldwork.Debrief, error)
	GetMany() ([]*fieldwork.Debrief, error)
	Slug(slug string) (*fieldwork.Debrief, error)
	Create(dbf *fieldwork.Debrief) error
	Update(dbf *fieldwork.Debrief) error
	Delete(id int) error
}

type GOIRepo interface {
	Get(id int) (*fieldwork.GOI, error)
	GetMany() ([]*fieldwork.GOI, error)
	Slug(slug string) (*fieldwork.GOI, error)
	Create(dbf *fieldwork.GOI) error
	Update(dbf *fieldwork.GOI) error
	Delete(id int) error
}

type MissionRepo interface {
	Get(id int) (*fieldwork.Mission, error)
	GetMany() ([]*fieldwork.Mission, error)
	Slug(slug string) (*fieldwork.Mission, error)
	Create(dbf *fieldwork.Mission) error
	Update(dbf *fieldwork.Mission) error
	Delete(id int) error
}

package types

import "time"

type Entity struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (e *Entity) Entity() *Entity {
	return e
}

package fieldwork

import (
	"image/color"
	"iris/domain/types/base"
	"iris/domain/values"
)

type GOI struct {
	base.Entity
	AuthorID       int               `json:"authorID"`
	Name           string            `json:"name"`
	Slug           string            `json:"slug"`
	PrimaryColor   color.RGBA        `json:"primaryColor"`
	SecondaryColor color.RGBA        `json:"secondaryColor"`
	MO             string            `json:"mo"`
	Locations      []values.Location `json:"locations"`
	Active         bool              `json:"active"`
	AssignedTo     int               `json:"assignedTo"`
}

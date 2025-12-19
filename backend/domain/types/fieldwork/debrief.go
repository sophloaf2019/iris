package fieldwork

import (
	"iris/domain/types/base"
)

type Debrief struct {
	base.Entity
	MissionID int    `json:"missionID"`
	Title     string `json:"title"`
	Slug      string `json:"slug"`
	AuthorID  int    `json:"authorID"`
	Summary   string `json:"summary"`
	UserIDs   []int  `json:"userIDs"`
}

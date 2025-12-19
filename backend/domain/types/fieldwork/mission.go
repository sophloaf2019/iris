package fieldwork

import (
	"iris/domain/types/base"
	"iris/domain/values"
)

type Mission struct {
	base.Entity
	Title           string          `json:"title"`
	Slug            string          `json:"slug"`
	Briefing        string          `json:"briefing"`
	Tags            []string        `json:"tags"`
	GoiID           int             `json:"goiID"`
	AuthorID        int             `json:"authorID"`
	Location        values.Location `json:"location"`
	InterestedUsers []int           `json:"interestedUsers"`
}

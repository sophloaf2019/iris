package social

import "iris/domain/types/base"

type Comment struct {
	base.Entity
	UserID     int    `json:"userID"`
	Message    string `json:"message"`
	ParentID   int    `json:"parentID"`
	ParentType string `json:"parentType"`
}

const (
	ContentUser    = "user"
	ContentPost    = "post"
	ContentComment = "comment"
	ContentGOI     = "goi"
	ContentDebrief = "debrief"
	ContentMission = "mission"
)

func (c *Comment) SetAuthor(id int) {
	c.UserID = id
}

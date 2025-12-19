package types

type Debrief struct {
	*Entity
	MissionID int
	UserIDs   []int
	Message   string
}

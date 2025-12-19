package types

type Mission struct {
	*Entity
	GroupID  int
	Title    string
	Briefing string
	Location string
}

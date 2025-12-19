package types

type Group struct {
	*Entity
	Name      string
	MO        string
	Locations []string
}

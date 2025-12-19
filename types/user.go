package types

type User struct {
	*Entity
	Username       string
	HashedPassword string
	Clearance      int
}

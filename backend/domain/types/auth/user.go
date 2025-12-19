package auth

import (
	"errors"
	"iris/domain/types/base"
)

type User struct {
	base.Entity
	Username       string    `json:"username"`
	HashedPassword string    `json:"-"`
	Clearance      Clearance `json:"clearance"`
}

// NewUser creates a new User object with the provided username, hashed password,
// and clearance level.
func NewUser(username, hashedPassword string, clearance Clearance) (*User, error) {
	if clearance <= 0 || clearance > 5 {
		return nil, errors.New("clearance must be between 0 and 5")
	}
	return &User{
		Username:       username,
		HashedPassword: hashedPassword,
		Clearance:      clearance,
	}, nil
}

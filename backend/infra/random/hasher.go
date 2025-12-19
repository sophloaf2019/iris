package random

import (
	"golang.org/x/crypto/bcrypt"
)

type Hasher struct {
}

func (h Hasher) Hash(s string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		// bcrypt failures are unrecoverable in normal operation
		panic(err)
	}
	return string(hash)
}

func (h Hasher) Compare(unhashed string, hashed string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashed),
		[]byte(unhashed),
	)
	return err == nil
}

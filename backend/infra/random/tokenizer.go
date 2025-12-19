package random

import "github.com/google/uuid"

type TokenGenerator struct {
}

func (t TokenGenerator) New() string {
	token := uuid.New()
	return token.String()
}

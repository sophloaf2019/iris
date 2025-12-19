package auth

import "iris/domain/types/auth"

type TokenProvider interface {
	New() string
}

type SessionManager interface {
	Save(token string, session *auth.Session) error
	Get(token string) (*auth.Session, error)
	Delete(token string) error
}

type Hasher interface {
	Hash(string) string
	Compare(unhashed string, hashed string) bool
}

type UserRepository interface {
	Create(user *auth.User) error
	Update(user *auth.User) error
	Get(id int) (*auth.User, error)
	GetMany() ([]*auth.User, error)
	GetByUsername(username string) (*auth.User, error)
}

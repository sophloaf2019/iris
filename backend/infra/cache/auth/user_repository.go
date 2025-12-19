package auth

import (
	"errors"
	"iris/domain/types/auth"
	"sync"
	"time"
)

type UserRepository struct {
	mu         sync.RWMutex
	cache      map[int]*auth.User
	lastUsedID int
}

func (u *UserRepository) GetMany() ([]*auth.User, error) {
	results := make([]*auth.User, 0)
	u.mu.RLock()
	defer u.mu.RUnlock()
	for _, v := range u.cache {
		results = append(results, v)
	}
	return results, nil
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		cache: make(map[int]*auth.User),
	}
}

func (u *UserRepository) Create(user *auth.User) error {
	user.ID = u.lastUsedID + 1
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.DeletedAt = time.Time{}
	u.mu.Lock()
	u.cache[user.ID] = user
	u.lastUsedID++
	u.mu.Unlock()
	return nil
}

func (u *UserRepository) Update(user *auth.User) error {
	user.UpdatedAt = time.Now()
	u.mu.Lock()
	u.cache[user.ID] = user
	u.lastUsedID++
	u.mu.Unlock()
	return nil
}

func (u *UserRepository) Get(id int) (*auth.User, error) {
	user, ok := u.cache[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	if !user.DeletedAt.IsZero() {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (u *UserRepository) GetByUsername(username string) (*auth.User, error) {
	for _, user := range u.cache {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}

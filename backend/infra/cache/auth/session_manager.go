package auth

import (
	"errors"
	"iris/domain/types/auth"
	"sync"
)

type SessionManager struct {
	mu    sync.RWMutex
	cache map[string]*auth.Session
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		cache: make(map[string]*auth.Session),
	}
}

func (s *SessionManager) Save(token string, session *auth.Session) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cache[token] = session
	return nil
}

func (s *SessionManager) Get(token string) (*auth.Session, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	sesh, ok := s.cache[token]
	if !ok {
		return nil, errors.New("session not found")
	}
	return sesh, nil
}

func (s *SessionManager) Delete(token string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.cache, token)
	return nil
}

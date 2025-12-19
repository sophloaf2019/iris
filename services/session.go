package services

import (
	"fmt"
	"iris/types"
	"time"
)

type Session struct {
	User    types.User
	Expires time.Time
}

func (s Session) IsExpired() bool {
	return s.Expires.Before(time.Now())
}

type SessionManager struct {
	sessions map[string]Session
}

func NewSessionManager() *SessionManager {
	return &SessionManager{sessions: make(map[string]Session)}
}

func (s *SessionManager) Store(user types.User) string {
	id := fmt.Sprintf("%d", time.Now().UnixNano())
	s.sessions[id] = Session{user, time.Now().Add(time.Hour * 3)}
	return id
}

func (s *SessionManager) Load(id string) (Session, error) {
	session, ok := s.sessions[id]
	if !ok {
		return Session{}, fmt.Errorf("session not found")
	}
	if session.IsExpired() {
		delete(s.sessions, id)
		return Session{}, fmt.Errorf("session is expired")
	}
	return session, nil
}

func (s *SessionManager) Delete(id string) {
	delete(s.sessions, id)
}

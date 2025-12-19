package auth

import "time"

type Session struct {
	ExpiresAt time.Time
	Context   *Context
}

func (s *Session) IsExpired() bool {
	return s.ExpiresAt.Before(time.Now())
}

func NewSession(duration time.Duration, context *Context) *Session {
	return &Session{
		ExpiresAt: time.Now().Add(duration),
		Context:   context,
	}
}

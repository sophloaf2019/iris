package auth

import (
	"errors"
	"iris/domain/types/auth"
	"time"
)

// Login
// Logout
// IssueNewUser
// ValidateToken
// ResetPassword

type Service struct {
	tp TokenProvider
	sm SessionManager
	h  Hasher
	ur UserRepository
}

func NewService(tp TokenProvider, sm SessionManager, h Hasher, ur UserRepository) *Service {
	return &Service{tp: tp, sm: sm, h: h, ur: ur}
}

func (s *Service) Login(username string, password string) (string, error) {
	user, err := s.ur.GetByUsername(username)
	if err != nil {
		return "", err
	}
	if !s.h.Compare(password, user.HashedPassword) {
		return "", errors.New("invalid password")
	}
	token := s.tp.New()
	newCtx := auth.NewContext(user, token)

	session := auth.NewSession(time.Hour*1, newCtx)
	if err := s.sm.Save(token, session); err != nil {
		return "", err
	}
	return token, nil
}

func (s *Service) Logout(token string) error {
	return s.sm.Delete(token)
}

func (s *Service) IssueNewUser(ctx auth.Context, username, password string, clearance auth.Clearance) (*auth.User, error) {
	if !auth.AtLeast(ctx.Clearance, auth.ClearanceTopSecret) {
		return nil, errors.New("invalid clearance")
	}
	user, err := auth.NewUser(username, s.h.Hash(password), clearance)
	if err != nil {
		return nil, err
	}
	err = s.ur.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) ValidateToken(token string) (auth.Context, error) {
	session, err := s.sm.Get(token)
	if err != nil {
		return auth.Context{}, err
	}
	if session.IsExpired() {
		return auth.Context{}, errors.New("token is expired")
	}
	return *session.Context, nil
}

func (s *Service) ResetPassword(ctx auth.Context, userID int, oldPassword, newPassword string) error {
	user, err := s.ur.Get(userID)
	if err != nil {
		return err
	}
	if ctx.UserID != userID {
		if !auth.AtLeast(ctx.Clearance, auth.ClearanceTopSecret) {
			return errors.New("invalid clearance")
		}
	} else {
		if !s.h.Compare(oldPassword, user.HashedPassword) {
			return errors.New("invalid old password")
		}
	}
	user.HashedPassword = s.h.Hash(newPassword)
	return s.ur.Update(user)
}

func (s *Service) GetUserByUsername(ctx auth.Context, username string) (*auth.User, error) {
	return s.ur.GetByUsername(username)
}

func (s *Service) GetUserByID(ctx auth.Context, userID int) (*auth.User, error) {
	return s.ur.Get(userID)
}

func (s *Service) GetUsers(ctx auth.Context) ([]*auth.User, error) {
	return s.ur.GetMany()
}

func (s *Service) UserCan(
	ctx auth.Context,
	userID int,
	action auth.Action,
	contentType auth.Content,
	isOwner bool,
) (bool, error) {
	user, err := s.GetUserByID(ctx, userID)
	if err != nil {
		return false, err
	}
	return auth.Can(user.Clearance, action, contentType, isOwner), nil
}

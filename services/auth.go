package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"iris/types"
)

type AuthService struct {
	sessionManager *SessionManager
	users          *Storage[types.User]
}

func NewAuthService(userStorage *Storage[types.User], manager *SessionManager) *AuthService {
	return &AuthService{users: userStorage, sessionManager: manager}
}

// Login
// Logout
// IssueNewUser
// ResetPassword

func (s *AuthService) Login(username, password string) (string, error) {
	users, err := s.users.GetAll()
	if err != nil {
		return "", err
	}
	for _, user := range users {
		if user.Username == username {
			err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
			if err != nil {
				return "", err
			}
			return s.sessionManager.Store(user), nil
		}
	}
	return "", errors.New("user not found")
}

func (s *AuthService) Logout(token string) error {
	s.sessionManager.Delete(token)
	return nil
}

func (s *AuthService) IssueNewUser(session Session, username, password string, clearance int) error {
	if session.User.Clearance < 5 {
		return errors.New("user lacks clearance")
	}
	var user types.User
	user.Username = username
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.HashedPassword = string(hashedPwd)
	user.Clearance = clearance
	_, err := s.users.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthService) ResetPassword(session Session, userID int, oldPassword, newPassword string) error {
	// WE MUST HAVE BOTH OLD AND NEW PASSWORDS
	// UNLESS ITS FOR A DIFFERENT USER THAN US AND WE'RE
	// C-5
	if oldPassword == "" || newPassword == "" {
		if session.User.ID == userID {
			return errors.New("need old password")
		}
		if session.User.Clearance < 5 {
			return errors.New("user lacks clearance")
		}
	}
	user, err := s.users.Get(userID)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(oldPassword))
	if err != nil {
		return err
	}
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	user.HashedPassword = string(hashedPwd)
	return s.users.Update(user)
}

func (s *AuthService) GetUser(userID int) (types.User, error) {
	user, err := s.users.Get(userID)
	if err != nil {
		return types.User{}, err
	}
	return user, nil
}

func (s *AuthService) GetUserByUsername(username string) (types.User, error) {
	users, err := s.users.GetAll()
	if err != nil {
		return types.User{}, err
	}
	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}
	return types.User{}, errors.New("user not found")
}

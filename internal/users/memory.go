package users

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidUserID    = errors.New("invalid user ID")
	ErrUserNotFound     = errors.New("user not found")
	ErrUsernameNotFound = errors.New("username not found")
	ErrUsernameExists   = errors.New("username already exists")
)

type MemoryRepository struct {
	users     map[string]User
	usernames map[string]string
	lock      sync.RWMutex
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		users:     make(map[string]User),
		usernames: make(map[string]string),
	}
}

func (m *MemoryRepository) FindByID(_ context.Context, userID string) (User, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	user, ok := m.users[userID]
	if !ok {
		return User{}, ErrUserNotFound
	}

	return user, nil
}

func (m *MemoryRepository) FindByUsername(_ context.Context, username string) (User, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()

	id, ok := m.usernames[username]
	if !ok {
		return User{}, ErrUsernameNotFound
	}

	user, ok := m.users[id]
	if !ok {
		return User{}, ErrUserNotFound
	}

	return user, nil
}

func (m *MemoryRepository) Create(_ context.Context, user User) (User, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if _, found := m.usernames[user.Username]; found {
		return User{}, ErrUsernameExists
	}

	id := uuid.New().String()

	user.ID = id
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Time{}

	m.users[id] = user
	m.usernames[user.Username] = id

	return user, nil
}

func (m *MemoryRepository) Update(_ context.Context, user User) (User, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if user.ID == "" {
		return User{}, ErrInvalidUserID
	}

	curr, found := m.users[user.ID]
	if !found {
		return User{}, ErrUserNotFound
	}

	if user.Username != curr.Username {
		if _, found := m.usernames[user.Username]; found {
			return User{}, ErrUsernameExists
		}

		m.usernames[user.Username] = user.ID
		delete(m.usernames, curr.Username)
	}

	user.UpdatedAt = time.Now()

	m.users[user.ID] = user

	return user, nil
}

func (m *MemoryRepository) Delete(_ context.Context, userID string) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	curr, found := m.users[userID]
	if !found {
		return ErrUserNotFound
	}

	delete(m.users, userID)
	delete(m.usernames, curr.Username)

	return nil
}

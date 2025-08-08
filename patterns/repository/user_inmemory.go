package repository

import (
	"context"
	"sync"
)

// InMemory Implementation for testing

type InMemoryUserRepository struct {
	mu    sync.RWMutex
	users map[string]*User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{users: make(map[string]*User)}
}

func (r *InMemoryUserRepository) Create(ctx context.Context, user *User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.ID]; exists {
		return ErrUserAlreadyExists
	}

	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) GetById(ctx context.Context, id string) (*User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	user, ok := r.users[id]
	if !ok {
		return nil, ErrUserNotFound
	}
	return user, nil
}

func (r *InMemoryUserRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, ErrUserNotFound
}

func (r *InMemoryUserRepository) Update(ctx context.Context, user *User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return ErrUserNotFound
	}

	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.users[id]; !exists {
		return ErrUserNotFound
	}

	delete(r.users, id)
	return nil
}

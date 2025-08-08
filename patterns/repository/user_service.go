package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	Repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) RegisterUser(ctx context.Context, email, name string) (*User, error) {
	// Check if user with the same email already exists
	if _, err := s.Repo.GetByEmail(ctx, email); err == nil {
		return nil, ErrUserAlreadyExists
	}

	user := &User{
		ID:        uuid.New().String(),
		Email:     email,
		Name:      name,
		CreatedAt: time.Now(),
	}

	if err := s.Repo.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return user, nil

}

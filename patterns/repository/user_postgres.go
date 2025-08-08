package repository

import (
	"context"
	"database/sql"
)

// PostgresUserRepository implementation
type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(ctx context.Context, user *User) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users (id, email, name, created_at) VALUES ($1, $2, $3, $4)", user.ID, user.Email, user.Name, user.CreatedAt)
	return err
}

func (r *PostgresUserRepository) GetById(ctx context.Context, id string) (*User, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, email, name, created_at FROM users WHERE id = $1", id)
	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, email, name, created_at FROM users WHERE email = $1", email)
	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) Update(ctx context.Context, user *User) error {
	_, err := r.db.ExecContext(ctx, "UPDATE users SET email = $1, name = $2 WHERE id = $3", user.Email, user.Name, user.ID)
	return err
}

func (r *PostgresUserRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "DELETE FROM users WHERE id = $1", id)
	return err
}

package db

import (
	"context"
	"errors"
	"time"

	"github.com/asm2212/go-gin-starter/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (username, password, created_at) VALUES ($1, $2, $3) RETURNING id`
	return r.db.QueryRow(ctx, query, user.Username, user.Password, time.Now()).Scan(&user.ID)
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	query := `SELECT id, username, password, created_at FROM users WHERE username=$1`
	row := r.db.QueryRow(ctx, query, username)
	var user model.User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

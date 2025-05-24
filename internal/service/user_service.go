package service

import (
	"context"
	"errors"

	"github.com/asm2212/go-gin-starter/internal/db"
	"github.com/asm2212/go-gin-starter/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *db.UserRepository
}

func NewUserService(repo *db.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, username, password string) (*model.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		Username: username,
		Password: string(hashed),
	}
	if err := s.repo.CreateUser(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) Authenticate(ctx context.Context, username, password string) (*model.User, error) {
	user, err := s.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New("invalid credentials")
	}
	return user, nil
}

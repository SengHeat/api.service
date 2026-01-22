package service

import (
	"context"
	"fmt"
	"time"

	"api.drsb-purchase-service/config"
	"api.drsb-purchase-service/internal/domain/user"
	"api.drsb-purchase-service/pkg/jwt"
)

type userService struct {
	repo   user.Repository
	config *config.Config
}

func NewUserService(repo user.Repository, config *config.Config) user.Service {
	return &userService{repo: repo, config: config}
}

func (s userService) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.User, error) {
	// Check if email already exists
	existing, _ := s.repo.FindByEmail(ctx, req.Email)
	if existing != nil {
		return nil, fmt.Errorf("email already exists")
	}

	// Create user
	u := &user.User{
		Name:  req.Name,
		Email: req.Email,
		Role:  req.Role,
	}

	if u.Role == "" {
		u.Role = "user"
	}

	// Set password
	u.Password = req.Password
	if err := u.HashPassword(); err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Save to database
	if err := s.repo.Create(ctx, u); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return u, nil
}

func (s userService) GetUser(ctx context.Context, id int) (*user.User, error) {
	u, err := s.repo.FindById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	return u, nil
}

func (s userService) GetUsers(ctx context.Context, page, pageSize int) ([]*user.User, int, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	users, err := s.repo.FindAll(ctx, pageSize, offset)

	if err != nil {
		return nil, 0, fmt.Errorf("failed to get users: %w", err)
	}

	total, err := s.repo.Count(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count users: %w", err)
	}

	return users, total, nil
}

func (s userService) UpdateUser(ctx context.Context, id int, req *user.UpdateUserRequest) (*user.User, error) {
	u, err := s.repo.FindById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	if req.Name != "" {
		u.Name = req.Name
	}
	if req.Email != "" {
		// Check if email is taken by another user
		existing, _ := s.repo.FindByEmail(ctx, req.Email)
		if existing != nil && existing.ID != id {
			return nil, fmt.Errorf("email already exists")
		}
		u.Email = req.Email
	}
	if req.Role != "" {
		u.Role = req.Role
	}

	// Save changes
	if err := s.repo.Update(ctx, u); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return u, nil
}

func (s userService) Delete(ctx context.Context, id int) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

func (s userService) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	// Find user by email
	u, err := s.repo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Check password
	if !u.CheckPassword(req.Password) {
		return nil, fmt.Errorf("invalid credentials")
	}

	// Generate JWT token
	token, err := jwt.GenerateToken(
		u.ID,
		u.Email,
		u.Role,
		s.config.JWT.Secret,
		time.Duration(s.config.JWT.Expiration)*time.Hour,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	return &user.LoginResponse{
		Token: token,
		User:  u,
	}, nil
}

func (s userService) ValidateToken(ctx context.Context, token string) (*user.User, error) {
	// Validate and parse token
	claims, err := jwt.ValidateToken(token, s.config.JWT.Secret)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	// Get user from database
	u, err := s.repo.FindById(ctx, claims.UserID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	return u, nil
}

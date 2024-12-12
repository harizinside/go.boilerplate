package auth

import (
	"context"
	"fmt"

	"go.boilerplate/internal/model"
	"go.boilerplate/pkg/utils"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) SignUpService(ctx context.Context, name string, email string, password string) (*model.User, error) {
	hashed_password, err := utils.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	resp, err := s.repo.SignUpRepository(ctx, name, email, hashed_password)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Service) SignInService(ctx context.Context, email string, password string) (*model.User, error) {
	resp, err := s.repo.FindUserRepository(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	valid, err := utils.VerifyPassword(password, resp.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to verify password: %v", err)
	}

	if !valid {
		return nil, fmt.Errorf("invalid email or password")
	}

	return resp, nil

}

func (s *Service) RecoveryService(ctx context.Context, email string) (*model.User, error) {
	resp, err := s.repo.FindUserRepository(ctx, email)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Service) ResetPasswordService(ctx context.Context, id string, password string) (bool, error) {
	hashed_password, err := utils.HashPassword(password)
	if err != nil {
		return false, fmt.Errorf("failed to hash password: %v", err)
	}

	resp, err := s.repo.ResetPasswordRepository(ctx, id, hashed_password)
	if err != nil {
		return false, err
	}

	return resp, nil
}

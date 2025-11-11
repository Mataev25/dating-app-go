package user

import (
	"context"
	"fmt"
	"strings"
)

type Service interface {
	Register(
		ctx context.Context,
		rq *CreateUserRequest,
	) (*UserResponse, error)

	GetUser(
		ctx context.Context, 
		id string,
	) (*UserResponse, error)
}

type userService struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &userService{repo: repo}
}

func (s *userService) Register(
	ctx context.Context, 
	rq *CreateUserRequest,
) (*UserResponse, error) {
	user := rq.ToUser()
	if err := user.Validate(); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	existingUser, err := s.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		if err.Error() != "user not found" && 
			!strings.Contains(
				err.Error(), 
				"user not found",
				) {
			return nil, fmt.Errorf(
				"failed to check email existence: %w", 
				err,
				)
		}
	} else if existingUser != nil {

		return nil, fmt.Errorf(
			"user with email %s already exists", 
			user.Email,
			)
	}
	
	if err := s.repo.CreateUser(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return FromUser(user), nil
}

func (s *userService) GetUser(
	ctx context.Context, 
	id string,
	) (*UserResponse, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	
	return FromUser(user), nil
}

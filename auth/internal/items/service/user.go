package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/projects/pro-sphere-backend/auth/genproto/genproto/auth"
	"github.com/projects/pro-sphere-backend/auth/internal/items/models"
	"github.com/projects/pro-sphere-backend/auth/internal/items/repos"
)

type UserService struct {
	repo      repos.UserRepositories
	validator *validator.Validate
	logger    *zap.Logger
	auth.UnimplementedUserManagementServer
}

func NewUserService(repo repos.UserRepositories, logger *zap.Logger) *UserService {
	return &UserService{
		repo:      repo,
		validator: validator.New(),
		logger:    logger,
	}
}

// Error variables
var (
	ErrInvalidInput = errors.New("invalid input")
	ErrUserNotFound = models.ErrUserNotFound
)

// GetUserByID retrieves a user by ID with validation and error handling
func (s *UserService) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	if id <= 0 {
		s.logger.Warn("invalid user ID", zap.Int("id", id))
		return nil, fmt.Errorf("%w: ID must be a positive integer", ErrInvalidInput)
	}

	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			s.logger.Warn("user not found", zap.Int("id", id))
			return nil, ErrUserNotFound
		}
		s.logger.Error("failed to retrieve user", zap.Error(err))
		return nil, err
	}

	return user, nil
}

// CreateUser creates a new user with validation
func (s *UserService) CreateUser(ctx context.Context, input *models.CreateUser) (*models.User, error) {
	if err := s.validator.Struct(input); err != nil {
		s.logger.Warn("validation failed", zap.Error(err))
		return nil, fmt.Errorf("%w: %v", ErrInvalidInput, err)
	}

	createdUser, err := s.repo.CreateUser(ctx, input)
	if err != nil {
		s.logger.Error("failed to create user", zap.Error(err))
		return nil, err
	}

	return createdUser, nil
}

// UpdateUser updates an existing user with validation
func (s *UserService) UpdateUser(ctx context.Context, input *models.UpdateUser) (*models.User, error) {
	if input.ID <= 0 {
		s.logger.Warn("invalid user ID for update", zap.Int("id", input.ID))
		return nil, fmt.Errorf("%w: ID must be a positive integer", ErrInvalidInput)
	}

	updatedUser, err := s.repo.UpdateUser(ctx, input)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			s.logger.Warn("user not found for update", zap.Int("id", input.ID))
			return nil, ErrUserNotFound
		}
		s.logger.Error("failed to update user", zap.Error(err))
		return nil, err
	}

	return updatedUser, nil
}

// DeleteUser deletes a user by ID with validation
func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	if id <= 0 {
		s.logger.Warn("invalid user ID for delete", zap.Int("id", id))
		return fmt.Errorf("%w: ID must be a positive integer", ErrInvalidInput)
	}

	err := s.repo.DeleteUser(ctx, id)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			s.logger.Warn("user not found for delete", zap.Int("id", id))
			return ErrUserNotFound
		}
		s.logger.Error("failed to delete user", zap.Error(err))
		return err
	}

	return nil
}

// GetUserByProID retrieves a user by ProID with validation
func (s *UserService) GetUserByProID(ctx context.Context, proID int) (*models.User, error) {
	if proID <= 0 {
		s.logger.Warn("invalid pro ID", zap.Int("proID", proID))
		return nil, fmt.Errorf("%w: ProID must be a positive integer", ErrInvalidInput)
	}

	user, err := s.repo.GetUserByPROID(ctx, proID)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			s.logger.Warn("user not found by ProID", zap.Int("proID", proID))
			return nil, ErrUserNotFound
		}
		s.logger.Error("failed to retrieve user by ProID", zap.Error(err))
		return nil, err
	}

	return user, nil
}

// GetAllUsers retrieves a list of users based on the provided filters and pagination.
func (s *UserService) GetAllUsers(ctx context.Context, filter *models.GetAllUsers) (*models.GetAllUsersRes, error) {
	if err := s.validator.Struct(filter); err != nil {
		s.logger.Warn("validation failed for GetAllUsers filter", zap.Error(err))
		return nil, fmt.Errorf("%w: %v", ErrInvalidInput, err)
	}

	users, err := s.repo.GetAllUsers(ctx, filter)
	if err != nil {
		s.logger.Error("failed to get all users", zap.Error(err))
		return nil, err
	}

	return users, nil
}

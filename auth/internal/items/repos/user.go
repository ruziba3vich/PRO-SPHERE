/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-12-05 01:04:04
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-05 03:38:42
 * @FilePath: /auth/internal/items/repo/user.go
 * @Description: Repository interface and models for user management.
 */
package repos

import (
	"context"
	"time"

	"github.com/projects/pro-sphere-backend/auth/internal/items/models"
)

type UserRepositories interface {
	// Basic CRUD
	CreateUser(ctx context.Context, user *models.CreateUser) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.UpdateUser) (*models.User, error)
	DeleteUser(ctx context.Context, id int) error
	GetUser(ctx context.Context, id int) (*models.User, error)
	GetUserByPROID(ctx context.Context, proID int) (*models.User, error)
	GetAllUsers(ctx context.Context, filter *models.GetAllUsers) (*models.GetAllUsersRes, error)

	// Filters and Search
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	SearchUsers(ctx context.Context, query string, limit, offset int) (*models.GetAllUsersRes, error)
	GetUsersByRole(ctx context.Context, role string) (*models.GetAllUsersRes, error)
	GetUsersByDateRange(ctx context.Context, startDate, endDate time.Time) (*models.GetAllUsersRes, error)

	// Status Management
	ActivateUser(ctx context.Context, id int) error
	DeactivateUser(ctx context.Context, id int) error
	ToggleUserStatus(ctx context.Context, id int) error

	// Authentication
	UpdatePassword(ctx context.Context, id int, hashedPassword string) error

	// Audit
	LogLastLogin(ctx context.Context, id int, loginTime time.Time) error
	GetRecentlyCreatedUsers(ctx context.Context, limit, offset int) (*models.GetAllUsersRes, error)

	// Custom
	AssignProID(ctx context.Context, userID, proID int) error
}

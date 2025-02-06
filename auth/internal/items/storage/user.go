/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-12-05 01:43:27
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-21 01:25:21
 * @FilePath: /auth/internal/items/storage/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package storage

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"

	"github.com/projects/pro-sphere-backend/auth/internal/items/models"
	"github.com/projects/pro-sphere-backend/auth/internal/items/repos"
)

type UserRepository struct {
	db         *sql.DB
	sqlBuilder sq.StatementBuilderType
	redis      *redis.Client
	logger     *zap.Logger
}

func NewUserRepository(db *sql.DB, redisClient *redis.Client, logger *zap.Logger) repos.UserRepositories {
	return &UserRepository{
		db:         db,
		sqlBuilder: sq.StatementBuilderType{}.PlaceholderFormat(sq.Dollar),
		redis:      redisClient,
		logger:     logger,
	}
}

func (r *UserRepository) GetUserByPROID(ctx context.Context, proID int) (*models.User, error) {
	cacheKey := fmt.Sprintf("userProID:%d", proID)
	cachedUser, err := r.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var user models.User
		if err := json.Unmarshal([]byte(cachedUser), &user); err == nil {
			return &user, nil
		}
	}

	query := r.sqlBuilder.
		Select("id", "pro_id", "first_name", "last_name", "email", "date_of_birth",
			"gender", "avatar", "role", "created_at", "updated_at", "phone").
		From("users").
		Where(sq.Eq{"pro_id": proID})

	sqlq, args, err := query.ToSql()
	if err != nil {
		r.logger.Error("failed to build query", zap.Error(err))
		return nil, fmt.Errorf("failed to build query: %v", err)
	}

	user := &models.User{}
	err = r.db.QueryRowContext(ctx, sqlq, args...).Scan(
		&user.ID, &user.ProID, &user.FirstName,
		&user.LastName, &user.Email, &user.DateOfBirth,
		&user.Gender, &user.Avatar, &user.Role,
		&user.CreatedAt, &user.UpdatedAt, &user.Phone,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, models.ErrUserNotFound // User not found
		}
		r.logger.Error("failed to get user", zap.Error(err))
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	// Cache the user
	if err := r.cacheUser(ctx, user); err != nil {
		r.logger.Warn("failed to cache user", zap.Error(err))
	}

	return user, nil
}

// CreateUser implements user creation with transaction and caching
func (r *UserRepository) CreateUser(ctx context.Context, user *models.CreateUser) (*models.User, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		r.logger.Error("failed to begin transaction", zap.Error(err))
		return nil, fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	query := r.sqlBuilder.
		Insert("users").
		Columns(
			"pro_id", "first_name", "last_name", "email",
			"date_of_birth", "gender", "avatar", "role", "phone",
		).
		Values(
			user.ProID, user.FirstName, user.LastName, user.Email,
			user.DateOfBirth, user.Gender, user.Avatar, user.Role, user.Phone,
		).
		Suffix("RETURNING id, created_at, updated_at")

	sql, args, err := query.ToSql()
	if err != nil {
		r.logger.Error("failed to build query", zap.Error(err))
		return nil, fmt.Errorf("failed to build query: %v", err)
	}

	createdUser := &models.User{
		ProID:       user.ProID,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		DateOfBirth: user.DateOfBirth,
		Gender:      user.Gender,
		Avatar:      user.Avatar,
		Role:        user.Role,
		Phone:       user.Phone,
	}

	err = tx.QueryRowContext(ctx, sql, args...).Scan(
		&createdUser.ID,
		&createdUser.CreatedAt,
		&createdUser.UpdatedAt,
	)
	if err != nil {
		r.logger.Error("failed to insert user", zap.Error(err))
		return nil, fmt.Errorf("failed to insert user: %v", err)
	}

	// Cache user in Redis
	if err := r.cacheUser(ctx, createdUser); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		r.logger.Error("failed to commit transaction", zap.Error(err))
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return createdUser, nil
}

// UpdateUser updates an existing user
func (r *UserRepository) UpdateUser(ctx context.Context, user *models.UpdateUser) (*models.User, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		r.logger.Error("failed to begin transaction", zap.Error(err))
		return nil, fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	updateMap := make(map[string]interface{})

	if user.ProID != nil {
		updateMap["pro_id"] = *user.ProID
	}
	if user.FirstName != nil {
		updateMap["first_name"] = *user.FirstName
	}
	if user.LastName != nil {
		updateMap["last_name"] = *user.LastName
	}
	if user.Email != nil {
		updateMap["email"] = *user.Email
	}
	if user.DateOfBirth != nil {
		updateMap["date_of_birth"] = user.DateOfBirth
	}
	if user.Gender != nil {
		updateMap["gender"] = *user.Gender
	}
	if user.Avatar != nil {
		updateMap["avatar"] = *user.Avatar
	}
	if user.Role != nil {
		updateMap["role"] = *user.Role
	}

	updateMap["updated_at"] = time.Now()

	query := r.sqlBuilder.
		Update("users").
		SetMap(updateMap).
		Where(sq.Eq{"id": user.ID}).
		Suffix("RETURNING *")

	sqlq, args, err := query.ToSql()
	if err != nil {
		r.logger.Error("failed to build update query", zap.Error(err))
		return nil, fmt.Errorf("failed to build update query: %v", err)
	}

	updatedUser := &models.User{}
	err = tx.QueryRowContext(ctx, sqlq, args...).Scan(
		&updatedUser.ID, &updatedUser.ProID, &updatedUser.FirstName,
		&updatedUser.LastName, &updatedUser.Email, &updatedUser.DateOfBirth,
		&updatedUser.Gender, &updatedUser.Avatar, &updatedUser.Role,
		&updatedUser.CreatedAt, &updatedUser.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		r.logger.Error("failed to update user", zap.Error(models.ErrUserNotFound))
		return nil, models.ErrUserNotFound
	}
	if err != nil {
		r.logger.Error("failed to update user", zap.Error(err))
		return nil, fmt.Errorf("failed to update user: %v", err)
	}

	// Update Redis cache
	if err := r.cacheUser(ctx, updatedUser); err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		r.logger.Error("failed to update user", zap.Error(err))
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return updatedUser, nil
}

// DeleteUser removes a user by ID
func (r *UserRepository) DeleteUser(ctx context.Context, id int) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		r.logger.Error("failed to begin transaction", zap.Error(err))
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	query := r.sqlBuilder.
		Delete("users").
		Where(sq.Eq{"id": id})

	sqlq, args, err := query.ToSql()
	if err != nil {
		r.logger.Error("failed to build delete query", zap.Error(err))
		return fmt.Errorf("failed to build delete query: %v", err)
	}

	_, err = tx.ExecContext(ctx, sqlq, args...)
	if err != nil {
		r.logger.Error("failed to delete user", zap.Error(err))
		return fmt.Errorf("failed to delete user: %v", err)
	}
	if err == sql.ErrNoRows {
		r.logger.Error("failed to delete user", zap.Error(models.ErrUserNotFound))
		return models.ErrUserNotFound
	}
	// Remove from Redis cache
	cacheKey := fmt.Sprintf("user:%d", id)
	if err := r.redis.Del(ctx, cacheKey).Err(); err != nil {
		r.logger.Error("failed to delete user from cache", zap.Error(err))
		return fmt.Errorf("failed to delete user from cache: %v", err)
	}

	return tx.Commit()
}

// GetUser retrieves a user by ID with cache
func (r *UserRepository) GetUser(ctx context.Context, id int) (*models.User, error) {
	// Check Redis cache first
	cacheKey := fmt.Sprintf("user:%d", id)
	cachedUser, err := r.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var user models.User
		if err := json.Unmarshal([]byte(cachedUser), &user); err == nil {
			return &user, nil
		}
	}

	// If not in cache, query database
	query := r.sqlBuilder.
		Select("*").
		From("users").
		Where(sq.Eq{"id": id})

	sqlq, args, err := query.ToSql()
	if err != nil {
		r.logger.Error("failed to delete user from cache", zap.Error(err))
		return nil, fmt.Errorf("failed to build query: %v", err)
	}

	user := &models.User{}
	err = r.db.QueryRowContext(ctx, sqlq, args...).Scan(
		&user.ID, &user.ProID, &user.FirstName,
		&user.LastName, &user.Email, &user.DateOfBirth,
		&user.Gender, &user.Avatar, &user.Role,
		&user.CreatedAt, &user.UpdatedAt, &user.Phone,
	)

	if err == sql.ErrNoRows {
		r.logger.Error("failed to get user", zap.Error(models.ErrUserNotFound))
		return nil, models.ErrUserNotFound
	}
	if err != nil {
		r.logger.Error("failed to get user", zap.Error(err))
		return nil, fmt.Errorf("failed to get user: %v", err)
	}

	// Cache the user
	if err := r.cacheUser(ctx, user); err != nil {
		r.logger.Error("failed to cache user", zap.Error(err))
		return nil, err
	}

	return user, nil
}

// GetAllUsers retrieves users with complex filtering
func (r *UserRepository) GetAllUsers(ctx context.Context, filter *models.GetAllUsers) (*models.GetAllUsersRes, error) {
	query := r.sqlBuilder.Select("*").From("users")

	// Apply filters
	if filter.FirstName != nil {
		query = query.Where(sq.ILike{"first_name": "%" + *filter.FirstName + "%"})
	}
	if filter.LastName != nil {
		query = query.Where(sq.ILike{"last_name": "%" + *filter.LastName + "%"})
	}
	if filter.Gender != nil {
		query = query.Where(sq.Eq{"gender": *filter.Gender})
	}
	if filter.Role != nil {
		query = query.Where(sq.Eq{"role": *filter.Role})
	}
	if filter.BirthRange != nil && filter.BirthRange.IsValid() {
		query = query.Where(
			sq.And{
				sq.GtOrEq{"date_of_birth": filter.BirthRange.StartDate},
				sq.LtOrEq{"date_of_birth": filter.BirthRange.EndDate},
			},
		)
	}

	// Pagination
	if filter.Limit > 0 {
		query = query.Limit(uint64(filter.Limit))
	}
	query = query.Offset(uint64(filter.Offset))

	// Get total count
	countQuery := r.sqlBuilder.Select("COUNT(*)").From("users")
	sql, args, err := countQuery.ToSql()
	if err != nil {
		r.logger.Error("failed to cache user", zap.Error(err))
		return nil, fmt.Errorf("failed to build count query: %v", err)
	}

	var totalItems int
	err = r.db.QueryRowContext(ctx, sql, args...).Scan(&totalItems)
	if err != nil {
		r.logger.Error("failed to get total count", zap.Error(err))
		return nil, fmt.Errorf("failed to get total count: %v", err)
	}

	// Execute main query
	sql, args, err = query.ToSql()
	if err != nil {
		r.logger.Error("failed to build query", zap.Error(err))
		return nil, fmt.Errorf("failed to build query: %v", err)
	}

	rows, err := r.db.QueryContext(ctx, sql, args...)
	if err != nil {
		r.logger.Error("failed to query users", zap.Error(err))
		return nil, fmt.Errorf("failed to query users: %v", err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(
			&user.ID, &user.ProID, &user.FirstName,
			&user.LastName, &user.Email, &user.DateOfBirth,
			&user.Gender, &user.Avatar, &user.Role,
			&user.CreatedAt, &user.UpdatedAt, &user.Phone,
		)
		if err != nil {
			r.logger.Error("failed to scan user", zap.Error(err))
			return nil, fmt.Errorf("failed to scan user: %v", err)
		}
		users = append(users, user)
	}

	return &models.GetAllUsersRes{
		TotalItems: totalItems,
		Users:      users,
	}, nil
}

// GetUserByEmail retrieves a user by email
func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	query := r.sqlBuilder.
		Select("*").
		From("users").
		Where(sq.Eq{"email": email})

	sqlq, args, err := query.ToSql()
	if err != nil {
		r.logger.Error("failed to build query", zap.Error(err))
		return nil, fmt.Errorf("failed to build query: %v", err)
	}

	user := &models.User{}
	err = r.db.QueryRowContext(ctx, sqlq, args...).Scan(
		&user.ID, &user.ProID, &user.FirstName,
		&user.LastName, &user.Email, &user.DateOfBirth,
		&user.Gender, &user.Avatar, &user.Role,
		&user.CreatedAt, &user.UpdatedAt, &user.Phone,
	)
	if err == sql.ErrNoRows {
		r.logger.Error("No user found by this email", zap.Error(err))
		return nil, models.ErrUserNotFound
	}
	if err != nil {
		r.logger.Error("failed to get user by email", zap.Error(err))
		return nil, fmt.Errorf("failed to get user by email: %v", err)
	}

	return user, nil
}

// SearchUsers performs a full-text search on users
func (r *UserRepository) SearchUsers(ctx context.Context, query string, limit, offset int) (*models.GetAllUsersRes, error) {
	searchQuery := r.sqlBuilder.
		Select("*").
		From("users").
		Where(
			sq.Or{
				sq.ILike{"first_name": "%" + query + "%"},
				sq.ILike{"last_name": "%" + query + "%"},
				sq.ILike{"email": "%" + query + "%"},
			},
		).
		Limit(uint64(limit)).
		Offset(uint64(offset))

	// Get total count
	countQuery := r.sqlBuilder.
		Select("COUNT(*)").
		From("users").
		Where(
			sq.Or{
				sq.ILike{"first_name": "%" + query + "%"},
				sq.ILike{"last_name": "%" + query + "%"},
				sq.ILike{"email": "%" + query + "%"},
			},
		)

	sql, args, err := countQuery.ToSql()
	if err != nil {
		r.logger.Error("failed to build count query", zap.Error(err))
		return nil, fmt.Errorf("failed to build count query: %v", err)
	}

	var totalItems int
	err = r.db.QueryRowContext(ctx, sql, args...).Scan(&totalItems)
	if err != nil {
		r.logger.Error("failed to get total count", zap.Error(err))
		return nil, fmt.Errorf("failed to get total count: %v", err)
	}

	sql, args, err = searchQuery.ToSql()
	if err != nil {
		r.logger.Error("failed to build query", zap.Error(err))

		return nil, fmt.Errorf("failed to build query: %v", err)
	}

	rows, err := r.db.QueryContext(ctx, sql, args...)
	if err != nil {
		r.logger.Error("failed to search users", zap.Error(err))
		return nil, fmt.Errorf("failed to search users: %v", err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(
			&user.ID, &user.ProID, &user.FirstName,
			&user.LastName, &user.Email, &user.DateOfBirth,
			&user.Gender, &user.Avatar, &user.Role,
			&user.CreatedAt, &user.UpdatedAt, &user.Phone,
		)
		if err != nil {
			r.logger.Error("failed to scan user", zap.Error(err))
			return nil, fmt.Errorf("failed to scan user: %v", err)
		}
		users = append(users, user)
	}

	return &models.GetAllUsersRes{
		TotalItems: totalItems,
		Users:      users,
	}, nil
}

// GetUsersByRole (continued)
func (r *UserRepository) GetUsersByRole(ctx context.Context, role string) (*models.GetAllUsersRes, error) {
	rows, err := r.db.QueryContext(ctx,
		"SELECT * FROM users WHERE role = $1",
		role,
	)
	if err != nil {
		r.logger.Error("failed to query users by role", zap.Error(err))
		return nil, fmt.Errorf("failed to query users by role: %v", err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(
			&user.ID, &user.ProID, &user.FirstName,
			&user.LastName, &user.Email, &user.DateOfBirth,
			&user.Gender, &user.Avatar, &user.Role,
			&user.CreatedAt, &user.UpdatedAt, &user.Phone,
		)
		if err != nil {
			r.logger.Error("failed to scan user", zap.Error(err))
			return nil, fmt.Errorf("failed to scan user: %v", err)
		}
		users = append(users, user)
	}

	return &models.GetAllUsersRes{
		TotalItems: len(users),
		Users:      users,
	}, nil
}

// GetUsersByDateRange retrieves users within a specific date range
func (r *UserRepository) GetUsersByDateRange(ctx context.Context, startDate, endDate time.Time) (*models.GetAllUsersRes, error) {
	query := r.sqlBuilder.
		Select("*").
		From("users").
		Where(
			sq.And{
				sq.GtOrEq{"date_of_birth": startDate},
				sq.LtOrEq{"date_of_birth": endDate},
			},
		)

	sql, args, err := query.ToSql()
	if err != nil {
		r.logger.Error("failed to build query", zap.Error(err))
		return nil, fmt.Errorf("failed to build query: %v", err)
	}

	rows, err := r.db.QueryContext(ctx, sql, args...)
	if err != nil {
		r.logger.Error("failed to query users by date range", zap.Error(err))
		return nil, fmt.Errorf("failed to query users by date range: %v", err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(
			&user.ID, &user.ProID, &user.FirstName,
			&user.LastName, &user.Email, &user.DateOfBirth,
			&user.Gender, &user.Avatar, &user.Role,
			&user.CreatedAt, &user.UpdatedAt, &user.Phone,
		)
		if err != nil {
			r.logger.Error("failed to scan user", zap.Error(err))
			return nil, fmt.Errorf("failed to scan user: %v", err)
		}
		users = append(users, user)
	}

	return &models.GetAllUsersRes{
		TotalItems: len(users),
		Users:      users,
	}, nil
}

// ActivateUser sets user status to active
func (r *UserRepository) ActivateUser(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx,
		"UPDATE users SET status = 'active' WHERE id = $1",
		id,
	)
	return err
}

// DeactivateUser sets user status to inactive
func (r *UserRepository) DeactivateUser(ctx context.Context, id int) error {
	_, err := r.db.ExecContext(ctx,
		"UPDATE users SET status = 'inactive' WHERE id = $1",
		id,
	)
	return err
}

// ToggleUserStatus switches between active and inactive
func (r *UserRepository) ToggleUserStatus(ctx context.Context, id int) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		r.logger.Error("failed to begin transaction", zap.Error(err))
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	var currentStatus string
	err = tx.QueryRowContext(ctx, "SELECT status FROM users WHERE id = $1", id).Scan(&currentStatus)
	if err != nil {
		r.logger.Error("failed to get current status", zap.Error(err))
		return fmt.Errorf("failed to get current status: %v", err)
	}

	newStatus := "inactive"
	if currentStatus == "inactive" {
		newStatus = "active"
	}

	_, err = tx.ExecContext(ctx,
		"UPDATE users SET status = $1 WHERE id = $2",
		newStatus, id,
	)
	if err != nil {
		r.logger.Error("failed to toggle user status", zap.Error(err))
		return fmt.Errorf("failed to toggle user status: %v", err)
	}

	return tx.Commit()
}

// UpdatePassword updates user's password hash
func (r *UserRepository) UpdatePassword(ctx context.Context, id int, hashedPassword string) error {
	_, err := r.db.ExecContext(ctx,
		"UPDATE users SET password_hash = $1 WHERE id = $2",
		hashedPassword, id,
	)
	return err
}

// LogLastLogin records the timestamp of last user login
func (r *UserRepository) LogLastLogin(ctx context.Context, id int, loginTime time.Time) error {
	_, err := r.db.ExecContext(ctx,
		"UPDATE users SET last_login = $1 WHERE id = $2",
		loginTime, id,
	)
	return err
}

// GetRecentlyCreatedUsers retrieves recently created users
func (r *UserRepository) GetRecentlyCreatedUsers(ctx context.Context, limit, offset int) (*models.GetAllUsersRes, error) {
	query := r.sqlBuilder.
		Select("*").
		From("users").
		OrderBy("created_at DESC").
		Limit(uint64(limit)).
		Offset(uint64(offset))

	// Get total count
	countQuery := r.sqlBuilder.
		Select("COUNT(*)").
		From("users")

	sql, args, err := countQuery.ToSql()
	if err != nil {
		r.logger.Error("failed to build count query", zap.Error(err))
		return nil, fmt.Errorf("failed to build count query: %v", err)
	}

	var totalItems int
	err = r.db.QueryRowContext(ctx, sql, args...).Scan(&totalItems)
	if err != nil {
		r.logger.Error("failed to get total count", zap.Error(err))
		return nil, fmt.Errorf("failed to get total count: %v", err)
	}

	sql, args, err = query.ToSql()
	if err != nil {
		r.logger.Error("failed to build query", zap.Error(err))
		return nil, fmt.Errorf("failed to build query: %v", err)
	}

	rows, err := r.db.QueryContext(ctx, sql, args...)
	if err != nil {
		r.logger.Error("failed to query recent users", zap.Error(err))
		return nil, fmt.Errorf("failed to query recent users: %v", err)
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(
			&user.ID, &user.ProID, &user.FirstName,
			&user.LastName, &user.Email, &user.DateOfBirth,
			&user.Gender, &user.Avatar, &user.Role,
			&user.CreatedAt, &user.UpdatedAt, &user.Phone,
		)
		if err != nil {
			r.logger.Error("failed to build count query", zap.Error(err))
			return nil, fmt.Errorf("failed to build count query: %v", err)
		}
		users = append(users, user)
	}

	return &models.GetAllUsersRes{
		TotalItems: totalItems,
		Users:      users,
	}, nil
}

// AssignProID assigns a project ID to a user
func (r *UserRepository) AssignProID(ctx context.Context, userID, proID int) error {
	_, err := r.db.ExecContext(ctx,
		"UPDATE users SET pro_id = $1 WHERE id = $2",
		proID, userID,
	)
	return err
}

// Helper method to cache user in Redis
func (r *UserRepository) cacheUser(ctx context.Context, user *models.User) error {
	cacheKey := fmt.Sprintf("user:%d", user.ID)
	value, err := json.Marshal(user)
	if err != nil {
		r.logger.Error("failed to marshal user", zap.Error(err))
		return fmt.Errorf("failed to marshal user: %v", err)
	}

	return r.redis.Set(ctx, cacheKey, value, time.Hour*2).Err()
}

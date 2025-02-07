/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-12-17
 * @Description: Handlers for UserManagement gRPC Service
 */

package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/projects/pro-sphere-backend/apigateway/genproto/auth"
	"github.com/projects/pro-sphere-backend/apigateway/internal/items/config"
	"go.uber.org/zap"
)

type UserHandler struct {
	cfg    *config.Config
	logger *zap.Logger
	user   auth.UserManagementClient
}

func NewUserHandler(cfg *config.Config, logger *zap.Logger, userClient auth.UserManagementClient) *UserHandler {
	return &UserHandler{
		cfg:    cfg,
		logger: logger,
		user:   userClient,
	}
}

// @Summary Create a new user
// @Description Creates a new user in the system
// @Tags User Management
// @Accept json
// @Produce json
// @Security     BearerAuth
// @Param user body auth.CreateUserRequest true "User data"
// @Success 201 {object} auth.CreateUserResponse "User created successfully"
// @Failure 400 {object} map[string]string "Invalid user data"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /v1/users [post]
func (u *UserHandler) CreateUser(c *gin.Context) {
	var req auth.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		u.logger.Warn("Invalid request body", zap.Error(err))
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	res, err := u.user.CreateUser(c, &req)
	if err != nil {
		u.logger.Error("Failed to create user", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.IndentedJSON(http.StatusCreated, res)
}

// @Summary Get user by ID
// @Description Retrieves a user's details by ID
// @Tags User Management
// @Produce json
// @Security     BearerAuth
// @Success 200 {object} auth.GetUserResponse "User data"
// @Failure 404 {object} map[string]string "User not found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /v1/users/me [get]
func (u *UserHandler) GetUserByID(c *gin.Context) {
	// Retrieve the user ID from the context
	userID, exists := c.Get("user_id")
	if !exists {
		u.logger.Warn("User ID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	id, ok := userID.(int) // Type assertion to int
	if !ok {
		u.logger.Error("Failed to assert user_id to int")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server error"})
		return
	}

	// Call the gRPC service
	res, err := u.user.GetUser(c, &auth.GetUserRequest{Id: int32(id)})
	if err != nil {
		u.logger.Error("Failed to retrieve user", zap.Int("id", id), zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary Delete user by ID
// @Description Deletes a user by ID
// @Tags User Management
// @Produce json
// @Security     BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} map[string]string "User deleted successfully"
// @Failure 404 {object} map[string]string "User not found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /v1/users/{id} [delete]
func (u *UserHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id") // or c.Query("id") for query parameter
	if idStr == "" {
		u.logger.Warn("ID parameter missing")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	// Convert `id` from string to int32
	id, err := strconv.Atoi(idStr)
	if err != nil || id > 2147483647 || id < -2147483648 { // Check for int32 range
		u.logger.Error("Invalid ID parameter", zap.String("id", idStr), zap.Error(err))
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "ID must be a valid 32-bit integer"})
		return
	}
	res, err := u.user.DeleteUser(c, &auth.DeleteUserRequest{Id: int32(id)})
	if err != nil {
		u.logger.Error("Failed to delete user", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}

// @Summary Update user by ID
// @Description Updates a user's information by ID
// @Tags User Management
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body auth.UpdateUserByIDRequest true "Updated user data"
// @Security     BearerAuth
// @Success 200 {object} auth.UpdateUserResponse "User updated successfully"
// @Failure 400 {object} map[string]string "Invalid user data"
// @Failure 404 {object} map[string]string "User not found"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /v1/users/{id} [put]
func (u *UserHandler) UpdateUserByID(c *gin.Context) {
	var req auth.UpdateUserByIDRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		u.logger.Warn("Invalid request body", zap.Error(err))
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
		return
	}

	res, err := u.user.UpdateUserByID(c, &req)
	if err != nil {
		u.logger.Error("Failed to update user", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}

// @Summary Get all users
// @Description Retrieves a list of all users with optional filters
// @Tags User Management
// @Produce json
// @Security     BearerAuth
// @Param first_name query string false "First Name"
// @Param last_name query string false "Last Name"
// @Param gender query string false "Gender"
// @Param role query string false "Role"
// @Param limit query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {object} auth.GetAllUsersResponse "List of users"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /v1/users [get]
func (u *UserHandler) GetAllUsers(c *gin.Context) {
	var req auth.GetAllUsersRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		u.logger.Warn("Invalid query parameters", zap.Error(err))
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid query parameters"})
		return
	}

	res, err := u.user.GetAllUsers(c, &req)
	if err != nil {
		u.logger.Error("Failed to get all users", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}

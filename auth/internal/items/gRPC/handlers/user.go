package handlers

import (
	"context"
	"log"
	"time"

	"github.com/projects/pro-sphere-backend/auth/genproto/genproto/auth"
	"github.com/projects/pro-sphere-backend/auth/internal/items/models"
	"github.com/projects/pro-sphere-backend/auth/internal/items/service"
	"go.uber.org/zap"
)

type UserHandler struct {
	service *service.UserService
	logger  *zap.Logger
	auth.UnimplementedUserManagementServer
}

func NewUserHandler(service *service.UserService, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		service: service,
		logger:  logger,
	}
}

// GetUserByID handles the GetUserByID gRPC request
func (h *UserHandler) GetUser(ctx context.Context, req *auth.GetUserRequest) (*auth.GetUserResponse, error) {
	user, err := h.service.GetUserByID(ctx, int(req.GetId()))
	if err != nil {
		h.logger.Error("failed to get user by ID", zap.Error(err))
		return nil, err
	}

	return &auth.GetUserResponse{
		User: convertModelToProto(user),
	}, nil
}

// CreateUser handles the CreateUser gRPC request
func (h *UserHandler) CreateUser(ctx context.Context, req *auth.CreateUserRequest) (*auth.CreateUserResponse, error) {
	input := convertProtoToCreateModel(req.GetNewUser())
	user, err := h.service.CreateUser(ctx, input)
	if err != nil {
		h.logger.Error("failed to create user", zap.Error(err))
		return nil, err
	}

	return &auth.CreateUserResponse{
		User: convertModelToProto(user),
	}, nil
}

// UpdateUserByID handles the UpdateUserByID gRPC request
func (h *UserHandler) UpdateUserByID(ctx context.Context, req *auth.UpdateUserByIDRequest) (*auth.UpdateUserResponse, error) {
	input := convertProtoToUpdateModel(req.GetUpdateUser())
	user, err := h.service.UpdateUser(ctx, input)
	if err != nil {
		h.logger.Error("failed to update user by ID", zap.Error(err))
		return nil, err
	}

	return &auth.UpdateUserResponse{
		User: convertModelToProto(user),
	}, nil
}

// DeleteUser handles the DeleteUser gRPC request
func (h *UserHandler) DeleteUser(ctx context.Context, req *auth.DeleteUserRequest) (*auth.DeleteUserResponse, error) {
	err := h.service.DeleteUser(ctx, int(req.GetId()))
	if err != nil {
		h.logger.Error("failed to delete user", zap.Error(err))
		return nil, err
	}

	return &auth.DeleteUserResponse{
		Success: true,
	}, nil
}

// GetUserByProID handles the GetUserByProID gRPC request
func (h *UserHandler) GetUserByPROID(ctx context.Context, req *auth.GetUserByProIDRequest) (*auth.GetUserResponse, error) {
	user, err := h.service.GetUserByProID(ctx, int(req.GetProId()))
	if err != nil {
		h.logger.Error("failed to get user by ProID", zap.Error(err))
		return nil, err
	}

	return &auth.GetUserResponse{
		User: convertModelToProto(user),
	}, nil
}

// GetAllUsers handles the GetAllUsers gRPC request.
func (h *UserHandler) GetAllUsers(ctx context.Context, req *auth.GetAllUsersRequest) (*auth.GetAllUsersResponse, error) {
	// Convert gRPC request to service filter model
	filter := convertProtoToFilterModel(req)

	// Call the service layer
	result, err := h.service.GetAllUsers(ctx, filter)
	if err != nil {
		h.logger.Error("failed to get all users", zap.Error(err))
		return nil, err
	}

	// Convert service response to gRPC response
	return convertServiceToProtoResponse(result), nil
}
func convertServiceToProtoResponse(res *models.GetAllUsersRes) *auth.GetAllUsersResponse {
	users := make([]*auth.User, len(res.Users))
	for i, user := range res.Users {
		users[i] = convertModelToProto(user)
	}

	return &auth.GetAllUsersResponse{
		TotalItems: int32(res.TotalItems),
		Users:      users,
	}
}

func convertModelToProto(user *models.User) *auth.User {
	if user == nil {
		return nil
	}
	return &auth.User{
		Id:          int32(user.ID),
		ProId:       int32(user.ProID),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Phone:       user.Phone,
		DateOfBirth: user.DateOfBirth,
		Gender:      string(user.Gender),
		Avatar:      user.Avatar,
		Role:        user.Role,
		CreatedAt:   user.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   user.UpdatedAt.Format(time.RFC3339),
	}
}

func convertProtoToCreateModel(proto *auth.User) *models.CreateUser {
	return &models.CreateUser{
		ProID:       int(proto.ProId),
		FirstName:   proto.FirstName,
		LastName:    proto.LastName,
		Email:       proto.Email,
		Phone:       proto.Phone,
		DateOfBirth: proto.DateOfBirth,
		Gender:      models.Gender(proto.Gender),
		Avatar:      proto.Avatar,
		Role:        proto.Role,
	}
}

func convertProtoToUpdateModel(proto *auth.User) *models.UpdateUser {
	return &models.UpdateUser{
		ID:          int(proto.GetId()),                // Convert int32 to *int
		ProID:       intPointer(int(proto.GetProId())), // Convert int32 to *int
		FirstName:   stringPointer(proto.GetFirstName()),
		LastName:    stringPointer(proto.GetLastName()),
		Email:       stringPointer(proto.GetEmail()),
		DateOfBirth: stringPointer(proto.GetDateOfBirth()),
		Gender:      (*models.Gender)(&proto.Gender),
		Avatar:      stringPointer(proto.GetAvatar()),
		Role:        stringPointer(proto.GetRole()),
	}
}
func intPointer(i int) *int {
	return &i
}

func stringPointer(s string) *string {
	return &s
}

func convertProtoToFilterModel(req *auth.GetAllUsersRequest) *models.GetAllUsers {
	return &models.GetAllUsers{
		FirstName:  stringPointer(req.GetFirstName()),
		LastName:   stringPointer(req.GetLastName()),
		Gender:     (*models.Gender)(&req.Gender),
		Role:       stringPointer(req.GetRole()),
		BirthRange: convertProtoToBirthRange(req.GetBirthRange()),
		Limit:      int(req.GetLimit()),
		Offset:     int(req.GetOffset()),
	}
}

func convertProtoToBirthRange(protoRange *auth.BirthRange) *models.BirthRange {
	if protoRange == nil {
		return nil
	}

	// Parse the StartDate
	startDateStr := protoRange.GetStartDate()
	const layout = time.RFC3339
	startDate, err := time.Parse(layout, startDateStr)
	if err != nil {
		log.Fatalf("Error parsing start date: %v", err)
	}

	// Parse the EndDate (if needed)
	endDateStr := protoRange.GetEndDate()
	var endDate time.Time
	if endDateStr != "" {
		endDate, err = time.Parse(layout, endDateStr)
		if err != nil {
			log.Fatalf("Error parsing end date: %v", err)
		}
	}

	// Return the converted BirthRange
	return &models.BirthRange{
		StartDate: startDate, // Use the parsed time.Time
		EndDate:   endDate,   // Use the parsed time.Time
	}
}

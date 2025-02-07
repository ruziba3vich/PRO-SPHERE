/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-12-05 02:39:29
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-26 01:54:54
 * @FilePath: /auth/internal/items/service/auth.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/projects/pro-sphere-backend/auth/genproto/genproto/auth"
	"github.com/projects/pro-sphere-backend/auth/internal/items/config"
	"github.com/projects/pro-sphere-backend/auth/internal/items/models"
	"github.com/projects/pro-sphere-backend/auth/internal/items/repos"
	"github.com/projects/pro-sphere-backend/auth/internal/items/storage/cache"
	"go.uber.org/zap"
)

type (
	AuthService struct {
		userRepo  repos.UserRepositories
		cfg       *config.Config
		logger    *zap.Logger
		userCache *cache.UserCache
	}
)

func NewAuthService(userRepo repos.UserRepositories, userCache *cache.UserCache, cfg *config.Config, logger *zap.Logger) *AuthService {
	return &AuthService{
		userRepo:  userRepo,
		logger:    logger,
		cfg:       cfg,
		userCache: userCache,
	}
}

func (s *AuthService) LoginUser(ctx context.Context, code, clientName string) (*string, error) {
	s.logger.Info("Starting user login process", zap.String("code", code))

	tokens, err := s.ExchangeCodeForToken(code, clientName)
	if err != nil {
		s.logger.Error("Failed to exchange code for token", zap.String("code", code), zap.Error(err))
		return nil, fmt.Errorf("failed to exchange code for token: %w", err)
	}

	proIDToken := tokens.AccessToken
	s.logger.Info("Successfully exchanged code for token", zap.String("access_token", proIDToken))

	proUser, err := s.FetchUserProfile(proIDToken)
	if err != nil {
		s.logger.Error("Failed to fetch user profile", zap.String("access_token", proIDToken), zap.Error(err))
		return nil, fmt.Errorf("failed to fetch user profile: %w", err)
	}

	user, err := s.userRepo.GetUserByPROID(ctx, proUser.ID)
	if err == sql.ErrNoRows || user == nil {
		s.logger.Info("User not found, creating a new user", zap.Any("pro_id", proUser.ID))

		user, err = s.userRepo.CreateUser(ctx, &models.CreateUser{
			ProID:       proUser.ID,
			FirstName:   proUser.FirstName,
			LastName:    proUser.LastName,
			Email:       proUser.Email,
			Phone:       proUser.PhoneNumber,
			DateOfBirth: proUser.DateOfBirth,
			Gender:      proUser.Gender,
			Avatar:      proUser.AvatarUrl,
			Role:        "user",
		})
		if err != nil {
			s.logger.Error("Failed to create new user", zap.Error(err))
			return nil, fmt.Errorf("failed to create new user: %w", err)
		}
		s.logger.Info("Successfully created new user", zap.Any("user_id", user.ID))
	}

	jwtToken, err := GenerateJWT(user.ID, user.ProID, []byte(s.cfg.ApiKey))
	if err != nil {
		s.logger.Error("Failed to generate JWT token", zap.Error(err))
		return nil, fmt.Errorf("failed to generate JWT token: %w", err)
	}
	s.logger.Info("Successfully generated JWT token", zap.Any("user_id", user.ID))

	userCode := uuid.NewString()
	if err := s.userCache.SaveUserWithCode(ctx, &models.Tokens{AccessToken: jwtToken}, userCode); err != nil {
		s.logger.Error("Failed to cache user with tokens", zap.String("user_code", userCode), zap.Error(err))
		return nil, fmt.Errorf("failed to cache user with tokens: %w", err)
	}
	s.logger.Info("Successfully cached user with tokens", zap.String("user_code", userCode))

	s.logger.Info("User login completed successfully", zap.String("user_code", userCode))

	return &userCode, nil
}

func (s *AuthService) GetTokensByCode(ctx context.Context, code string) (*auth.TokenResponse, error) {
	tokens, err := s.userCache.GetUserTokensByCode(ctx, code)
	if err != nil {
		s.logger.Error("Failed to get tokens from cache", zap.Error(err))
		return nil, err
	}

	return &auth.TokenResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		ExpiresIn:    tokens.ExpiresIn,
	}, nil
}

func (s *AuthService) GenerateOAuth2AuthorizationURL(clientId, redirectUrl string) (string, error) {
	baseURL := s.cfg.ProID.BaseURL
	params := url.Values{}
	params.Add("client_id", clientId)
	params.Add("redirect_uri", redirectUrl)
	params.Add("response_type", "code")

	return fmt.Sprintf("%s?%s", baseURL, params.Encode()), nil
}

func (s *AuthService) ExchangeCodeForToken(code, clientName string) (*models.TokenResponse, error) {
	endpoint := s.cfg.ProID.Endpoint
	data := url.Values{}
	data.Add("grant_type", s.cfg.ProID.GrantType)
	data.Add("client_id", s.cfg.ProIDOther[clientName].ClientID)
	data.Add("client_secret", s.cfg.ProIDOther[clientName].SecretKey)
	data.Add("redirect_uri", s.cfg.ProIDOther[clientName].RedirectURL)
	data.Add("code", code)

	resp, err := http.PostForm(endpoint, data)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to exchange code for token: %s", resp.Status)
	}

	var tokenResponse models.TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return nil, err
	}

	return &tokenResponse, nil
}

func (s *AuthService) FetchUserProfile(accessToken string) (*models.ProIDUser, error) {
	client := &http.Client{
		Timeout: 10 * time.Second, // Set a timeout
	}

	req, err := http.NewRequest("GET", "https://api.id.sfere.pro/api/user", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch user profile: HTTP %d (%s)", resp.StatusCode, http.StatusText(resp.StatusCode))
	}
	var user models.ProIDUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	s.logger.Info("User: ", zap.Any("User", user))
	user.DateOfBirth = "1970-01-01"
	user.Gender = models.Gender(models.NormalizeGender(string(user.Gender)))

	user.DateOfBirth = "1970-01-01"

	s.logger.Info("User: ", zap.Any("User", user))

	return &user, nil
}

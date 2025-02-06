/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-12-08 23:28:44
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-26 01:59:00
 * @FilePath: /auth/internal/items/gRPC/handlers/auth.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package handlers

import (
	"context"

	"github.com/projects/pro-sphere-backend/auth/genproto/genproto/auth"
	"github.com/projects/pro-sphere-backend/auth/internal/items/config"
	"github.com/projects/pro-sphere-backend/auth/internal/items/repos"
	"github.com/projects/pro-sphere-backend/auth/internal/items/service"
	"go.uber.org/zap"
)

type (
	AuthHandler struct {
		authService *service.AuthService
		userRepo    repos.UserRepositories
		logger      *zap.Logger
		cfg         *config.Config
		auth.UnimplementedAuthenticationServer
	}
)

func NewAuthHandler(authService *service.AuthService, userRepo repos.UserRepositories, logger *zap.Logger, cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		userRepo:    userRepo,
		authService: authService,
		logger:      logger,
		cfg:         cfg,
	}
}

func (a *AuthHandler) GenerateOuthUrl(ctx context.Context, req *auth.GenerateOuthUrlRequest) (*auth.GenerateOuthUrlResponse, error) {
	var url string
	var clientName string
	if req.Name == "admin" {
		clientName = "admin"
	} else if req.Name == "web" {
		clientName = "web"
	} else {
		a.logger.Error("Client name must be admin, web or mobile", zap.String("GotName", clientName))
		return nil, service.ErrInvalidInput
	}

	clientId := a.cfg.ProIDOther[clientName].ClientID
	redirectUrl := a.cfg.ProIDOther[clientName].RedirectURL

	url, err := a.authService.GenerateOAuth2AuthorizationURL(clientId, redirectUrl)
	if err != nil {
		a.logger.Error("Failed to generate url for proid", zap.Error(err))
		return nil, err
	}

	return &auth.GenerateOuthUrlResponse{
		Url: url,
	}, nil
}

func (a *AuthHandler) LoginUser(ctx context.Context, req *auth.CodeRequest) (*auth.CodeResponse, error) {
	clients := make(map[string]string, 4)
	clients["web"] = ""
	clients["admin"] = ""
	clients["mobile"] = ""
	clients["news"] = ""

	if cl, ok := clients[req.ClientName]; !ok {
		a.logger.Error("Client must be named admin, web, mobile or news", zap.String("gotname", cl))
		return nil, service.ErrInvalidInput
	}
	code, err := a.authService.LoginUser(ctx, req.ProCode, req.ClientName)
	if err != nil {
		a.logger.Error("Failed to exchage code for token", zap.Error(err))
		return nil, err
	}
	// fmt.Println(tokens)
	a.logger.Info("code", zap.Any("code", *code))

	return &auth.CodeResponse{Code: *code}, nil
}

func (a *AuthHandler) GetTokens(ctx context.Context, req *auth.TokenRequest) (*auth.TokenResponse, error) {
	return a.authService.GetTokensByCode(ctx, req.Code)
}

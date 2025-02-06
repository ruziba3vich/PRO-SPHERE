/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-12-08 22:34:50
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-26 01:49:05
 * @FilePath: /apigateway/internal/items/http/handlers/auth.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-12-08 22:34:50
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-18 13:57:00
 * @FilePath: /apigateway/internal/items/http/handlers/auth.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/projects/pro-sphere-backend/apigateway/genproto/auth"
	"github.com/projects/pro-sphere-backend/apigateway/internal/items/config"
	"go.uber.org/zap"
)

type AuthHandler struct {
	cfg    *config.Config
	logger *zap.Logger
	auth   auth.AuthenticationClient
}

func NewAuthHandler(cfg *config.Config, logger *zap.Logger, authClient auth.AuthenticationClient) *AuthHandler {
	return &AuthHandler{
		cfg:    cfg,
		logger: logger,
		auth:   authClient,
	}
}

// @Summary Start OAuth flow and redirect to Pro ID
// @Description Initiates the OAuth process and redirects the user to the Pro ID login page
// @Tags Authentication
// @Produce json
// Param client_name query string true "client name e.g(admin, web, mobile or news(soon))"
// @Success 302 {string} string "Redirect URL"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /v1/auth/oauth/start [get]
func (a *AuthHandler) StartAndRedirectToProID(c *gin.Context) {
	clientName := c.Query("client_name")
	if clientName == "" {
		clientName = "web"
		a.logger.Warn("Client name is not provided: changed to default -> web")
	}
	res, err := a.auth.GenerateOuthUrl(c, &auth.GenerateOuthUrlRequest{Name: clientName})
	if err == errors.New("invalid input") {
		a.logger.Error("Invalid client name", zap.String("Got", clientName))
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Must be mobile, web, admin or news"})
		return
	}

	if err != nil {
		a.logger.Error("Failed to get response from auth service", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to initiate OAuth flow"})
		return
	}

	c.Redirect(http.StatusFound, res.Url)
}

// @Summary Handle OAuth callback
// @Description Processes the OAuth callback and redirects to the configured URL with a code
// @Tags Authentication
// @Produce json
// @Param code query string true "Authorization Code"
// @Success 302 {string} string "Redirect URL with code"
// @Failure 400 {object} map[string]string "Invalid code"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /v1/auth/oauth/callback [get]
func (a *AuthHandler) HandleCallBack(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		a.logger.Warn("Authorization code missing")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Authorization code is required"})
		return
	}

	res, err := a.auth.LoginUser(c, &auth.CodeRequest{ProCode: code, ClientName: "web"})
	if err != nil {
		a.logger.Error("Failed to exchange code for tokens", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange authorization code"})
		return
	}

	c.Redirect(http.StatusFound, a.cfg.FrontEnd.RedirectURL+res.Code)
}

// @Summary Handle OAuth callback For Admin
// @Description Processes the OAuth callback and redirects to the configured URL with a code
// @Tags Authentication
// @Produce json
// @Param code query string true "Authorization Code"
// @Success 302 {string} string "Redirect URL with code"
// @Failure 400 {object} map[string]string "Invalid code"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /v1/auth/oauth/admin/callback [get]
func (a *AuthHandler) HandleCallBackForAdmin(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		a.logger.Warn("Authorization code missing")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Authorization code is required"})
		return
	}

	res, err := a.auth.LoginUser(c, &auth.CodeRequest{ProCode: code, ClientName: "admin"})
	if err != nil {
		a.logger.Error("Failed to exchange code for tokens", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange authorization code"})
		return
	}

	c.Redirect(http.StatusFound, a.cfg.FrontEndAdmin.RedirectURL+res.Code)
}

// @Summary Get tokens by code
// @Description Retrieves tokens using the provided authorization code
// @Tags Authentication
// @Produce json
// @Param code query string true "Authorization Code"
// @Success 200 {object} auth.TokenResponse "Tokens"
// @Failure 400 {object} map[string]string "Invalid code"
// @Failure 500 {object} map[string]string "Internal Server Error"
// @Router /v1/auth/oauth/tokens [get]
func (a *AuthHandler) GetTokensByCode(c *gin.Context) {
	code := c.Query("code")
	if err := uuid.Validate(code); err != nil {
		a.logger.Warn("Authorization code is not valid! Code must be UUID")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid authorization: must be UUID"})
		return
	}
	if code == "" {
		a.logger.Warn("Authorization code missing")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Authorization code is required"})
		return
	}

	res, err := a.auth.GetTokens(c, &auth.TokenRequest{Code: code})
	if err != nil {
		a.logger.Error("Failed to retrieve tokens", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve tokens"})
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}

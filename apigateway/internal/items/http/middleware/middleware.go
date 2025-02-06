/*
	* @Author: javohir-a abdusamatovjavohir@gmail.com
	* @Date: 2024-10-04 17:31:49
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-17 13:26:14
	* @FilePath: /sfere_backend/internal/items/http/app/middleware.go
*/
package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set CORS headers
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins or specify a domain
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		// Continue to the next middleware or route handler
		c.Next()
	}
}

type Claims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
	ProID  int `json:"pro_id"`
}

func AuthMiddleware(jwtSecretKey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Missing or invalid token"})
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		})

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Invalid token: " + err.Error()})
			return
		}

		ctx.Set("claims", claims)
		ctx.Set("user_id", claims.UserID)
		ctx.Next()
	}
}

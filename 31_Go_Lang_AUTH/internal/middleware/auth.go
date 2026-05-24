package middleware

import (
	"go-auth/internal/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// This file contains the authentication middleware for the application.
//store -> auth data info -> gin context

const (
	CtxUserIDKey = "auth.user_id"
	CtxRoleKey = "auth.role"
)

func AuthRequired(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := strings.TrimSpace(c.GetHeader("Authorization"))
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is required",
			})
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid Authorization header format",
			})
			return
		}

		scheme := strings.TrimSpace(parts[0])
		tokenString := strings.TrimSpace(parts[1])

		if !strings.EqualFold(scheme, "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{	
				"error": "Authorization header must start with Bearer",
			})
			return
		}

		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Token is required",
			})
			return
		}
		claims, err := auth.ParseToken(jwtSecret, tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token: " + err.Error(),
			})
			return
		}
		c.Set(CtxUserIDKey, claims.Subject)
		c.Set(CtxRoleKey, claims.Role)
		c.Next()

	}	
}

func GetUserID(c *gin.Context) (string, bool) {	
	res, ok := c.Get(CtxUserIDKey)
	if !ok {
		return "", false
	}
	userID, ok := res.(string)
	return userID, ok
}

func GetRole(c *gin.Context) (string, bool) {
	res, ok := c.Get(CtxRoleKey)
	if !ok {
		return "", false
	}
	role, ok := res.(string)
	return role, ok
}

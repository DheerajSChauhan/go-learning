package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func RequireRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get role stored by AuthRequired middleware
		role, ok := c.Get(CtxRoleKey)
		if !ok {
			c.AbortWithStatusJSON(403, gin.H{
				"error": "role information is missing",
			})
			return
		}

		roleStr, ok := role.(string)
		if !ok {
			c.AbortWithStatusJSON(403, gin.H{
				"error": "invalid role information",
			})
			return
		}

		// Compare role (case-insensitive)
		if !strings.EqualFold(roleStr, requiredRole) {
			c.AbortWithStatusJSON(403, gin.H{
				"error": "forbidden: insufficient permissions, required role: " + requiredRole,
			})
			return
		}

		c.Next()
	}
}
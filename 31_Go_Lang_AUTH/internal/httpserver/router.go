package httpserver

import (
	"go-auth/internal/app"
	"go-auth/internal/middleware"
	"go-auth/internal/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(a *app.App) *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", health)

	userRepo := user.NewRepo(a.DB)
	userSrv := user.NewService(userRepo, a.Config.JWTSecret)
	userHandler := user.NewHandler(userSrv)

	// Public routes
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	// Protected/API routes
	api := r.Group("/api")

	api.Use(middleware.AuthRequired(a.Config.JWTSecret))

	api.GET("/files", func(c *gin.Context) {
		userID, ok := middleware.GetUserID(c)
		if !ok || userID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "user id is missing from auth context",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"ok":      true,
			"user_id": userID,
			"userId":  userID,
			"files":   []any{},
		})
	})

	api.GET("/products", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ok":       true,
			"products": []any{},
		})
	})

	admin := api.Group("/admin")
	admin.Use(middleware.RequireRole("admin"))
	admin.GET("/resctricted", func(c *gin.Context) {
		role,_ := middleware.GetRole(c)
		c.JSON(http.StatusOK, gin.H{
			"ok": true,	
			"role": role,
		})
	})
	return r
}

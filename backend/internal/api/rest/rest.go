package rest

import (
	"net/http"
	"spiropoulos94/emailchaser/invite/internal/api/rest/handlers"

	"github.com/gin-gonic/gin"
)

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract user email from the request header.
		userEmail := c.GetHeader("user-email")

		// If user email is empty, return an error.
		if userEmail == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User email not provided in headers"})
			c.Abort()
			return
		}

		// Save user email in the Gin context.
		c.Set("user-email", userEmail)

		// Continue processing the request.
		c.Next()
	}
}

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(UserMiddleware())

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	apiRouter := r.Group("/api")
	{
		handlers.RegisterInvitationGroup(apiRouter)
		handlers.RegisterUserGroup(apiRouter)
	}

	return r
}

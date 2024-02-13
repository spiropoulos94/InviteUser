package rest

import (
	"net/http"
	"spiropoulos94/emailchaser/invite/internal/api/rest/handlers"

	"github.com/gin-gonic/gin"
)

func NewRouter () *gin.Engine {
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})


	apiRouter := r.Group("/api")
	{
		handlers.RegisterUserGroup(apiRouter)
		// handlers.NewInvitationHandler(api)
	}
	
	return r
}
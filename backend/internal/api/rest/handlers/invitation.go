package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InvitationHandlerAll (c *gin.Context) {
	c.String(http.StatusOK, "Got Invitations")
}

func InvitationHandlerOne (c *gin.Context) {
	c.String(http.StatusOK, "Got Invitations One")
}

func NewInvitationHandler(r *gin.RouterGroup) *gin.RouterGroup {
	invitationGroup := r.Group("/invitation")
	{
		invitationGroup.GET("/", InvitationHandlerAll)
		invitationGroup.GET("/:id", InvitationHandlerOne)
	}
	return invitationGroup
}
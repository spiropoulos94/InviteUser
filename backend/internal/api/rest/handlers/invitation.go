package handlers

import (
	"net/http"
	"spiropoulos94/emailchaser/invite/ent"
	"spiropoulos94/emailchaser/invite/internal/db"

	"github.com/gin-gonic/gin"
)

type InvitationHandler struct {
	db *ent.Client
}

func NewInvitationHandler() *InvitationHandler {
	return &InvitationHandler{
		db: db.GetDB(),
	}
}

func RegisterInvitationGroup(r *gin.RouterGroup) {

	invitationH := NewInvitationHandler()

	invitationGroup := r.Group("/invitation")
	{
		invitationGroup.GET("/", invitationH.All)
	}
}

func (u *InvitationHandler) All(c *gin.Context) {

	c.JSON(http.StatusAccepted, gin.H{
		"invitaton": "skata",
	})

}

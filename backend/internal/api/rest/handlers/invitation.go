package handlers

import (
	"net/http"
	"spiropoulos94/emailchaser/invite/ent"
	"spiropoulos94/emailchaser/invite/internal/db"
	"spiropoulos94/emailchaser/invite/internal/utils"

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

	invitationGroup := r.Group("/invites")
	{
		invitationGroup.GET("/", invitationH.All)
		invitationGroup.POST("/", invitationH.Create)
	}
}

func (h *InvitationHandler) All(c *gin.Context) {
	invitations, err := h.db.Invitation.Query().All(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch invitations"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"invitations": invitations})
}

func (h *InvitationHandler) Create(c *gin.Context) {
	var inviteInput struct {
		InviterEmail string `json:"inviter_email" binding:"required,email"`
		InviteeEmail string `json:"invitee_email" binding:"required,email"`
	}

	// Bind request body to the inviteInput struct
	if err := c.ShouldBindJSON(&inviteInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Logic for validating inviter and invitee email domain goes here.
	// Assuming the validation passes for now.
	// TODO

	// Extract domain from inviter and invitee emails
	inviterDomain := utils.ExtractDomain(inviteInput.InviterEmail)
	inviteeDomain := utils.ExtractDomain(inviteInput.InviteeEmail)

	// Check if the domains are the same
	if inviterDomain != inviteeDomain {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Inviter and invitee emails must be from the same domain"})
		return
	}

	// Create an invitation record in the database
	invitation, err := h.db.Invitation.Create().
		SetInviteeEmail(inviteInput.InviteeEmail).
		SetStatus("pending").
		Save(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create invitation"})
		return
	}

	// Logic for sending email notification to the invitee goes here.

	c.JSON(http.StatusCreated, gin.H{"message": "Invitation sent successfully", "invitation": invitation})
}

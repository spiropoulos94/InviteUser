package handlers

import (
	"net/http"
	"spiropoulos94/emailchaser/invite/ent"
	"spiropoulos94/emailchaser/invite/ent/invitation"
	"spiropoulos94/emailchaser/invite/ent/user"
	"spiropoulos94/emailchaser/invite/internal/db"
	"spiropoulos94/emailchaser/invite/internal/utils"
	"strconv"

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
		invitationGroup.POST("/:id/accept", invitationH.AcceptInvitation)
	}
}

func (h *InvitationHandler) AcceptInvitation(c *gin.Context) {

	// Get the invitation ID from the request
	invitationID := c.Param("id")

	invitationIDstr, err := strconv.Atoi(invitationID)
	if err != nil {
		c.JSON(504, gin.H{"error": "Could not convert invitation id to string"})
		return
	}

	// Retrieve the invitation from the database
	invitation, err := h.db.Invitation.Query().
		Where(invitation.ID(invitationIDstr)).
		Only(c)

	if err != nil {
		// Handle the error, e.g., return 404 if invitation is not found
		c.JSON(404, gin.H{"error": "Invitation not found"})
		return
	}

	// Set the status of the invitation to "accepted"
	invitation, err = invitation.Update().
		SetStatus("accepted").
		Save(c)

	if err != nil {
		// Handle the error, e.g., return 500 if unable to update the invitation
		c.JSON(500, gin.H{"error": "Failed to accept invitation"})
		return
	}

	// Create a new user with all fields and a blank password
	newUser, err := h.db.User.Create().
		SetName(invitation.InviteeEmail).
		SetEmail(invitation.InviteeEmail).
		SetPassword("").
		Save(c)

	if err != nil {
		// Handle the error, e.g., return 500 if unable to create the user
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Retrieve the inviter's team
	inviterTeam, err := invitation.QueryInviter().QueryTeams().Only(c)

	if err != nil {
		// Handle the error, e.g., return 500 if unable to retrieve the inviter's team
		c.JSON(500, gin.H{"error": "Failed to retrieve inviter's team"})
		return
	}

	// Save the new user on the same team
	_, err = inviterTeam.Update().
		AddUsers(newUser).
		Save(c)

	if err != nil {
		// Handle the error, e.g., return 500 if unable to add the user to the team
		c.JSON(500, gin.H{"error": "Failed to add user to team"})
		return
	}

	// Respond with success message
	c.JSON(200, gin.H{"message": "Invitation accepted and user added to team", "user": newUser})
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
		// InviterEmail string `json:"inviter_email" binding:"required,email"`
		InviteeEmail string `json:"invitee_email" binding:"required,email"`
	}

	inviterEmailRaw, exist := c.Get("user-email")
	if !exist {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User inviter email is missing"})
		return
	}

	inviterEmail, ok := inviterEmailRaw.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User inviter email is not a string"})
		return
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
	inviterDomain := utils.ExtractDomain(inviterEmail)
	inviteeDomain := utils.ExtractDomain(inviteInput.InviteeEmail)

	// Check if the domains are the same
	if inviterDomain != inviteeDomain {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Inviter and invitee emails must be from the same domain"})
		return
	}

	inviterUser, err := h.db.User.Query().Where(user.EmailEQ(inviterEmail)).Only(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Inviter user was not found"})
		return
	}

	// Create an invitation record in the database
	invitation, err := h.db.Invitation.Create().
		SetInviteeEmail(inviteInput.InviteeEmail).
		SetStatus("pending").
		SetInviter(inviterUser).
		Save(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create invitation"})
		return
	}

	// Logic for sending email notification to the invitee goes here.

	c.JSON(http.StatusCreated, gin.H{"message": "Invitation sent successfully", "invitation": invitation})
}

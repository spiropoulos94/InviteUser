package handlers

import (
	"fmt"
	"net/http"
	"spiropoulos94/emailchaser/invite/ent"
	"spiropoulos94/emailchaser/invite/ent/team"
	"spiropoulos94/emailchaser/invite/ent/user"
	"spiropoulos94/emailchaser/invite/internal/db"
	"spiropoulos94/emailchaser/invite/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	db *ent.Client
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		db: db.GetDB(),
	}
}

func RegisterUserGroup(r *gin.RouterGroup) {

	userH := NewUserHandler()

	userGroup := r.Group("/user")
	{
		userGroup.GET("/", userH.All)
		userGroup.GET("/:id", userH.FindById)
		userGroup.POST("/", userH.Create)
	}
}

func (u *UserHandler) Create(c *gin.Context) {
	var userInput struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	// Bind request body to the userInput struct
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the email is already registered
	existingUser, err := u.db.User.Query().
		Where(user.EmailEQ(userInput.Email)).
		Exist(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if existingUser {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
		return
	}

	// Extract host from email address
	host := utils.ExtractDomain(userInput.Email)

	// Check if team exists for the host
	team, err := u.db.Team.Query().
		Where(team.NameEQ(host)).
		Only(c)
	if err != nil {
		// Create a new team if it doesn't exist
		team, err = u.db.Team.Create().
			SetName(host).
			Save(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create team"})
			return
		}
	}

	// Create the user in the database
	newUser, err := u.db.User.Create().
		SetName(userInput.Name).
		SetEmail(userInput.Email).
		SetPassword(userInput.Password).AddGroups(team).
		Save(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Fetch user's teams
	userTeams, err := newUser.QueryGroups().All(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user teams"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": newUser, "user_teams": userTeams})
}

func (u *UserHandler) All(c *gin.Context) {
	users, err := u.db.User.Query().All(c)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
func (u *UserHandler) FindById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := u.db.User.Get(c, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

package handlers

import (
	"fmt"
	"net/http"
	"spiropoulos94/emailchaser/invite/ent"
	"spiropoulos94/emailchaser/invite/ent/user"
	"spiropoulos94/emailchaser/invite/internal/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	db *ent.Client
}

func NewUserHandler() *UserHandler{
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
	}
}

func (u *UserHandler) All (c *gin.Context) {

	us, err := u.db.User.Query().Where(user.Name("test")).Only(c)
	if err!= nil{
		fmt.Println(err)
	}
    
	c.JSON(http.StatusAccepted, gin.H{
		"user":us,
	})
	
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
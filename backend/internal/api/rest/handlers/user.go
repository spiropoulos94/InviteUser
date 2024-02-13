package handlers

import (
	"fmt"
	"net/http"
	"spiropoulos94/emailchaser/invite/ent"
	"spiropoulos94/emailchaser/invite/ent/user"
	"spiropoulos94/emailchaser/invite/internal/db"

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


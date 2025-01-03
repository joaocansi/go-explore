package users

import (
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *UserService
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (u *UserHandler) Create(c *gin.Context) {
	createUserInput := CreateUserInput{}
	if err := c.ShouldBind(&createUserInput); err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": "bad format",
		})
		return
	}

	user, err := u.service.Create(createUserInput)
	if err != nil {
		c.JSON(400, gin.H{
			"status":  400,
			"message": err.Error(),
		})
		return
	}

	c.JSON(201, user)
}

package users

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRoutes(router *gin.Engine, db *gorm.DB) {
	repository := NewUserRepository(db)
	service := NewUserService(repository)
	handler := NewUserHandler(service)
	router.POST("/users", handler.Create)
}

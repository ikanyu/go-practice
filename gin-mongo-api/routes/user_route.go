package routes

import (
	"gin-mongo-api/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controller.CreateUser)
	router.GET("/user/:userId", controller.GetUser)
	router.PUT("/user/:userId", controller.EditAUser)
	router.DELETE("/user/:userId", controller.DeleteUser)
	router.GET("/users", controller.GetAllUsers)
}

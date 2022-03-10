package routes

import (
	"gin-mongo-api/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/api/v2/users", controllers.GetAllUsers())
	router.GET("/api/v2/users/:id", controllers.GetUser())
	router.POST("/api/v2/users", controllers.CreateUser())
	router.PUT("/api/v2/users", controllers.UpdateUser())
	router.DELETE("/api/v2/users/:id", controllers.DeleteUser())
}

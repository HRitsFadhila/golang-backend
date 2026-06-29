package routes

import (
	"github.com/HRitsFadhila/golang-backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	router := gin.Default()

	router.POST("/api/register", controllers.Register)
	router.POST("api/login", controllers.Login)

	return  router
}
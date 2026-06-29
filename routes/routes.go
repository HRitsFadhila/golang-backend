package routes

import (
	"github.com/HRitsFadhila/golang-backend/controllers"
	"github.com/HRitsFadhila/golang-backend/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	router.POST("/api/register", controllers.Register)
	router.POST("api/login", controllers.Login)
	router.GET("/api/users", middlewares.AuthMiddleware(), controllers.FindUsers)
	router.POST("/api/users", middlewares.AuthMiddleware(), controllers.CreateUser)
	router.GET("/api/users/:id", middlewares.AuthMiddleware(), controllers.FinsUserById)
	router.PATCH("api/users/:id", middlewares.AuthMiddleware(), controllers.UpdateUser)
	router.DELETE("api/users/:id", middlewares.AuthMiddleware(), controllers.DeleteUser)

	return  router
}
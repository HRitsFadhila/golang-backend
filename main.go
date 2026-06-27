package main

import (
	"github.com/HRitsFadhila/golang-backend/config"
	"github.com/HRitsFadhila/golang-backend/database"
	"github.com/gin-gonic/gin"
)

func main(){
	config.LoadEnv()
	database.InitDB()
	router := gin.Default()

	router.GET("/", func (c *gin.Context){
		c.JSON(200, gin.H{
			"message":"Hello World!",
		})
	})

	router.Run(":" + config.GetEnv("APP_PORT","3000"))
}